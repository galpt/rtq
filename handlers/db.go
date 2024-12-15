package handlers

import (
	"log"

	badger "github.com/dgraph-io/badger/v4"
)

var (
	opts = badger.DefaultOptions("./db").WithIndexCacheSize(100 << 20)
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func openDB() *badger.DB {
	db, err := badger.Open(opts)
	handleError(err)
	return db
}

func closeDB(db *badger.DB) {
	db.Close()
}

func writeToDB(key string, value string) error {

	db := openDB()
	defer closeDB(db)

	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(value))
		return err
	})

	return err
}

func getFromDB(key string) ([]byte, error) {

	db := openDB()
	defer closeDB(db)

	var valCopy []byte
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		valCopy, err = item.ValueCopy(nil)
		handleError(err)
		return nil
	})
	handleError(err)

	return valCopy, nil
}
