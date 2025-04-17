package repository

import (
    "github.com/syndtr/goleveldb/leveldb"
)

// variabel global database leveldb
var db *leveldb.DB

// membukan database leveldb
func OpenDB() (*leveldb.DB, error) {
    var err error
    db, err = leveldb.OpenFile("../internal/db", nil)
    if err != nil {
        return nil, err
    }
    return db, nil
}

// membaca data dari database 
func RedUsers() ([]byte, error) {
    if db == nil {
        return nil, leveldb.ErrClosed
    }
    return db.Get([]byte("users"), nil)
}

// meyimpan data ke database
func SaveUsers(data []byte) error {
    if db == nil {
        return leveldb.ErrClosed
    }
    return db.Put([]byte("users"), data, nil)
}

// meghapus data dari database
func DeleteUsers() error {
    if db == nil {
        return leveldb.ErrClosed
    }
    return db.Delete([]byte("users"), nil)
}