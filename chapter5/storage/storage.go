package storage

import "context"

// Item represents an item at a shop
type Item struct {
	Name  string
	Price int64
}

// Storage is our storage interface
// We'll implement it with:
// Mongo and Bolt
// @collection string is a collection of Item
// in SQL it will be a table
// in Bolt it will be a bucket
// In Mongo it will be a collection
type Storage interface {
	GetByName(ctx context.Context, collection string, key string) (interface{}, error)
	Put(ctx context.Context, collection string, key string, item interface{})
}
