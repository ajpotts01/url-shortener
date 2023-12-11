package database

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

type DB interface {
	CreateUrl(ctx context.Context, urlParams CreateUrlParams)
	FetchUrl(ctx context.Context, urlKey string)
}

type Database struct {
	DbClient *firestore.Client
}

func (db *Database) InitConnection(ctx context.Context) *firestore.Client {
	return &firestore.Client{}
}

func CreateConnection(ctx context.Context, projectId string, databaseId string) *firestore.Client {
	client, err := firestore.NewClientWithDatabase(ctx, projectId, databaseId)
	//client, err := firestore.NewClient(ctx, projectId)

	if err != nil {
		log.Fatalf("Failed to init db: %v", err)
	}

	return client
}
