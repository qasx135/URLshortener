package main

import (
	"context"
	"fmt"
	"url-shortener/internal/config"
	"url-shortener/pkg/logger"
	"url-shortener/pkg/postgres"
)

func main() {
	//TODO: init config: cleanenv
	cfg := config.NewConfig()

	fmt.Println(cfg)

	//TODO: init logger: zap
	ctx := context.Background()
	ctx, _ = logger.New(ctx)

	//TODO: init storage: postgres
	db, _ := postgres.New(cfg.PostgresConfig)
	err := postgres.PrepareTables(ctx, db)
	if err != nil {
		return
	}

	//TODO: init router

	//TODO: init server
}
