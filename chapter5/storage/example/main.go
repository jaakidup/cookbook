package main

import (
	"context"
	"fmt"

	"github.com/jaakidup/go-cookbook/chapter5/storage"
)

func main() {
	storage := storage.NewBoltStorage("sample")
	Logic(storage)
}

// Logic does a few steps against the storage interface
func Logic(store storage.Storage) {
	ctx := context.Background()

	collection := "testcollection"

	item := storage.Item{Name: "Test Item", Price: 5000}
	store.Put(ctx, collection, "one", item)

	next := storage.Item{Name: "Another Item", Price: 2345}
	store.Put(ctx, collection, "two", next)

	fmt.Println("==========================================")
	fmt.Println("Let's do request")
	fmt.Println("")

	storedItem, err := store.GetByName(ctx, collection, "one")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(storedItem)
	storedItem, err = store.GetByName(ctx, collection, "two")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(storedItem)
}
