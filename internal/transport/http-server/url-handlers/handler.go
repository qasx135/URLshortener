package url_handlers

import (
	"context"
	"errors"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"net/http"
	"url-shortener/internal/URL/model"
	"url-shortener/internal/api/response"
	"url-shortener/internal/randomAlias"
)

const aliasLength = 7

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	response.Response
	Alias string `json:"alias,omitempty"`
}

type Service interface {
	Create(ctx context.Context, URL *model.URLModel) error
	Get(ctx context.Context) ([]model.URLModel, error)
	GetOne(ctx context.Context, alias string) (model.URLModel, error)
	Update(ctx context.Context, id int, URL string, alias string) error
	Delete(ctx context.Context, id string) error
}

type Handler struct {
	service Service
	Ctx     context.Context
}

func NewHandler(service Service, ctx context.Context) *Handler {
	return &Handler{service: service, Ctx: ctx}
}
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		render.JSON(w, r, response.Error("Failed to decode request body"))
		return // Stopping handler
	}
	if err = validator.New().Struct(req); err != nil {
		var validateErr validator.ValidationErrors
		errors.As(err, &validateErr)

		render.JSON(w, r, response.Error("Failed to validate request body"))
		render.JSON(w, r, response.ValidationError(validateErr))

		return
	}

	alias := req.Alias
	if alias == "" {
		alias = randomAlias.NewRandomAlias(aliasLength)
	}
	err = h.service.Create(h.Ctx, &model.URLModel{Url: req.URL, Alias: alias})
	if err != nil {
		return
	}
	render.JSON(w, r, Response{
		Response: response.OK(),
		Alias:    alias,
	})
}
func (h *Handler) Get(w http.ResponseWriter, r *http.Request)    {}
func (h *Handler) GetOne(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {}
