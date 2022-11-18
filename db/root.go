package db

import (
	"github.com/boltdb/bolt"
	"strconv"
	"time"
)

var Db *bolt.DB

func OpenDB() error {
	db, err := bolt.Open("tasks.db", 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		return err
	}

	Db = db
	createBucket("tasks")
	createBucket("completed")

	return nil
}

func AddToBucket(bucketName string, task string) error {
	err := Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))

		_ = b.Put([]byte(strconv.Itoa(bucketLength(b)+1)), []byte(task))

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func RemoveFromBucket(bucketName string, task string) {}

func ListBucketItems(bucketName string) map[string]string {
	items := make(map[string]string)

	_ = Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			items[string(k[:])] = string(v)
		}

		return nil
	})

	return items
}

func bucketLength(b *bolt.Bucket) int {
	c := b.Cursor()
	var sum int

	for k, _ := c.First(); k != nil; k, _ = c.Next() {
		sum++
	}

	return sum
}

func createBucket(name string) *bolt.Bucket {
	var bucket *bolt.Bucket

	_ = Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(name))
		if b == nil {
			b, _ := tx.CreateBucket([]byte(name))
			bucket = b
		} else {
			bucket = b
		}

		return nil
	})

	return bucket
}

func CloseDB() error {
	err := Db.Close()
	if err != nil {
		return err
	}

	return nil
}
