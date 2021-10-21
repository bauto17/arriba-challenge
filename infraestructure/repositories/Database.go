package repositories

import (
	"arriba-challenge/domain/model"
	"arriba-challenge/infraestructure/config"
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"strings"
	"time"
)

type Database interface {
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string) (pgx.Rows, error)
}

type DatabaseImpl struct {
	conn *pgxpool.Pool
}

const connStringTemplate = "host=%s user=%s dbname=%s password=%s port=%d sslmode=disable"

func NewDatabase(config config.DatabaseConfig) (Database, error) {

	uri := fmt.Sprintf(
		connStringTemplate,
		config.Host,
		config.User,
		config.DbName,
		config.Password,
		config.Port,
	)

	dbCfg, err := pgxpool.ParseConfig(uri)
	if err != nil {
		return nil, err
	}

	dbCfg.MaxConns = int32(config.PoolSize)
	dbCfg.MaxConnLifetime = time.Duration(config.ConnLifeTimeSecond) * time.Second

	conn, err := pgxpool.ConnectConfig(context.Background(), dbCfg)
	if err != nil {
		return nil, err
	}

	_, err = conn.Exec(context.Background(), CreateTable)

	if err != nil {
		return nil, err
	}

	return &DatabaseImpl{
		conn: conn,
	}, nil
}

func (d *DatabaseImpl) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	return d.conn.Exec(ctx, sql, args...)
}

func (d *DatabaseImpl) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return d.conn.QueryRow(ctx, sql, args...)
}

func (d *DatabaseImpl) Query(ctx context.Context, sql string) (pgx.Rows, error) {
	return d.conn.Query(ctx, sql)
}

func Handle(err error) error {
	if strings.Contains(err.Error(), "ch_positive_usd") {
		return model.NotEnoughFiat{}
	} else if strings.Contains(err.Error(), "ch_positive_btc") {
		return model.NotEnoughBtc{}
	} else if strings.Contains(err.Error(), "ch_positive_eth") {
		return model.NotEnoughETH{}
	} else {
		return model.SqlInternalError{}
	}
}
