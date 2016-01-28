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
			log.Println("Scanning file token", string(fileToken))

			possibleDeletionDate := time.Now()

			var deleteDate time.Time
			err := deleteDate.UnmarshalBinary(fileBucket.Get([]byte("delete-date")))
			if err != nil {
				log.Println("Delete date unmarshal error:", err)
				continue
			}
			d, m, y := possibleDeletionDate.Date()
			dd, mm, yy := deleteDate.Date()
			if d == dd && m == mm && y == yy {
				log.Println("Removing file...")
				err = os.RemoveAll(path.Join(uppath.UploadedPath, string(fileToken)))
				if err != nil {
					log.Println("Failed to remove file:", fileToken, " -", err)
					continue
				}
				files.DeleteBucket(fileToken)
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

func DbStats() bolt.Stats {
	return db.Stats()
}

func Inspect() ([]interface{}, error) {
	ret := make([]interface{}, 0)
	err := db.View(func(tx *bolt.Tx) error {
		files := tx.Bucket([]byte("files"))
		files.ForEach(func(fileToken, _ []byte) error {
			fileInfo := files.Bucket(fileToken)
			days, _ := strconv.Atoi(string(fileInfo.Get([]byte("expires-in"))))
			var creationTime, deleteDate time.Time
			var err error
			err = creationTime.UnmarshalBinary(fileInfo.Get([]byte("creation-time")))
			if err != nil {
				log.Println("Created date unmarshal error:", err)
			}
			err = deleteDate.UnmarshalBinary(fileInfo.Get([]byte("delete-date")))
			if err != nil {
				log.Println("Delete date unmarshal error:", err)
			}
			data := struct {
				Days       int
				Filename   string
				Created    string
				DeleteDate string
			}{
				days,
				string(fileInfo.Get([]byte("filename"))),
				creationTime.String(),
				deleteDate.String(),
			}
			ret = append(ret, data)
			return nil
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return ret, nil
}
