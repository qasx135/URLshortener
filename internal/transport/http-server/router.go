package httpserver

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"net/http"
	"time"
	"url-shortener/internal/transport/http-server/url-handlers"
)

type Config struct {
	Addr        string        `yaml:"address" env:"HTTP_ADDR" env-default:"localhost:8080" env-required:"true"`
	TimeOut     time.Duration `yaml:"timeout" env:"HTTP_TIMEOUT" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"HTTP_IDLE_TIMEOUT" env-default:"60s"`
}

type Router struct {
	config  Config
	Router  *chi.Mux
	Handler url_handlers.Handler
}

func NewRouter(cfg Config, h *url_handlers.Handler) *Router {
	cs := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},                   // Укажите разрешённые источники (фронт)
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Разрешённые методы
		AllowedHeaders:   []string{"Content-Type", "Authorization"},           // Разрешённые заголовки
		AllowCredentials: true,                                                // Если нужны cookies или токены
	})

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(cs.Handler)
	r.Post("/url", h.Create)
	r.Get("/url/all-urls", h.Get)
	r.Get("/url/{alias}", h.GetOne)
	r.Put("/url/{id}", h.Update)
	return &Router{config: cfg, Router: r, Handler: *h}
}

func (r *Router) Run(cfg Config, router *Router) {
	srv := &http.Server{
		Addr:         cfg.Addr,
		Handler:      router.Router,
		ReadTimeout:  cfg.TimeOut,
		WriteTimeout: cfg.TimeOut,
		IdleTimeout:  cfg.IdleTimeout,
	}
	fmt.Println("Router.Run")
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		_ = fmt.Errorf("failed to start http server: %w", err)
	}
	fmt.Println("Server stopped")

}
