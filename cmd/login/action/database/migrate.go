package database

import (
	"context"
	"github.com/feditools/login/cmd/login/action"
	"github.com/feditools/login/internal/config"
	"github.com/feditools/login/internal/db/bun"
	"github.com/feditools/login/internal/metrics/statsd"
	"github.com/spf13/viper"
)

// Migrate runs database migrations
var Migrate action.Action = func(ctx context.Context) error {
	l := logger.WithField("func", "Migrate")

	// create metrics collector
	metricsCollector, err := statsd.New()
	if err != nil {
		l.Errorf("metrics: %s", err.Error())
		return err
	}
	defer func() {
		err := metricsCollector.Close()
		if err != nil {
			l.Errorf("closing metrics: %s", err.Error())
		}
	}()

	// create database client
	l.Info("running database migration")
	dbClient, err := bun.New(ctx, metricsCollector)
	if err != nil {
		l.Errorf("db: %s", err.Error())
		return err
	}
	defer func() {
		err := dbClient.Close(ctx)
		if err != nil {
			l.Errorf("closing db: %s", err.Error())
		}
	}()

	err = dbClient.DoMigration(ctx)
	if err != nil {
		l.Errorf("migration: %s", err.Error())
		return err
	}

	if viper.GetBool(config.Keys.DbLoadTestData) {
		err = dbClient.LoadTestData(ctx)
		if err != nil {
			l.Errorf("migration: %s", err.Error())
			return err
		}
	}

	return nil
}
