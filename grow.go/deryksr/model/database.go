package model

import (
	"errors"
	"strconv"
	"strings"
	"sync"
)

type LocalDatabase struct {
	data sync.Map
	size sync.Map
}

func NewDatabase() *LocalDatabase {
	var database LocalDatabase
	database.size.Store("size", 0)
	return &database
}

func GenerateKey(record Growth) string {
	return strconv.Itoa(record.Year) +
		strings.ToLower(record.Indicator) +
		strings.ToLower(record.Country)
}

func (db *LocalDatabase) updateCounter(value int) {
	count, _ := db.size.Load("size")
	db.size.Store("size", count.(int)+value)
}

func (db *LocalDatabase) Save(record Growth) {
	key := GenerateKey(record)
	_, loaded := db.data.LoadOrStore(key, record)
	if !loaded {
		db.updateCounter(1)
	}
}

func (db *LocalDatabase) Read(key string) *Growth {
	value, ok := db.data.Load(key)
	if ok {
		return &Growth{
			Country:   value.(Growth).Country,
			Indicator: value.(Growth).Indicator,
			Value:     value.(Growth).Value,
			Year:      value.(Growth).Year,
		}
	}
	return nil
}

func (db *LocalDatabase) Delete(key string) error {
	_, ok := db.data.Load(key)
	if !ok {
		return errors.New("the key <" + key + "> has not been found!")
	}

	db.data.Delete(key)
	db.updateCounter(-1)
	return nil
}

func (db *LocalDatabase) Upsert(record Growth) {
	key := GenerateKey(record)
	_, ok := db.data.Load(key)
	if !ok {
		db.updateCounter(1)
	}
	db.data.Store(key, record)
}

func (db *LocalDatabase) Size() int {
	count, _ := db.size.Load("size")
	return count.(int)
}
