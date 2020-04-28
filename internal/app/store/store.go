package store

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

//
type Store struct {
	config *Config
	db     *mongo.Database
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {

	// Base context.
	ctx := context.Background()

	clientOpts := options.Client().ApplyURI(s.config.DatabaseURL)
	client, err := mongo.Connect(ctx, clientOpts)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	db = client.Database(s.config.DataBaseName)

	fmt.Println("Connected to MongoDB!")

	return nil
}

func (s *Store) Close() {
	// ...

	s.db.Client().Disconnect(context.TODO())

}

func GetStoreDB() *mongo.Database {
	return db
}
