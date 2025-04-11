package db

import (
	"github.com/boltdb/bolt"
	"time"
	"encoding/binary"
)

var db = *bolt.DB
var taskBucket = []byte("tasks")

type Task struct {
	Key int
	Value string
} 

func Init(dbPath string) err error {
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil{
		return err
	}
	return db.Update(func (tx *bolt.Tx) error {
		_, err := tx.CreateIfNotExists(taskBucket)
	}) 
}

// Adds task to our db Bucket

func AddTask(task string) (id int, error) {
	err := db.Update(func (tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		return b.Put(intToByte(id64), []byte(task))
	})

	if err := nil {
		return -1, err
	}
	return id, nil
}

func intToByte(i int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

func byteToInt(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
