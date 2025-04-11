package db

import (
	"github.com/boltdb/bolt"
	"time"
	"encoding/binary"
)

var db *bolt.DB
var taskBucket = []byte("tasks")

type Task struct {
	Key int
	Value string
} 

func Init(dbPath string) (err error) {
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil{
		return err
	}
	return db.Update(func (tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	}) 
}

// Adds task to our db Bucket

func AddTask(task string) (id int, err error) {
	err = db.Update(func (tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		return b.Put(intToByte(id), []byte(task))
	})

	if err != nil {
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

// Deletes our current bucket removing all tasks

func Clear() error {
	return db.Update(func (tx *bolt.Tx) error {
		return tx.DeleteBucket(taskBucket)
	})
}

// Returns all tasks to be listed to user 

func GetTasks() (tasks []Task, err error) {
	err = db.View(func (tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)

		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next(){
			tasks = append(tasks, Task{byteToInt(k), string(v[:])})
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// Removes tasks from db 

func DoTasks(keys []int) error {
	return db.Update(func (tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)

		for _, key := range keys {
			bucket.Delete(intToByte(key))	
		}
		return nil
	})
} 








