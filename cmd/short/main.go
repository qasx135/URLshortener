package main

import (
	"context"
	"url-shortener/internal/URL/model"
	URLrepository "url-shortener/internal/URL/repository"
	URLservice "url-shortener/internal/URL/service"
	"url-shortener/internal/config"
	"url-shortener/pkg/logger"
	"url-shortener/pkg/postgres"
)

func main() {
	cfg := config.NewConfig()

	ctx := context.Background()
	ctx, _ = logger.New(ctx)

	db, _ := postgres.New(ctx, cfg.PostgresConfig)

	repo := URLrepository.NewRepository(db)
	myService := URLservice.NewService(repo)

	testModel := model.URLModel{
		Url:   "https://google.com",
		Alias: "google",
	}
	err := myService.Create(ctx, &testModel)
	if err != nil {
		return
	}

	//TODO: init router

	//TODO: init server
}
