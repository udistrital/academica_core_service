package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// DB es la instancia global del pool de conexiones Oracle.
var DB *sqlx.DB

// InitDB abre la conexión Oracle y configura el pool para carga concurrente.
func InitDB(dsn string) error {
	db, err := sqlx.Open("oracle", dsn)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(30 * time.Minute)
	if err = db.Ping(); err != nil {
		return err
	}
	DB = db
	return nil
}
