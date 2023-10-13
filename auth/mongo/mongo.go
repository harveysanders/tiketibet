package mongo

import (
	"net/url"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	db     *mongo.Client
	dbName string
}

func NewStore(client *mongo.Client, dbName string) *Store {
	return &Store{
		db:     client,
		dbName: dbName,
	}
}

// ParseDBName parses the database name from a connection string.
func ParseDBName(connectionString string) (string, error) {
	u, err := url.Parse(connectionString)
	if err != nil {
		return "", err
	}
	return strings.TrimPrefix(u.Path, "/"), nil
}
