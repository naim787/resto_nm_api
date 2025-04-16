package repository

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

func OpenDB() *leveldb.DB {
	var err error
	db, err = leveldb.OpenFile("../internal/db", nil);
	if err != nil {
		log.Fatal("Failed to open the database :", err)
		return nil
	}

	return db
}

func RedUsers() ([]byte, error) {
	return db.Get([]byte("users"), nil)
}

func SaveUsers(data []byte) error {
	return db.Put([]byte("users"), data, nil)
}

func DeleteUsers() error {
	err := db.Delete([]byte("users"), nil)
	if err != nil {
		return err
	}

	return nil
}