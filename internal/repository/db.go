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

/////////////////////// function /////////////////////////////////////////////

func ReadDB(key string) ([]byte, error) {
    if db == nil {
        return nil, leveldb.ErrClosed
    }
    return db.Get([]byte(key), nil)
}



/////////////////////// USERS /////////////////////////////////////////////

// meyimpan data ke database
func SaveUsers(data []byte, key string) error {
    if db == nil {
        return leveldb.ErrClosed
    }
    return db.Put([]byte(key), data, nil)
}

// meghapus data dari database
func DeleteUsers() error {
    if db == nil {
        return leveldb.ErrClosed
    }
    return db.Delete([]byte("users"), nil)
}




/////////////////////// PRODUCT /////////////////////////////////////////////
