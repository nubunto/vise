package persistence

import (
	"errors"
	"log"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/nubunto/vise/persistence/types"
	"github.com/nubunto/vise/uppath"
)

var db *bolt.DB

func Init() error {
	var err error
	db, err = bolt.Open("data/vise.db", 0600, nil)
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("files"))
		if err != nil {
			return errors.New("Failed to create bucket \"files\".")
		}
		return nil
	})
}

func Close() {
	db.Close()
}

func CheckExpiration() error {
	return db.Update(func(tx *bolt.Tx) error {
		files := tx.Bucket([]byte("files"))
		filesCursor := files.Cursor()
		for fileToken, _ := filesCursor.First(); fileToken != nil; fileToken, _ = filesCursor.Next() {
			fileBucket := files.Bucket(fileToken)
			// TODO: check if bucket exists.
			created, err := time.Parse(time.UnixDate, string(fileBucket.Get([]byte("created-in"))))
			if err != nil {
				log.Println("Time parse error:", err)
				continue
			}
			availability, err := strconv.Atoi(string(fileBucket.Get([]byte("expires-in"))))
			if err != nil {
				log.Println("Days are in a strange format:", err)
				continue
			}
			if created.Day()-time.Now().Day() >= availability {
				filename := string(fileBucket.Get([]byte("filename")))
				err = os.Remove(path.Join(string(fileToken), filename))
				if err != nil {
					log.Println("Failed to remove file:", filename, " -", err)
					continue
				}
				fileBucket.Delete(fileToken)
			}
		}
		return nil
	})
}

func GetLinks(match func(string) bool, uriBuilder func(string) string) ([]string, error) {
	links := make([]string, 0)
	err := db.View(func(tx *bolt.Tx) error {
		files := tx.Bucket([]byte("files"))
		return files.ForEach(func(fileToken, _ []byte) error {
			fileBucket := files.Bucket(fileToken)
			userToken := string(fileBucket.Get([]byte("user-token")))
			if match(string(userToken)) {
				links = append(links, uriBuilder(string(fileToken)))
			}
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return links, nil
}

func FindFile(token []byte) (*os.File, error) {
	var file *os.File
	var fileErr error
	err := db.View(func(tx *bolt.Tx) error {
		files := tx.Bucket([]byte("files"))
		c := files.Cursor()
		for ftoken, _ := c.First(); ftoken != nil; ftoken, _ = c.Next() {
			fileBucket := files.Bucket(ftoken)
			if string(ftoken) == string(token) {
				filename := fileBucket.Get([]byte("filename"))
				file, fileErr = os.Open(path.Join(uppath.UploadedPath, string(token), string(filename)))
				return nil
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return file, fileErr
}

func Save(f types.File) error {
	return f.Save(db)
}
