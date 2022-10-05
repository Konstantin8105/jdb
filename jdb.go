// Package is tiny JSON database
package jdb

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

type Db[T any] struct {
	mutex    sync.Mutex
	filename string
	vs       []T
}

func Open[T any](filename string) (_ *Db[T], err error) {
	db := Db[T]{}
	db.filename = filename

	var dat []byte
	dat, err = os.ReadFile(filename)
	if err != nil {
		// ignore error
		err = nil
	} else {
		err = json.Unmarshal(dat, &db.vs)
		if err != nil {
			return
		}
	}
	return &db, err
}

func (db *Db[T]) Add(v T) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.vs = append(db.vs, v)
	return
}

func (db *Db[T]) Get() []T {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	return db.vs
}

func (db *Db[T]) Write() (err error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	b, err := json.MarshalIndent(&db.vs, "", "\t")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(db.filename, b, 0644); err != nil {
		return err
	}

	return
}
