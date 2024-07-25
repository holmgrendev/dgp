package database

import (
	"os"

	_ "github.com/lib/pq"
)

//var db *sql.DB

type DatabaseCredentials struct {
	Host     string
	Database string
	User     string
	Password string
}

func InitDatabase() DatabaseCredentials {
	// Set default credentials
	dbc := DatabaseCredentials{Host: "localhost", Database: os.Getenv("POSTGRES_DB"), User: "postgres", Password: os.Getenv("POSTGRES_PASSWORD")}

	// Check user
	if os.Getenv("POSTGRES_USER") != "" {
		dbc.User = os.Getenv("POSTGRES_USER")
	}

	// Check database
	if dbc.Database == "" {
		dbc.Database = dbc.User
	}

	// Check host
	if os.Getenv("POSTGRES_HOST") != "" {
		dbc.Host = os.Getenv("POSTGRES_HOST")
	}

	return dbc
}
