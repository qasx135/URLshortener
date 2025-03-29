package URLrepository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5"
	"url-shortener/internal/URL/model"
	URLservice "url-shortener/internal/URL/service"
	"url-shortener/pkg/logger"
)

type Repository struct {
	db *pgx.Conn
}

func (r *Repository) Create(ctx context.Context, URL *model.URLModel) error {
	q := `INSERT INTO 
    		urls_table (url, alias) 
			VALUES ($1, $2) 
		`

	if _, err := r.db.Exec(ctx, q, URL.Url, URL.Alias); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// Проверяем код ошибки (unique constraint violation)
			if pgErr.Code == "23505" {
				logger.GetLoggerFromCtx(ctx).Info(ctx, "URL already exists, skipping insertion")
				return nil // Игнорируем это, так как строка уже существует
			}
		}
		// Для остальных ошибок вызываем фатальный логгер
		//logger.GetLoggerFromCtx(ctx).Fatal(ctx, "Unexpected error: %v", zap.Error(err))
	}
	return nil
}

func (r *Repository) Get(ctx context.Context) ([]model.URLModel, error) {
	q := `SELECT id, url, alias FROM urls_table`
	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	urls := make([]model.URLModel, 0)
	for rows.Next() {
		var url model.URLModel
		err = rows.Scan(&url.ID, &url.Url, &url.Alias)
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}
	return urls, nil
}

func (r *Repository) GetOne(ctx context.Context, alias string) (model.URLModel, error) {

	q := `SELECT id, url, alias FROM urls_table WHERE ALIAS = $1`
	var url model.URLModel
	err := r.db.QueryRow(ctx, q, alias).Scan(&url.ID, &url.Url, &url.Alias)
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Info(ctx, "No data with the given alias")
		//return model.URLModel{}, nil
	}
	fmt.Println(url)
	return url, nil
}

func (r *Repository) Update(ctx context.Context, id string, alias string) error {
	q := `UPDATE urls_table SET alias = $1 WHERE id = $2`
	commandTag, err := r.db.Exec(ctx, q, alias, id)
	if err != nil {
		return fmt.Errorf("failed to execute update query: %w", err)
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no rows were updated for id %d", id)
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	//TODO: Выбрать как реализовать (флагом без удаления/полное удаление)
	panic("implement me")
}

func NewRepository(db *pgx.Conn) URLservice.Repository {
	return &Repository{db: db}
}
