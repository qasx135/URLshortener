package URLservice

import "context"

type Repository interface {
	Create(ctx context.Context, URL string, alias string) error
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
