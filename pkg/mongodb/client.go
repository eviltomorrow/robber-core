package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//
var (
	DSN     string
	MaxOpen uint64 = 10
	DB      *mongo.Client
)

//
var (
	DefaultTimeout = 10 * time.Second
)

func Build() {
	pool, err := build(DSN)
	if err != nil {
		panic(err)
	}
	DB = pool
}

func Close() error {
	if DB == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	return DB.Disconnect(ctx)
}

func build(dsn string) (*mongo.Client, error) {
	client, err := mongo.NewClient(
		options.Client().ApplyURI(dsn).SetMaxPoolSize(MaxOpen),
	)
	if err != nil {
		return nil, err
	}

	ctx1, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	if err := client.Connect(ctx1); err != nil {
		return nil, err
	}

	ctx2, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	if err := client.Ping(ctx2, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}
