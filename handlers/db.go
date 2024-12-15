package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	badger "github.com/dgraph-io/badger/v4"
)

var (
	opts       = badger.DefaultOptions("./db").WithIndexCacheSize(100 << 20)
	maxAntrian = 8
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func openDB() *badger.DB {
	db, err := badger.Open(opts)
	handleError(err)
	initConfig(db) // initialize config for max number of queue per day.
	return db
}

func closeDB(db *badger.DB) {
	db.Close()
}

func initConfig(db *badger.DB) {
	err := db.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte("config:maxantrian"))
		if err != badger.ErrKeyNotFound {
			return nil
		}

		txn.Set([]byte("config:maxantrian"), []byte(fmt.Sprintf("%v", maxAntrian)))
		return nil
	})
	handleError(err)
}

func writeToDB(key string, value interface{}) error {
	db := openDB()
	defer closeDB(db)

	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), jsonData)
		return err
	})

	return err
}

func getFromDB(key string) (interface{}, error) {
	db := openDB()
	defer closeDB(db)
	var valCopy []byte
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		valCopy, err = item.ValueCopy(nil)
		handleError(err)

		return nil
	})

	if err != nil {
		return nil, err
	}

	var data interface{}
	err = json.Unmarshal(valCopy, &data)
	handleError(err)

	return data, nil
}

func getAllKeysWithPrefix(prefix string) (map[string]interface{}, error) {
	db := openDB()
	defer closeDB(db)

	result := make(map[string]interface{})

	err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = true
		it := txn.NewIterator(opts)
		defer it.Close()

		prefixBytes := []byte(prefix)
		for it.Seek(prefixBytes); it.ValidForPrefix(prefixBytes); it.Next() {
			item := it.Item()
			key := string(item.Key())
			valCopy, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}
			var data interface{}
			err = json.Unmarshal(valCopy, &data)
			handleError(err)
			result[key] = data

		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
