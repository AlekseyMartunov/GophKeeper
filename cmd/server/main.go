package main

import (
	"GophKeeper/internal/app/server"
	"context"
)

func main() {
	ctx := context.Background()

	err := server.Run(ctx)
	if err != nil {
		panic(err)
	}
}
