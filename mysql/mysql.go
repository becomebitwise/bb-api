package mysql

import (
	"context"
	"errors"

	api "github.com/becomebitwise/bb-api"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // need to import for sqlx
	"github.com/rs/xid"
)

// A Client is a connection to a mysql table that knows how to query for becomebitwise data.
type Client struct {
	db *sqlx.DB
}

// New creates a client for becomebitwise data.
func New(dsn string) (Client, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return Client{}, err
	}

	return Client{db: db}, nil
}

// Authenticate determines if a set of credentials are correct.
func (c Client) Authenticate(ctx context.Context, creds api.Creds) (string, error) {
	var dbCreds api.Creds
	if err := c.db.GetContext(ctx, &dbCreds, "select email, password from users where email = $1", creds.Email); err != nil {
		return "", err
	}

	if dbCreds == creds {
		return "", errors.New("passwords don't match")
	}

	return "", nil
}

// CreateUser adds a User to the database with the given, hashed password.
func (c Client) CreateUser(ctx context.Context, user api.User, password string) (api.Identifier, error) {
	user.ID = xid.New().Bytes()
	if _, err := c.db.ExecContext(ctx, "insert into users (id, email, password, first_name, last_name) values ($1, $2, $3, $4)", user.ID, user.Email, password, user.FirstName, user.LastName); err != nil {
		return nil, err
	}

	return user.ID, nil
}

// GetUser retrieves a User by ID or email.
func (c Client) GetUser(ctx context.Context, id string) (api.User, error) {
	// TODO (erik): Make this function handle emails properly.
	var user api.User
	if err := c.db.GetContext(ctx, &user, "select * from users where id = ?", id); err != nil {
		return user, err
	}

	return user, nil
}
