package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DB struct {
	master *bun.DB
}

type Config struct {
	Database string
	User     string
	Host     string
	Port     string
	SSLMode  string
	Password string
}

func NewBunDB(ctx context.Context, cfg *Config) *DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())
	err := db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	return &DB{
		master: db,
	}
}

func (d *DB) AddQueryHook(hook bun.QueryHook) {
	d.master.AddQueryHook(hook)
}

func (d *DB) GetMaster() *bun.DB {
	return d.master
}

func (d *DB) StartTransaction(ctx context.Context) (bun.Tx, error) {
	return d.master.BeginTx(ctx, nil)
}
