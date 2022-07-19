package main

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectToDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "postgres://user:mypassword@postgres:5432/testdb?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}

//func MigrateDB() error {
//	m, err := migrate.New(
//		"file://migrations",
//		"postgres://user:mypassword@postgres:5432/testdb?sslmode=disable",
//	)
//	if err != nil {
//		return err
//	}
//
//	err = m.Up()
//	if err == migrate.ErrNoChange {
//		return nil
//	}
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
