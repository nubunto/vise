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
	"github.com/rs/cors"

	"github.com/nubunto/vise/api"
	"github.com/nubunto/vise/destroyer"
)

type Hosts map[string]http.Handler

func (h Hosts) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, ok := h[r.Host]; ok {
		handler.ServeHTTP(w, r)
	} else {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

func dbInit(db *bolt.DB) {
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("users"))
		if err != nil {
			return errors.New("Fail to create bucket \"users\".")
		}
		return nil
	})
}

func main() {
	hosts := make(Hosts)

	apiEndpoint := echo.New()
	apiEndpoint.SetLogPrefix("api")
	apiEndpoint.Use(mw.Logger())
	apiEndpoint.Use(mw.Recover())

	db, err := bolt.Open("data/vise.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	dbInit(db)

	hosts["api.localhost:8080"] = apiEndpoint

	api.RegisterHandlers(apiEndpoint, db)

	website := echo.New()
	website.SetLogPrefix("site")
	website.Use(mw.Logger())
	website.Use(mw.Recover())

	hosts["localhost:8080"] = website

	website.Static("/", "public")
	website.Get("/:token/:file/download", func(c *echo.Context) error {
		token, file := c.P(0), c.P(1)
		filePath := path.Join(api.UploadedPath, token, file)

		src, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer src.Close()

		res := c.Response()
		res.Header().Set("Content-Disposition", "attachment")

		if _, err := io.Copy(res.Writer(), src); err != nil {
			return err
		}
		return nil
	})
	destroyer.Scan(db)
	log.Fatal(http.ListenAndServe(":8080", cors.Default().Handler(hosts)))
}
