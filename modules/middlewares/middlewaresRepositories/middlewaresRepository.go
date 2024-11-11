package middlewaresRepositories

import "github.com/jmoiron/sqlx"

type IMiddlewareRepository interface{}

type MiddlewareRepository struct {
	db *sqlx.DB
}

func MiddlewareRepositoryInit(db *sqlx.DB) IMiddlewareRepository {
	return &MiddlewareRepository{db: db}
}
