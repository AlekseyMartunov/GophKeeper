package server

import (
	cardsRepo "GophKeeper/internal/adapters/db/postgres/card"
	"GophKeeper/internal/adapters/db/postgres/migration"
	pairsRepo "GophKeeper/internal/adapters/db/postgres/pairs"
	tokenRepo "GophKeeper/internal/adapters/db/postgres/token"
	usersRepo "GophKeeper/internal/adapters/db/postgres/users"
	cardHandlers "GophKeeper/internal/adapters/http/cards/handlers"
	"GophKeeper/internal/adapters/http/cards/router"
	pairHandlers "GophKeeper/internal/adapters/http/pair/handlers"
	pairRouter "GophKeeper/internal/adapters/http/pair/router"
	userHandlers "GophKeeper/internal/adapters/http/users/handlers"
	userRouter "GophKeeper/internal/adapters/http/users/router"
	"GophKeeper/internal/config"
	"GophKeeper/internal/hasher"
	"GophKeeper/internal/jwt"
	"GophKeeper/internal/logger"
	"GophKeeper/internal/middleware/authenticationhttp"
	middlewareHTTPLogin "GophKeeper/internal/middleware/loginhttp"
	cardService "GophKeeper/internal/usecase/creditcard"
	pairService "GophKeeper/internal/usecase/pairs"
	usersService "GophKeeper/internal/usecase/users"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const tokenExpTime = time.Second * 60 * 60 * 24

func Run(ctx context.Context) error {
	cfg := config.NewConfig()

	err := migration.MigrationsUp(cfg.PostgresDSN())
	if err != nil {
		return fmt.Errorf("migration err: %w", err)
	}

	pool, err := pgxpool.New(ctx, cfg.PostgresDSN())
	if err != nil {
		return fmt.Errorf("pool creation err: %w", err)
	}

	tokenRepo := tokenRepo.NewTokenStorage(pool)
	token := jwt.NewTokenManager(tokenExpTime, cfg.SecretKey(), tokenRepo)

	hash := hasher.NewHasher(cfg.Salt())
	log := logger.NewLogger()

	middlewareHTTPLogin := middlewareHTTPLogin.NewLoggerMiddleware(log)
	middlewareHTTPAuth := authenticationhttp.NewAuthMiddleware(token, log)

	userStorage := usersRepo.NewUserStorage(pool)
	userService := usersService.NewUserService(userStorage, hash)
	userHandler := userHandlers.NewUserHandler(userService, log, token)
	userRouter := userRouter.NewUserControllerHTTP(userHandler, middlewareHTTPLogin)

	pairStorage := pairsRepo.NewPairsStorage(pool)
	pairService := pairService.NewPairService(pairStorage)
	pairHandler := pairHandlers.NewPairHandler(log, pairService)
	pairRouter := pairRouter.NewPairControllerHTTP(pairHandler, middlewareHTTPLogin, middlewareHTTPAuth)

	cardStorage := cardsRepo.NewCardStorage(pool)
	cardService := cardService.NewCardService(cardStorage)
	cardHandler := cardHandlers.NewCardHandler(log, cardService)
	cardRouter := cardrouter.NewCardControllerHTTP(cardHandler, middlewareHTTPLogin, middlewareHTTPAuth)

	e := echo.New()
	userRouter.Route(e)
	pairRouter.Route(e)
	cardRouter.Route(e)

	srv := http.Server{
		Addr:    cfg.RunAddr(),
		Handler: e,
	}

	if err = srv.ListenAndServe(); err != http.ErrServerClosed {
		return fmt.Errorf("listen and serve error: %w", err)
	}

	return nil
}
