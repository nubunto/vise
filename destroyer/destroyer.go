package destroyer

import (
	"encoding/json"
	"os"
	"path"

	"log"

	"github.com/boltdb/bolt"
	"github.com/robfig/cron"

	"time"

	"github.com/nubunto/vise/api"
)

func Scan(db *bolt.DB) {
	c := cron.New()
	c.AddFunc("@daily", func() {
		err := db.Update(func(tx *bolt.Tx) error {
			users := tx.Bucket([]byte("users"))
			c := users.Cursor()

			for token, fileList := c.First(); token != nil; token, fileList = c.Next() {
				var files api.UserFiles
				err := json.Unmarshal(fileList, &files)
				if err != nil {
					log.Println("Strange file list for user", token, ". Aborting file deletion (for this user only")
					continue
				}
				newFiles := make([]api.FileInfo, 0)
				directoryPath := path.Join(api.UploadedPath, string(token))
				for _, file := range files.Files {
					pathToFile := path.Join(directoryPath, file.Filename)
					fileInfo, err := os.Lstat(pathToFile)
					if err != nil {
						log.Println("Strange file info for user", token, ". Aborting.")
						continue
					}

					fileExistenceTime, availability := fileInfo.ModTime(), file.DaysAvailable
					log.Println("File was created at", fileExistenceTime, "and his availability is", availability)
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
		if err != nil {
			log.Println("Failed to run today's daily scanner:", err)
		}
	})
	c.Start()
}
