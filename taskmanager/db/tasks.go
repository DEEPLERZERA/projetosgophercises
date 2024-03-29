package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks") // bucket name
var db *bolt.DB // database

type Task struct { // Task struct
	Key   int
	Value string
}

func Init(dbPath string) error { // initialize database
	var err error 
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second}) // open database
	if err != nil { 		// if error is not nil
		return err          // return error
	}
	return db.Update(func(tx *bolt.Tx) error { // update database
		_, err := tx.CreateBucketIfNotExists(taskBucket) // create bucket if not exists
		return err // return error
	})
}

func CreateTask(task string) (int, error) { // create task
	var id int
	err := db.Update(func(tx *bolt.Tx) error { // update database
		b := tx.Bucket(taskBucket) // get bucket
		id64, _ := b.NextSequence() // get next sequence
		id = int(id64) // convert to int
		key := itob(id) // convert to byte
		return b.Put(key, []byte(task)) // put task
	})
	if err != nil { // if error is not nil
		return -1, err // return error
	}
	return id, nil // return id
}

func AllTasks() ([]Task, error) { // get all tasks
	var tasks []Task // make task slice
	err := db.View(func(tx *bolt.Tx) error { // view database
		b := tx.Bucket(taskBucket) // get bucket
		c := b.Cursor() // get cursor
		for k, v := c.First(); k != nil; k, v = c.Next() { // loop through cursor
			tasks = append(tasks, Task{ // append task
				Key:   btoi(k), // convert to int
				Value: string(v), // convert to string
			})
		}
		return nil // return nil
	})
	if err != nil { // if error is not nil
		return nil, err // return error
	}
	return tasks, nil // return tasks
}

func DeleteTask(key int) error { // delete task
	return db.Update(func(tx *bolt.Tx) error { // update database
		b := tx.Bucket(taskBucket) // get bucket
		return b.Delete(itob(key)) // delete task
	})
}

func itob(v int) []byte { // convert int to byte
	b := make([]byte, 8) // make byte with length 8
	binary.BigEndian.PutUint64(b, uint64(v)) // put uint64
	return b // return byte
}


func btoi (b []byte) int { // convert byte to int
	return int(binary.BigEndian.Uint64(b)) // return int
}