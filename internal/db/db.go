package db

import (
	"context"
	"github.com/feditools/login/internal/models"
)

// DB represents a database client
type DB interface {
	// Close closes the db connections
	Close(ctx context.Context) Error
	// Create stores the object
	Create(ctx context.Context, i any) Error
	// DoMigration runs database migrations
	DoMigration(ctx context.Context) Error
	// LoadTestData adds test data to the database
	LoadTestData(ctx context.Context) Error
	// ReadByID returns a model by its ID
	ReadByID(ctx context.Context, id int64, i any) Error
	// Update updates stored data
	Update(ctx context.Context, i any) Error

	// FediAccount

	// CreateFediAccount stores the federated instance and caches it
	CreateFediAccount(ctx context.Context, account *models.FediAccount) (err Error)
	// ReadFediAccount returns one federated social account
	ReadFediAccount(ctx context.Context, id int64) (account *models.FediAccount, err Error)
	// ReadFediAccountByUsername returns one federated social account
	ReadFediAccountByUsername(ctx context.Context, instanceID int64, username string) (account *models.FediAccount, err Error)
	// UpdateFediAccount updates the stored federated instance and caches it
	UpdateFediAccount(ctx context.Context, account *models.FediAccount) (err Error)

	// FediInstance

	// CreateFediInstance stores the federated instance and caches it
	CreateFediInstance(ctx context.Context, instance *models.FediInstance) (err Error)
	// ReadFediInstance returns one federated social instance
	ReadFediInstance(ctx context.Context, id int64) (instance *models.FediInstance, err Error)
	// ReadFediInstanceByDomain returns one federated social instance
	ReadFediInstanceByDomain(ctx context.Context, domain string) (instance *models.FediInstance, err Error)
	// UpdateFediInstance updates the stored federated instance and caches it
	UpdateFediInstance(ctx context.Context, instance *models.FediInstance) (err Error)

	// OauthClient

	// ReadOauthClient returns one oauth client
	ReadOauthClient(ctx context.Context, id int64) (instance *models.OauthClient, err Error)
}
