package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/korasdor/go-chess/internal/config"
)

const (
	usersTable = "users"
)

func NewPostgresDB(cfg *config.PostgresConfig) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=%s`, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DatabaseName, cfg.SSLMode)

	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		return nil, err
	}

	return db, err
}
