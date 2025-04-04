package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type Config struct {
	Host     string `yaml:"host" env:"POSTGRES_HOST" env-required:"true"`
	Port     string `yaml:"port" env:"POSTGRES_PORT" env-required:"true"`
	Username string `yaml:"username" env:"POSTGRES_USER" env-required:"true"`
	Password string `yaml:"password" env:"POSTGRES_PASSWORD" env-required:"true"`
	Database string `yaml:"database" env:"POSTGRES_DB" env-required:"true"`
}

func New(ctx context.Context, config Config) (*pgx.Conn, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %v", err)
	}

	PrepareTables(ctx, conn)
	return conn, nil
}

func PrepareTables(ctx context.Context, conn *pgx.Conn) {
	queryText := `create table urls_table
(
    id    serial not null
        constraint urls_table_pk
            primary key,
    url   text   not null
        constraint urls_table_pk_2
            unique,
    alias text   not null
        constraint urls_table_pk_3
            unique
);`

	_, err := conn.Exec(ctx, queryText)
	if err != nil {
		_ = fmt.Errorf("cannot create table")
	}
}
