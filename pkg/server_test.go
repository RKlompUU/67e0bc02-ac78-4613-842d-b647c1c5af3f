package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServerTestSuite struct {
	suite.Suite
	e *echo.Echo
}

func (suite *ServerTestSuite) SetupTest() {
	e := NewServer(
		newMockMetrics(
			withPingReturns(func() error {
				return nil
			}),
			withGetEOATransactionFeesHourlyReturns(func() ([]*EOATransactionFeesHourlyResult, error) {
				return nil, nil
			}),
		))

	suite.e = e
}

func (suite *ServerTestSuite) TestHealthzRouteConfigured() {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)

	rec := httptest.NewRecorder()
	suite.e.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusOK, rec.Code)
}

func (suite *ServerTestSuite) TestMetricEthereumTransactionFeesHourlyRouteConfigured() {
	req := httptest.NewRequest(http.MethodGet, "/metrics/ethereum/transaction_fees_hourly?date=2020-01-01", nil)

	rec := httptest.NewRecorder()
	suite.e.ServeHTTP(rec, req)

	// 404 will be the default response when no data is availble
	assert.Equal(suite.T(), http.StatusNotFound, rec.Code)
}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}
