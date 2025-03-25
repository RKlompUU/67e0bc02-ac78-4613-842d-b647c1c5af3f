package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite
	e *echo.Echo
}

func (suite *HandlerTestSuite) SetupTest() {
	suite.e = echo.New()
}

func (suite *HandlerTestSuite) TestHealthzMetrics() {
	cases := []struct {
		name         string
		path         string
		mockMetrics  Metrics
		expectStatus int
		expect       string
	}{
		{
			name:         "OK",
			path:         "/healthz",
			mockMetrics:  newMockMetrics(withPingReturns(func() error { return nil })),
			expectStatus: http.StatusOK,
			expect:       `"OK"`,
		},
		{
			name:         "Unavailable",
			path:         "/healthz",
			mockMetrics:  newMockMetrics(withPingReturns(func() error { return errors.New("connection refused") })),
			expectStatus: http.StatusServiceUnavailable,
			expect:       `{"error":"error pinging database: connection refused"}`,
		},
	}
	for _, tc := range cases {
		suite.T().Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tc.path, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := suite.e.NewContext(req, rec)

			if assert.NoError(suite.T(), Healthz(tc.mockMetrics)(c)) {
				assert.Equal(suite.T(), tc.expectStatus, rec.Code)
				assert.Equal(suite.T(), tc.expect, strings.TrimRight(rec.Body.String(), "\n"))
			}
		})
	}
}

func (suite *HandlerTestSuite) TestTransferFeesHourly() {
	cases := []struct {
		name         string
		path         string
		mockMetrics  Metrics
		expectStatus int
		expect       string
	}{
		{
			name: "OK",
			path: "/metrics/ethereum/transaction_fees_hourly?date=2020-09-07",
			mockMetrics: newMockMetrics(
				withGetEOATransactionFeesHourlyReturns(func() ([]*EOATransactionFeesHourlyResult, error) {
					return []*EOATransactionFeesHourlyResult{{
						Hour: time.Unix(0, 0),
						Fees: 123,
					}}, nil
				}),
			),
			expectStatus: http.StatusOK,
			expect:       fmt.Sprintf(`[{"t": %d, "v": 123}]`, time.Unix(0, 0).Unix()),
		},
		{
			name: "Metrics.GetEOATransactionFeesHourlyErrors",
			path: "/metrics/ethereum/transaction_fees_hourly?date=2020-09-07",
			mockMetrics: newMockMetrics(
				withGetEOATransactionFeesHourlyReturns(func() ([]*EOATransactionFeesHourlyResult, error) {
					return nil, errors.New("error retrieving metric: GetEOATransactionFeesHourly")
				}),
			),
			expectStatus: http.StatusInternalServerError,
			expect:       `{"error": "error retrieving metric: GetEOATransactionFeesHourly"}`,
		},
		{
			name: "DateNotFound",
			path: "/metrics/ethereum/transaction_fees_hourly?date=2020-01-01",
			mockMetrics: newMockMetrics(
				withGetEOATransactionFeesHourlyReturns(func() ([]*EOATransactionFeesHourlyResult, error) {
					return nil, nil
				}),
			),
			expectStatus: http.StatusNotFound,
			expect:       `{"error": "no metrics found for date: 2020-01-01"}`,
		},
		{
			name:         "MissingDateParameter",
			path:         "/metrics/ethereum/transaction_fees_hourly",
			mockMetrics:  newMockMetrics(),
			expectStatus: http.StatusBadRequest,
			expect:       `{"error":"query parameter 'date' is required in format YYYY-MM-DD"}`,
		},
		{
			name:         "MalformedDateFormatWithTimestamp",
			path:         "/metrics/ethereum/transaction_fees_hourly?date=2020-01-01T00:00:00+00:00",
			mockMetrics:  newMockMetrics(),
			expectStatus: http.StatusBadRequest,
			expect:       `{"error":"query parameter 'date' is required in format YYYY-MM-DD"}`,
		},
		{
			name:         "MalformedDateFormatPeriodSeparator",
			path:         "/metrics/ethereum/transaction_fees_hourly?date=2020.01.01",
			mockMetrics:  newMockMetrics(),
			expectStatus: http.StatusBadRequest,
			expect:       `{"error":"query parameter 'date' is required in format YYYY-MM-DD"}`,
		},
		{
			name:         "WrongParameter",
			path:         "/metrics/ethereum/transaction_fees_hourly?time=2020-01-01",
			mockMetrics:  newMockMetrics(),
			expectStatus: http.StatusBadRequest,
			expect:       `{"error":"query parameter 'date' is required in format YYYY-MM-DD"}`,
		},
	}
	for _, tc := range cases {
		suite.T().Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tc.path, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := suite.e.NewContext(req, rec)

			if assert.NoError(suite.T(), TransferFeesHourly(tc.mockMetrics)(c)) {
				assert.Equal(suite.T(), tc.expectStatus, rec.Code)
				assert.JSONEq(suite.T(), tc.expect, rec.Body.String())
			}
		})
	}
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
