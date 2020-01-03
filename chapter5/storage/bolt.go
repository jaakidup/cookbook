package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

// BoltStorage conforms to Storage interface
type BoltStorage struct {
	DB          *bolt.DB
	DBName      string
	saveChannel chan SaveObject
}

// SaveObject is sent over the SaveChannel
// to be saved in database
type SaveObject struct {
	Collection string
	Key        string
	Object     interface{}
}

// GetByName an item with key
func (bs *BoltStorage) GetByName(ctx context.Context, collection string, key string) (interface{}, error) {
	fmt.Println("GET from", collection, "Item with key:", key)

	var object interface{}

	// bs.DB.View(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("testcollection"))

	// 	c := b.Cursor()

	// 	for k, v := c.First(); k != nil; k, v = c.Next() {
	// 		fmt.Printf("key=%s, value=%s\n", k, v)
	// 	}

	// 	return nil
	// })

	fn := func(tx *bolt.Tx) error {
		fmt.Println("Inside bolt view")
		b := tx.Bucket([]byte(collection))
		v := b.Get([]byte(key))
		// fmt.Println("v ", v)
		json.Unmarshal(v, &object)

		return nil
	}
	// // Access data from within a read-only transactional block.
	if err := bs.DB.View(fn); err != nil {
		fmt.Println("Failed getting bucket and key combo")
		// log.Fatal(err)
		return nil, err
	}
	// fmt.Println("OK, value found, but need to send it back")
	return object, nil
}

// Put item into storage
func (bs *BoltStorage) Put(ctx context.Context, collection string, key string, item interface{}) {
	// fmt.Println("PUT into", collection, "Item:", item, " with Key: ", key)
	bs.saveChannel <- SaveObject{Collection: collection, Key: key, Object: item}
}

// ******************************************************************* //
//
// TODO
// Bolt needs to run in goroutines
// Write can only handle one write connection at a time
//
// ******************************************************************* //

// NewBoltStorage creates a new bolt db with filename
// @name is the filename to use for the bolt database
func NewBoltStorage(name string) *BoltStorage {
	db, err := bolt.Open(name+".db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	// readerdb, err := bolt.Open(name+".db", 0666, &bolt.Options{Timeout: 1 * time.Second, ReadOnly: true})
	// //readerdb, err := bolt.Open(name+".db", 0666, &bolt.Options{ReadOnly: true})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// readerdb, err := bolt.Open(name+".db", 0600, &bolt.Options{ReadOnly: true})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// create saveChannel with no buffer for TESTING Only
	// TODO: Test to see what size buffer is needed
	saveChannel := make(chan SaveObject, 0)

	// TODO
	// bucket name
	// key
	// object to store
	go func() {
		for {
			saveObject := <-saveChannel

			// db.View(func(tx *bolt.Tx) error {
			// 	b := tx.Bucket([]byte(saveObject.Collection))

			// 	c := b.Cursor()

			// 	for k, v := c.First(); k != nil; k, v = c.Next() {
			// 		fmt.Printf("key=%s, value=%s\n", k, v)
			// 	}

			// 	return nil
			// })

			fn := func(tx *bolt.Tx) error {
				b, err := tx.CreateBucketIfNotExists([]byte(saveObject.Collection))
				//b := tx.Bucket([]byte(saveObject.Collection))
				object, err := json.Marshal(saveObject.Object)
				if err != nil {
					fmt.Println(err)
				}
				err = b.Put([]byte(saveObject.Key), object)
				return nil
			}

			if err = db.Update(fn); err != nil {
				fmt.Println(err)
			}

		}
	}()

	bs := &BoltStorage{
		DB:          db,
		DBName:      name,
		saveChannel: saveChannel,
	}
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	// defer db.Close()

	return bs
}
