package persistence

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path"
	"time"

	"github.com/boltdb/bolt"
	"github.com/nubunto/vise/persistence/types"
	"github.com/nubunto/vise/uppath"
)

var db *bolt.DB

func init() {
	var err error
	db, err = bolt.Open("data/vise.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("users"))
		if err != nil {
			return errors.New("Fail to create bucket \"users\".")
		}
		return nil
	})
}

func Close() {
	db.Close()
}

func CheckExpiration() error {
	return db.Update(func(tx *bolt.Tx) error {
		users := tx.Bucket([]byte("users"))
		c := users.Cursor()

		for token, fileList := c.First(); token != nil; token, fileList = c.Next() {
			var files types.UserFiles
			err := json.Unmarshal(fileList, &files)
			if err != nil {
				log.Println("Strange file list for user", token, ". Aborting file deletion (for this user only")
				continue
			}
			newFiles := make([]types.FileInfo, 0)
			directoryPath := path.Join(uppath.UploadedPath, string(token))
			for _, file := range files.Files {
				pathToFile := path.Join(directoryPath, file.Filename)
				fileInfo, err := os.Lstat(pathToFile)
				if err != nil {
					log.Println("Strange file info for user", token, ". Aborting.")
					continue
				}

				fileExistenceTime, availability := fileInfo.ModTime(), file.DaysAvailable
				possibleDeleteDay := time.Now().Day()
				fileExistenceDay := fileExistenceTime.Day()
				if possibleDeleteDay-fileExistenceDay >= availability {
					log.Println("Removing file", pathToFile)
					// should remove file.
					// update database
					// etc.
					err = os.Remove(pathToFile)
					if err != nil {
						log.Println("Failed to remove file ", pathToFile)
					}
				} else {
					// he gets to live another day.
					newFiles = append(newFiles, file)
				}
			}
			if len(newFiles) == 0 {
				users.Delete(token)
				os.RemoveAll(directoryPath)
			} else {
				files.Files = newFiles
				newData, err := json.Marshal(files)
				if err != nil {
					log.Println("Failed to encode (json) files:", files)
				}
				users.Put(token, newData)
			}
		}
		return nil
	})

}

func UpdateFiles(fi types.FileInfo, token string) error {
	return db.Update(func(tx *bolt.Tx) error {
		token := []byte(token)

		b := tx.Bucket([]byte("users"))

		userFiles := b.Get(token)

		var files types.UserFiles
		if userFiles == nil {
			files = types.UserFiles{
				Files: make([]types.FileInfo, 0),
			}
		} else {
			err := json.Unmarshal(userFiles, &files)
			if err != nil {
				return err
			}
		}

		files.Files = append(files.Files, fi)
		marshaled, err := json.Marshal(files)
		if err != nil {
			return err
		}
		b.Put(token, marshaled)
		return nil
	})
}

func GetLinks(match func(string) bool, uriBuilder func(string, string) string) ([]types.Link, error) {
	links := make([]types.Link, 0)
	err := db.View(func(tx *bolt.Tx) error {
		users := tx.Bucket([]byte("users"))
		return users.ForEach(func(token, fileInfo []byte) error {
			if match(string(token)) {
				var files types.UserFiles
				err := json.Unmarshal(fileInfo, &files)
				if err != nil {
					return err
				}
				for _, f := range files.Files {
					links = append(links, types.Link{
						FileInfo: f,
						URL:      uriBuilder(string(token), f.Filename),
					})
				}
			}
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return links, nil
}
