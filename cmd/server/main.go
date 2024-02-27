package main

import (
	"context"

	"GophKeeper/internal/server/app"
)

func main() {
	ctx := context.Background()

	err := app.Run(ctx)
	if err != nil {
		panic(err)
	}
}
