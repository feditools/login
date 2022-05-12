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
	// ResetCache clears any caches in the module
	ResetCache(ctx context.Context) Error
	// Update updates stored data
	Update(ctx context.Context, i any) Error

	// ApplicationToken

	// CountApplicationTokens returns the number of application tokens
	CountApplicationTokens(ctx context.Context) (count int64, err Error)
	// CreateApplicationToken stores the application token
	CreateApplicationToken(ctx context.Context, applicationToken *models.ApplicationToken) (err Error)
	// ReadApplicationToken returns one application token
	ReadApplicationToken(ctx context.Context, id int64) (applicationToken *models.ApplicationToken, err Error)
	// ReadApplicationTokenByToken returns one application token
	ReadApplicationTokenByToken(ctx context.Context, token string) (applicationToken *models.ApplicationToken, err Error)
	// ReadApplicationTokensPage returns a page of application tokens
	ReadApplicationTokensPage(ctx context.Context, index, count int) (applicationToken []*models.ApplicationToken, err Error)
	// UpdateApplicationToken updates the stored application token
	UpdateApplicationToken(ctx context.Context, applicationToken *models.ApplicationToken) (err Error)

	// FediAccount

	// CountFediAccounts returns the number of federated social account
	CountFediAccounts(ctx context.Context) (count int64, err Error)
	// CountFediAccountsForInstance returns the number of federated social account for an instance
	CountFediAccountsForInstance(ctx context.Context, instanceID int64) (count int64, err Error)
	// CreateFediAccount stores the federated social account
	CreateFediAccount(ctx context.Context, account *models.FediAccount) (err Error)
	// IncFediAccountLoginCount updates the login count of a stored federated instance
	IncFediAccountLoginCount(ctx context.Context, account *models.FediAccount) (err Error)
	// ReadFediAccount returns one federated social account
	ReadFediAccount(ctx context.Context, id int64) (account *models.FediAccount, err Error)
	// ReadFediAccountByUsername returns one federated social account
	ReadFediAccountByUsername(ctx context.Context, instanceID int64, username string) (account *models.FediAccount, err Error)
	// ReadFediAccountsPage returns a page of federated social accounts
	ReadFediAccountsPage(ctx context.Context, index, count int) (instances []*models.FediAccount, err Error)
	// UpdateFediAccount updates the stored federated instance
	UpdateFediAccount(ctx context.Context, account *models.FediAccount) (err Error)

	// FediInstance

	// CountFediInstances returns the number of federated instances
	CountFediInstances(ctx context.Context) (count int64, err Error)
	// CreateFediInstance stores the federated instance
	CreateFediInstance(ctx context.Context, instance *models.FediInstance) (err Error)
	// ReadFediInstance returns one federated social instance
	ReadFediInstance(ctx context.Context, id int64) (instance *models.FediInstance, err Error)
	// ReadFediInstanceByDomain returns one federated social instance
	ReadFediInstanceByDomain(ctx context.Context, domain string) (instance *models.FediInstance, err Error)
	// ReadFediInstancesPage returns a page of federated social instances
	ReadFediInstancesPage(ctx context.Context, index, count int) (instances []*models.FediInstance, err Error)
	// UpdateFediInstance updates the stored federated instance
	UpdateFediInstance(ctx context.Context, instance *models.FediInstance) (err Error)

	// OauthClient

	// CountOauthClients returns the number of oauth clients
	CountOauthClients(ctx context.Context) (count int64, err Error)
	// CreateOauthClient stores the oauth client
	CreateOauthClient(ctx context.Context, client *models.OauthClient) (err Error)
	// ReadOauthClient returns one oauth client
	ReadOauthClient(ctx context.Context, id int64) (client *models.OauthClient, err Error)
	// ReadOauthClientsPage returns a page of oauth clients
	ReadOauthClientsPage(ctx context.Context, index, count int) (clients []*models.OauthClient, err Error)
	// UpdateOauthClient updates the stored oauth client
	UpdateOauthClient(ctx context.Context, client *models.OauthClient) (err Error)
}
