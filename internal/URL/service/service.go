package URLservice

import (
	"context"
	"url-shortener/internal/URL/model"
)

type Repository interface {
	Create(ctx context.Context, URL *model.URLModel) error
	Get(ctx context.Context) ([]string, error)
	GetOne(ctx context.Context, alias string) (string, error)
	Update(ctx context.Context, URL string) error
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
func (s *Service) Get(ctx context.Context) ([]string, error) {
	return s.repo.Get(ctx)
}
func (s *Service) GetOne(ctx context.Context, alias string) (string, error) {
	return s.repo.GetOne(ctx, alias)
}
func (s *Service) Update(ctx context.Context, URL string) error {
	return s.repo.Update(ctx, URL)
}
func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
