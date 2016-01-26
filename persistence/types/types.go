package types

import (
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

// New API
type File struct {
	UserToken     string
	FileToken     string
	Filename      string
	DaysAvailable int
}

func (file File) Save(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		filesBucket := tx.Bucket([]byte("files"))

		fileBucket, err := filesBucket.CreateBucketIfNotExists([]byte(file.FileToken))
		if err != nil {
			return err
		}

		err = fileBucket.Put([]byte("user-token"), []byte(file.UserToken))
		err = fileBucket.Put([]byte("filename"), []byte(file.Filename))
		err = fileBucket.Put([]byte("expires-in"), []byte(strconv.Itoa(file.DaysAvailable)))
		err = fileBucket.Put([]byte("creation-time"), []byte(time.Now().String()))
		if err != nil {
			return err
		}
		return nil
	})
}
