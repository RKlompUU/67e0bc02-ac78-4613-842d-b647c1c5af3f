package main

import (
	"context"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

// Metrics provides the interface for querying various metrics
type Metrics interface {
	Ping(ctx context.Context) error
	GetEOATransactionFeesHourly(date time.Time) ([]*EOATransactionFeesHourlyResult, error)
}

type Database struct {
	*pgxpool.Pool
}

type EOATransactionFeesHourlyResult struct {
	Hour time.Time `db:"hour"`
	Fees float64   `db:"fees"`
}

func NewDatabase(ctx context.Context, connString string) *Database {
	pool, _ := pgxpool.Connect(ctx, connString)
	return &Database{Pool: pool}
}

func (db *Database) GetEOATransactionFeesHourly(date time.Time) ([]*EOATransactionFeesHourlyResult, error) {
	query := `
		SELECT date_trunc('hour', t.block_time) AS "hour", sum(t.gas_used * t.gas_price) / 1E18 AS "fees"
		FROM transactions t
		LEFT JOIN contracts c ON c.address = t.to
		WHERE t.to != '0x0000000000000000000000000000000000000000'
		AND date_trunc('day', t.block_time) = $1
		AND c.address IS NULL
		GROUP BY date_trunc('hour', t.block_time);
	`
	var fees []*EOATransactionFeesHourlyResult

	if err := pgxscan.Select(
		context.Background(),
		db,
		&fees,
		query,
		date,
	); err != nil {
		log.Errorf("error querying GetEOATransactionFeesHourly: ", err)
		return nil, err
	}

	return fees, nil
}
