package db

import (
	"github.com/boltdb/bolt"
	"github.com/mitchellh/go-homedir"
	"os"
	"strconv"
	"time"
)

var Db *bolt.DB

func OpenDB() error {
	dir, _ := homedir.Dir()

	if _, err := os.Stat(dir + "/.tasks"); os.IsNotExist(err) {
		_ = os.Mkdir(dir+"/.tasks", 0766)
	}

	db, err := bolt.Open(dir+"/.tasks/tasks.db", 0600, &bolt.Options{Timeout: 1 * time.Second})

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

		k, _ := b.NextSequence()

		_ = b.Put([]byte(strconv.Itoa(int(k))), []byte(task))

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func RemoveFromBucket(bucketName string, taskNumber string) error {
	task, _ := strconv.Atoi(taskNumber)

	err := Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))

		if bucketName == "tasks" {
			c := b.Cursor()
			current := 1

			completed := tx.Bucket([]byte("completed"))

			for k, _ := c.First(); current <= task; k, _ = c.Next() {
				if current == task {
					t := b.Get(k)
					_ = completed.Put(k, t)

					err := b.Delete(k)
					if err != nil {
						return err
					}
				}

				current++
			}

			return nil
		} else {
			return nil
		}
	})

	if err != nil {
		return err
	}

	return nil
}

func DeleteFromBucket(taskNumber string) error {
	task, _ := strconv.Atoi(taskNumber)

	err := Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))

		c := b.Cursor()
		current := 1

		for k, _ := c.First(); current <= task; k, _ = c.Next() {
			if current == task {
				err := b.Delete(k)
				if err != nil {
					return err
				}
			}

			current++
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func ListBucketItems(bucketName string) map[string]string {
	items := make(map[string]string)

	_ = Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			items[string(k)] = string(v)
		}

		return nil
	})

	return items
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

//func bucketLength(b *bolt.Bucket) int {
//	c := b.Cursor()
//	var sum int
//
//	for k, _ := c.First(); k != nil; k, _ = c.Next() {
//		sum++
//	}
//
//	return sum
//}
