package database

import (
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/premwitthawas/basic-api/config"
)

func DbConnect(cfg config.IDbConfig) *sqlx.DB {
	db, err := sqlx.Connect("pgx", cfg.Url())
	if err != nil {
		log.Fatalf("Connected DB Failed ❌: %v\n", err)
	}
	log.Printf("Connected DB Success ✅: %v\n", cfg.Url())
	db.DB.SetMaxOpenConns(cfg.MaxConnection())
	return db
}
