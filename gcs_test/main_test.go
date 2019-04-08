package main

import (
	"context"
	"log"
	"os"
	"testing"

	"cloud.google.com/go/storage"
)

func TestList(t *testing.T) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	bucket := "dev-cloudsql-test"
	objects, err := list(os.Stdout, client, bucket)

	if err != nil {
		t.Fatal("error: ", err)
	}

	if len(objects) != 2 {
		t.Fatal("objects size not equal")
	}
}
