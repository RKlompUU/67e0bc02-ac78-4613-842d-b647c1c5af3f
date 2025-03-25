package main

import (
	"context"
	"time"
)

type mockDatabase struct {
	*Database
	pingReturns                        func() error
	getEOATransactionFeesHourlyReturns func() ([]*EOATransactionFeesHourlyResult, error)
}

func (m *mockDatabase) GetEOATransactionFeesHourly(
	date time.Time,
) ([]*EOATransactionFeesHourlyResult, error) {
	return m.getEOATransactionFeesHourlyReturns()
}

func (m *mockDatabase) Ping(ctx context.Context) error {
	return m.pingReturns()
}

type mockMetricsOption func(*mockDatabase)

func withPingReturns(f func() error) mockMetricsOption {
	return func(m *mockDatabase) {
		m.pingReturns = f
	}
}

func withGetEOATransactionFeesHourlyReturns(f func() ([]*EOATransactionFeesHourlyResult, error)) mockMetricsOption {
	return func(m *mockDatabase) {
		m.getEOATransactionFeesHourlyReturns = f
	}
}

func newMockMetrics(opts ...mockMetricsOption) *mockDatabase {
	m := &mockDatabase{}

	for _, opt := range opts {
		opt(m)
	}

	return m
}
