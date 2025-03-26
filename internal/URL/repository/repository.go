package URLrepository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
	"url-shortener/internal/URL/model"
	URLservice "url-shortener/internal/URL/service"
	"url-shortener/pkg/logger"
)

type Repository struct {
	db *pgx.Conn
}

func (r *Repository) Create(ctx context.Context, URL *model.URLModel) error {
	q := `INSERT INTO 
    		url (url, alias) 
			VALUES ($1, $2) 
			RETURNING id
		`
	err := r.db.QueryRow(ctx, q, URL.Url, URL.Alias).Scan(&URL.ID)
	if err != nil {
		//TODO: Разобраться и доделать
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "Stop ruin my program", zap.Error(err))
		return err
	}
	return nil
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
