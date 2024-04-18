package app

import (
	fileStorage "GophKeeper/internal/adapters/db/minio/file"
	cardsRepo "GophKeeper/internal/adapters/db/postgres/card"
	"GophKeeper/internal/adapters/db/postgres/migration"
	pairsRepo "GophKeeper/internal/adapters/db/postgres/pairs"
	tokenRepo "GophKeeper/internal/adapters/db/postgres/token"
	usersRepo "GophKeeper/internal/adapters/db/postgres/users"
	cardHandlers "GophKeeper/internal/adapters/http/cards/handlers"
	cardRouter "GophKeeper/internal/adapters/http/cards/router"
	fileHandlers "GophKeeper/internal/adapters/http/files/handlers"
	fileRouter "GophKeeper/internal/adapters/http/files/router"
	pairHandlers "GophKeeper/internal/adapters/http/pair/handlers"
	pairRouter "GophKeeper/internal/adapters/http/pair/router"
	tokenHandlers "GophKeeper/internal/adapters/http/tokens/handlers"
	tokenRouter "GophKeeper/internal/adapters/http/tokens/router"
	userHandlers "GophKeeper/internal/adapters/http/users/handlers"
	userRouter "GophKeeper/internal/adapters/http/users/router"
	"GophKeeper/internal/config"
	"GophKeeper/internal/jwt"
	"GophKeeper/internal/middleware/authenticationhttp"
	middlewareHTTPLogin "GophKeeper/internal/middleware/loginhttp"
	cardService "GophKeeper/internal/usecase/creditcard"
	fileService "GophKeeper/internal/usecase/file"
	pairService "GophKeeper/internal/usecase/pairs"
	tokenService "GophKeeper/internal/usecase/token"
	usersService "GophKeeper/internal/usecase/users"
	"GophKeeper/pkg/hasher"
	"GophKeeper/pkg/logger"
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
	cfg.ParseFlags()

	err := migration.MigrationsUp(cfg.PostgresDSN())
	if err != nil {
		return fmt.Errorf("migration err: %w", err)
	}

	pool, err := pgxpool.New(ctx, cfg.PostgresDSN())
	if err != nil {
		return fmt.Errorf("pool creation err: %w", err)
	}

	hash := hasher.NewHasher(cfg.Salt())
	log := logger.NewLogger()

	tokenRepo := tokenRepo.NewTokenStorage(pool)
	tokenManager := jwt.NewTokenManager(tokenExpTime, cfg.SecretKey())
	tokenService := tokenService.NewTokenService(tokenRepo, hash, tokenManager, cfg)

	middlewareHTTPLogin := middlewareHTTPLogin.NewLoggerMiddleware(log)
	middlewareHTTPAuth := authenticationhttp.NewAuthMiddleware(log, tokenService)

	userStorage := usersRepo.NewUserStorage(pool)
	userService := usersService.NewUserService(userStorage, hash)
	userHandler := userHandlers.NewUserHandler(userService, log)
	userRouter := userRouter.NewUserControllerHTTP(userHandler, middlewareHTTPLogin)

	pairStorage := pairsRepo.NewPairsStorage(pool)
	pairService := pairService.NewPairService(pairStorage)
	pairHandler := pairHandlers.NewPairHandler(log, pairService)
	pairRouter := pairRouter.NewPairControllerHTTP(pairHandler, middlewareHTTPLogin, middlewareHTTPAuth)

	cardStorage := cardsRepo.NewCardStorage(pool)
	cardService := cardService.NewCardService(cardStorage)
	cardHandler := cardHandlers.NewCardHandler(log, cardService)
	cardRouter := cardRouter.NewCardControllerHTTP(cardHandler, middlewareHTTPLogin, middlewareHTTPAuth)

	tokenHandler := tokenHandlers.NewTokenHandler(log, tokenService, userService)
	tokenRouter := tokenRouter.NewTokenControllerHTTP(tokenHandler, middlewareHTTPLogin, middlewareHTTPAuth)

	fileStorage, err := fileStorage.NewFileStorage(cfg)
	if err != nil {
		return err
	}
	fileService := fileService.NewFileService(fileStorage)
	fileHandler := fileHandlers.NewFileHandler(log, fileService)
	fileRouter := fileRouter.NewFileControllerHTTP(fileHandler, middlewareHTTPLogin, middlewareHTTPAuth)

	e := echo.New()

	userRouter.Route(e)
	pairRouter.Route(e)
	cardRouter.Route(e)
	tokenRouter.Route(e)
	fileRouter.Route(e)

	srv := http.Server{
		Addr:    cfg.RunAddr(),
		Handler: e,
	}

	if err = srv.ListenAndServe(); err != http.ErrServerClosed {
		return fmt.Errorf("listen and serve error: %w", err)
	}
	return nil
}
