package main

import (
	"context"
	URLrepository "url-shortener/internal/URL/repository"
	URLservice "url-shortener/internal/URL/service"
	"url-shortener/internal/config"
	http_server "url-shortener/internal/transport/http-server"
	"url-shortener/internal/transport/http-server/url-handlers"
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
	if myService == nil {
		logger.GetLoggerFromCtx(ctx).Info(ctx, "Failed to initialize service")
	}

	//TODO: init router
	handler := url_handlers.NewHandler(myService, ctx)
	router := http_server.NewRouter(cfg.RouterConfig, handler)
	router.Run(cfg.RouterConfig, router)

	//TODO: init server
}
