package database

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	db *sql.DB
)

type TransactionFunc func(context.Context, boil.ContextTransactor) error

func Setup(name string) error {
	wd, err := os.Getwd()

	if err != nil {
		return err
	}

	dsn := "file:" + filepath.Join(wd, name) + "?cache=shared&mode=rwc"

	db, err = sql.Open(`sqlite3`, dsn)

	if err != nil {
		return err
	}

	return nil
}

func Teardown() error {
	return db.Close()
}

func Transaction(ctx context.Context, fn TransactionFunc) error {
	tx, err := db.Begin()

	if err != nil {
		return err
	}
	if err = fn(ctx, tx); err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return err
}
