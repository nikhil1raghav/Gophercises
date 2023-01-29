package db

import (
	"encoding/binary"
	"github.com/boltdb/bolt"
	"time"
)
var taskbucket = []byte("tasks")
var db *bolt.DB
type Task struct{
	Key int
	Value string
}
func Init(dbPath string)error{
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1*time.Second})
	if err!=nil{
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err:=tx.CreateBucketIfNotExists(taskbucket)
		return err
	})
}

func CreateTask(task string) (int, error){
	var id int
	err:=db.Update(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket(taskbucket)
		id64, _ :=bucket.NextSequence()
		id = int(id64)
		id64b := itob(id)
		return bucket.Put(id64b, []byte(task))
	})
	if err!=nil{
		return -1, err
	}
	return id, nil
}
func AllTasks() ([]Task, error){
	var tasks []Task
	err:=db.View(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket(taskbucket)
		cursor:=bucket.Cursor()
		for k,v:=cursor.First();k!=nil;k,v =cursor.Next(){
			tasks=append(tasks, Task{
				btoi(k),
				string(v),
			})
		}
		return nil
	})
	if err!=nil{
		return nil, err
	}
	return tasks, nil
}
func DeleteTask(key int) error{
	err:=db.Update(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket(taskbucket)
		return bucket.Delete(itob(key))
	})
	return err
}
func itob(v int) []byte{
	b:=make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
