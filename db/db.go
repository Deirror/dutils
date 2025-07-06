package db

import (
	"database/sql"

	"github.com/Deirror/dutils/cfg"
)

// SQLDatabase wraps a standard sql.DB instance and provides
// lifecycle management for SQL database connections.
type SQLDatabase struct {
	Db *sql.DB
}

// NewSQLDatabase initializes and returns a new SQLDatabase instance
// using the provided configuration. It sets connection pool parameters
// such as maximum open/idle connections and connection lifetime.
//
// Returns an error if the connection could not be established.
func NewSQLDatabase(cfg *cfg.DBConfig) (*SQLDatabase, error) {
	db, err := Connect(cfg.Driver, cfg.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(int(cfg.PoolSize))
	db.SetMaxIdleConns(int(cfg.MaxIdle))
	db.SetConnMaxLifetime(cfg.MaxLifetime)

	return &SQLDatabase{
		Db: db,
	}, nil
}

// Close closes the underlying database connection.
// It is safe to call Close multiple times; if the DB is already nil, it does nothing.
func (db *SQLDatabase) Close() error {
	if db.Db != nil {
		return db.Db.Close()
	}
	return nil
}

// Connect opens a new database connection using the given driver name and DSN (Data Source Name).
// It pings the database to verify the connection is valid before returning.
//
// Returns an error if the connection could not be opened or pinged.
func Connect(driver, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// Ping verifies that the database connection is still alive.
// It is useful for health checks or readiness probes.
func (db *SQLDatabase) Ping() error {
	return db.Db.Ping()
}
