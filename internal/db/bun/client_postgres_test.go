//go:build postgres

package bun

import (
	"context"
	"github.com/feditools/go-lib/mock"
	"github.com/feditools/login/internal/config"
	"github.com/feditools/login/internal/db"
	"github.com/feditools/login/internal/models/testdata"
	"github.com/spf13/viper"
	"testing"
)

func TestNew_Postgres(t *testing.T) {
	dbAddress := "postgres"
	dbDatabase := "test"
	dbPassword := "test"
	dbPort := 5432
	dbTLSMode := dbTLSModeDisable
	dbUser := "test"

	viper.Reset()

	viper.Set(config.Keys.DBType, "postgres")

	viper.Set(config.Keys.DBAddress, dbAddress)
	viper.Set(config.Keys.DBDatabase, dbDatabase)
	viper.Set(config.Keys.DBPassword, dbPassword)
	viper.Set(config.Keys.DBPort, dbPort)
	viper.Set(config.Keys.DBTLSMode, dbTLSMode)
	viper.Set(config.Keys.DBUser, dbUser)

	metricsCollector, _ := mock.NewMetricsCollector()

	bun, err := New(context.Background(), metricsCollector)
	if err != nil {
		t.Errorf("unexpected error initializing bun connection: %s", err.Error())
		return
	}
	if bun == nil {
		t.Errorf("client is nil")
		return
	}
}

func TestPgConn(t *testing.T) {
	dbAddress := "postgres"
	dbDatabase := "test"
	dbPassword := "test"
	dbPort := 5432
	dbTLSMode := dbTLSModeDisable
	dbUser := "test"

	viper.Reset()

	viper.Set(config.Keys.DBType, "postgres")

	viper.Set(config.Keys.DBAddress, dbAddress)
	viper.Set(config.Keys.DBDatabase, dbDatabase)
	viper.Set(config.Keys.DBPassword, dbPassword)
	viper.Set(config.Keys.DBPort, dbPort)
	viper.Set(config.Keys.DBTLSMode, dbTLSMode)
	viper.Set(config.Keys.DBUser, dbUser)

	bun, err := pgConn(context.Background())
	if err != nil {
		t.Errorf("unexpected error initializing pg connection: %s", err.Error())
		return
	}
	if bun == nil {
		t.Errorf("client is nil")
		return
	}
}

func testNewPostresClient() (db.DB, error) {
	viper.Reset()

	viper.Set(config.Keys.DBType, "postgres")

	viper.Set(config.Keys.DBAddress, "postgres")
	viper.Set(config.Keys.DBDatabase, "test")
	viper.Set(config.Keys.DBPassword, "test")
	viper.Set(config.Keys.DBPort, 5432)
	viper.Set(config.Keys.DBUser, "test")
	viper.Set(config.Keys.DBEncryptionKey, testdata.TestEncryptionKey)

	metricsCollector, _ := mock.NewMetricsCollector()

	client, err := New(context.Background(), metricsCollector)
	if err != nil {
		return nil, err
	}

	err = client.DoMigration(context.Background())
	if err != nil {
		return nil, err
	}

	err = client.LoadTestData(context.Background())
	if err != nil {
		return nil, err
	}

	return client, nil
}
