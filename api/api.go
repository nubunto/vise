package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"

	"code.google.com/p/go-uuid/uuid"
	"github.com/boltdb/bolt"
	"github.com/labstack/echo"
)

var UploadedPath = "uploaded/"
var db *bolt.DB

func DB(d *bolt.DB) {
	db = d
}

func SaveFile(c *echo.Context) error {
	token := c.Form("user-token")
	if len(token) == 0 {
		token = uuid.New()
	}

	req := c.Request()
	req.ParseMultipartForm(16 << 20)

	days, err := strconv.Atoi(c.Form("days"))
	if err != nil {
		return err
	}

	file, fh, err := req.FormFile("file")
	if err != nil {
		fmt.Println("formfile err", err)
		return err
	}
	defer file.Close()

	directoryPath := path.Join(UploadedPath, token)
	err = os.MkdirAll(directoryPath, 0755)
	if err != nil {
		fmt.Println("mkdir err", err)
		return err
	}

	filePath := path.Join(directoryPath, fh.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		fmt.Println("dst create err", err)
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		token := []byte(token)

		b := tx.Bucket([]byte("users"))

		userFiles := b.Get(token)

		var files UserFiles
		if userFiles == nil {
			files = UserFiles{
				Files: make([]FileInfo, 0),
			}
		} else {
			err := json.Unmarshal(userFiles, &files)
			if err != nil {
				return err
			}
		}

		files.Files = append(files.Files, FileInfo{fh.Filename, days})
		marshaled, err := json.Marshal(files)
		if err != nil {
			return err
		}
		b.Put(token, marshaled)
		return nil
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, FileUploadResponse{ResponseOK, token})
}

func GetLinks(c *echo.Context) error {
	links := make([]Link, 0)
	err := db.View(func(tx *bolt.Tx) error {
		users := tx.Bucket([]byte("users"))
		c := users.Cursor()

		for token, fileInfo := c.First(); token != nil; token, fileInfo = c.Next() {
			var files UserFiles
			err := json.Unmarshal(fileInfo, &files)
			if err != nil {
				return err
			}
			for _, f := range files.Files {
				links = append(links, Link{
					FileInfo: f,
					URL:      string(token) + "/" + f.Filename + "/download",
				})
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, LinksResponse{ResponseOK, links})
}

func GetTokenLinks(c *echo.Context) error {
	userToken := c.P(0)
	links := make([]Link, 0)
	err := db.View(func(tx *bolt.Tx) error {
		users := tx.Bucket([]byte("users"))
		c := users.Cursor()

		for token, fileInfo := c.First(); token != nil; token, fileInfo = c.Next() {
			if string(token) == userToken {
				var files UserFiles
				err := json.Unmarshal(fileInfo, &files)
				if err != nil {
					return err
				}
				for _, f := range files.Files {
					links = append(links, Link{
						FileInfo: f,
						URL:      string(token) + "/" + f.Filename + "/download",
					})
				}
				break
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, LinksResponse{ResponseOK, links})
}

func DBStats(c *echo.Context) error {
	return c.JSON(http.StatusOK, db.Stats())
}
