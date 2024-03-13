package server

import (
	"GophKeeper/internal/adapters/db/postgres/migration"
	pairsRepo "GophKeeper/internal/adapters/db/postgres/pairs"
	usersRepo "GophKeeper/internal/adapters/db/postgres/users"
	userHandlers "GophKeeper/internal/adapters/http/users/handlers"
	userRouter "GophKeeper/internal/adapters/http/users/router"
	"GophKeeper/internal/config"
	"GophKeeper/internal/hasher"
	"GophKeeper/internal/jwt"
	"GophKeeper/internal/logger"
	middlewareHTTPLogin "GophKeeper/internal/middleware/loginhttp"
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

	token := jwt.NewTokenManager(tokenExpTime, cfg.SecretKey())
	hash := hasher.NewHasher(cfg.Salt())
	log := logger.NewLogger()
	middlewareHTTPLogin := middlewareHTTPLogin.NewLoggerMiddleware(log)

	userStorage := usersRepo.NewUserStorage(pool)
	userService := usersService.NewUserService(userStorage, hash)
	userHandler := userHandlers.NewUserHandler(userService, log, token)
	userRouter := userRouter.NewUserControllerHTTP(userHandler, middlewareHTTPLogin)

	pairStorage := pairsRepo.NewPairsStorage(pool)
	pairService := pairService.NewPairService(pairStorage)

	fmt.Println(pairService)

	e := echo.New()
	userRouter.Route(e)

	srv := http.Server{
		Addr:    cfg.RunAddr(),
		Handler: e,
	}

	if err = srv.ListenAndServe(); err != http.ErrServerClosed {
		return fmt.Errorf("listen and serve error: %w", err)
	}

	return nil
}
