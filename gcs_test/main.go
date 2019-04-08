// Sample storage-quickstart creates a Google Cloud Storage bucket.
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func list(w io.Writer, client *storage.Client, bucket string) ([]string, error) {
	ctx := context.Background()
	it := client.Bucket(bucket).Objects(ctx, nil)
	objects := []string{}

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		fmt.Fprintln(w, attrs.Name)

		objects = append(objects, attrs.Name)
	}

	return objects, nil
}

func main() {
	ctx := context.Background()

	// projectID := "gcpdev"

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	bucket := "dev-cloudsql-test"
	objects, _ := list(os.Stdout, client, bucket)
	fmt.Println(objects)
}
