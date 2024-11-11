package safesql

import (
	"context"
	"database/sql"
	"github.com/Gianluigi/def-prog-exercises/safesql/internal/raw"
)

type compileTimeConstant string

type TrustedSQL struct {
	s string
}

type DB struct {
	db *sql.DB
}
type Rows = sql.Rows
type Result = sql.Result

func init() {
	raw.TrustedSQLCtor = func(unsafe string) TrustedSQL {
		return TrustedSQL{unsafe}
	}
}

func New(text compileTimeConstant) TrustedSQL {
	return TrustedSQL{s: string(text)}
}

func Open(driverName string, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	return &DB{db}, err
}

func (db *DB) QueryContext(ctx context.Context, query TrustedSQL, args ...any) (*Rows, error) {
	return db.db.QueryContext(ctx, query.s, args...)
}

func (db *DB) ExecContext(ctx context.Context, query TrustedSQL, args ...any) (Result, error) {
	return db.db.ExecContext(ctx, query.s, args...)
}
