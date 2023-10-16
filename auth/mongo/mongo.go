package mongo

import (
	"context"
	"net/url"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID        string    `bson:"_id"`
	Email     string    `bson:"email" validate:"required,email"`
	Password  string    `bson:"password" validate:"required,min=6,max=30"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

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

func (s *Store) DB() *mongo.Database {
	return s.db.Database(s.dbName)
}

func (s *Store) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	res := s.DB().Collection("users").FindOne(ctx, bson.M{"email": email})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, res.Err()
	}

	var user *User
	err := res.Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
