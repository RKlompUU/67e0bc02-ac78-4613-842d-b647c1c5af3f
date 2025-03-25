package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewServer creates the Metrics API Server
// The following routes are available:
//   /healthz
//   /metrics/ethereum/transaction_fees_hourly?date=YYYY-MM-DD
func NewServer(m Metrics) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/healthz", Healthz(m))
	e.GET("/metrics/ethereum/transaction_fees_hourly", TransferFeesHourly(m))

	return e
}
