package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"errors"

	"github.com/boltdb/bolt"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"

	"github.com/nubunto/vise/api"
	"github.com/nubunto/vise/destroyer"
)

func dbInit(db *bolt.DB) {
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("users"))
		if err != nil {
			return errors.New("Fail to create bucket \"users\".")
		}
		return nil
	})
}

func ensureDirectories(dirs ...string) error {
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return nil
}

func dbOpen(dir string) *bolt.DB {
	db, err := bolt.Open("data/vise.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	err := ensureDirectories("data", "uploaded")
	if err != nil {
		log.Fatal(err)
	}

	vise := echo.New()

	vise.SetLogPrefix("vise")
	vise.Use(mw.Logger())
	vise.Use(mw.Recover())

	db := dbOpen("data")
	defer db.Close()
	dbInit(db)

	api.DB(db)
	apiEndpoint := vise.Group("/api")
	apiEndpoint.Post("/save", api.SaveFile)
	apiEndpoint.Get("/links", api.GetLinks)
	apiEndpoint.Get("/:token/links", api.GetTokenLinks)

	vise.Use(mw.Logger())
	vise.Use(mw.Recover())

	vise.Index("public/index.html")
	vise.Get("/:token/:file/download", func(c *echo.Context) error {
		token, file := c.P(0), c.P(1)
		filePath := path.Join(api.UploadedPath, token, file)

		src, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer src.Close()

		res := c.Response()
		if _, err := io.Copy(res.Writer(), src); err != nil {
			return err
		}
		return nil
	})
	destroyer.Scan(db)
	log.Fatal(http.ListenAndServe(":8080", vise))
}
