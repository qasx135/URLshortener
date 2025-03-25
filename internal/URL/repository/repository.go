package URLrepository

import (
	"context"
	"github.com/jackc/pgx/v5"
	URLservice "url-shortener/internal/URL/service"
)

type Repository struct {
	db *pgx.Conn
}

func (r *Repository) Create(ctx context.Context, URL string, alias string) error {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Get(ctx context.Context) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) GetOne(ctx context.Context, alias string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Update(ctx context.Context, URL string) error {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(db *pgx.Conn) URLservice.Repository {
	return &Repository{db: db}
}
