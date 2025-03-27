package URLservice

import (
	"context"
	"url-shortener/internal/URL/model"
)

type Repository interface {
	Create(ctx context.Context, URL *model.URLModel) error
	Get(ctx context.Context) ([]model.URLModel, error)
	GetOne(ctx context.Context, alias string) (model.URLModel, error)
	Update(ctx context.Context, id int, URL string, alias string) error
	Delete(ctx context.Context, id string) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, URL *model.URLModel) error {
	return s.repo.Create(ctx, URL)
}
func (s *Service) Get(ctx context.Context) ([]model.URLModel, error) {
	return s.repo.Get(ctx)
}
func (s *Service) GetOne(ctx context.Context, alias string) (model.URLModel, error) {
	return s.repo.GetOne(ctx, alias)
}
func (s *Service) Update(ctx context.Context, id int, URL string, alias string) error {
	return s.repo.Update(ctx, id, URL, alias)
}
func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
