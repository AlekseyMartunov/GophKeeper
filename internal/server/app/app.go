package app

import (
	usersRepo "GophKeeper/internal/server/adapters/db/postgres/users"
	"GophKeeper/internal/server/adapters/db/postgres/users/migration"
	userHandlers "GophKeeper/internal/server/adapters/http/users/handlers"
	userRouter "GophKeeper/internal/server/adapters/http/users/router"
	"GophKeeper/internal/server/config"
	"GophKeeper/internal/server/hasher"
	"GophKeeper/internal/server/jwt"
	usersService "GophKeeper/internal/server/usecase/users"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

const tokenExpTime = time.Second * 60 * 60 * 24

func Run(ctx context.Context) error {
	cfg := config.NewConfig()

	err := migration.UsersMigration(cfg.PostgresDSN())
	if err != nil {
		return fmt.Errorf("migration err: %w", err)
	}

	pool, err := pgxpool.New(ctx, cfg.PostgresDSN())
	if err != nil {
		return fmt.Errorf("pool creation err: %w", err)
	}

	token := jwt.NewTokenManager(tokenExpTime, cfg.SecretKey())
	hash := hasher.NewHasher(cfg.Salt())

	userStorage := usersRepo.NewUserStorage(pool)
	userService := usersService.NewUserService(userStorage, hash)
	userHandler := userHandlers.NewUserHandler(userService, nil, token)
	userRouter := userRouter.NewUserControllerHTTP(userHandler)

	userApp := userRouter.Route()

	mainApp := fiber.New()

	mainApp.Mount("/api/users", userApp)

	//mainApp.ListenTLS("", "", "")
	mainApp.Listen(":8080")

	return nil
}
