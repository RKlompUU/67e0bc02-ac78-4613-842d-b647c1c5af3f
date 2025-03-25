package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// dateFormat specifies the YYYY-MM-DD format
const dateFormat = "2006-01-02"

type ResponseError struct {
	Error string `json:"error"`
}

// TransferFeesHourlyResponse is the response wrapper type around HourlyTransferFees for the handler TransferFeesHourly
type TransferFeesHourlyResponse []*HourlyTransferFees

type HourlyTransferFees struct {
	Hour int64   `json:"t"`
	Fees float64 `json:"v"`
}

func Healthz(m Metrics) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := m.Ping(context.Background())
		if err != nil {
			return c.JSON(http.StatusServiceUnavailable,
				&ResponseError{fmt.Sprintf("error pinging database: %s", err)})
		}
		return c.JSON(http.StatusOK, "OK")
	}
}

// TransferFeesHourly calculates all the sum of EOA transfer fees broken down by hour for the given date.
// If the date parameter is missing or improperly formatted, a http.StatusBadRequest is returned and the
// handler exits prematurely.
func TransferFeesHourly(m Metrics) echo.HandlerFunc {
	return func(c echo.Context) error {
		date, err := time.Parse(dateFormat, c.QueryParam("date"))
		if err != nil {
			return c.JSON(http.StatusBadRequest,
				&ResponseError{"query parameter 'date' is required in format YYYY-MM-DD"},
			)
		}

		fees, err := m.GetEOATransactionFeesHourly(date)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				&ResponseError{"error retrieving metric: GetEOATransactionFeesHourly"},
			)
		}
		if fees == nil {
			return c.JSON(http.StatusNotFound,
				&ResponseError{fmt.Sprintf("no metrics found for date: %s", date.Format(dateFormat))},
			)
		}

		var resp TransferFeesHourlyResponse
		for _, f := range fees {
			resp = append(resp, &HourlyTransferFees{
				Hour: f.Hour.Unix(),
				Fees: f.Fees,
			})
		}

		return c.JSON(http.StatusOK, resp)
	}
}
