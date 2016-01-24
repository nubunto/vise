package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"

	"github.com/nubunto/vise/api"
	"github.com/nubunto/vise/destroyer"
	"github.com/nubunto/vise/persistence"
	"github.com/nubunto/vise/uppath"
	"github.com/nubunto/vise/website"
)

func main() {
	uppath.UploadedPath = "uploaded/"
	err := uppath.EnsureDirectories()
	if err != nil {
		log.Fatal(err)
	}

	defer persistence.Close()

	vise := echo.New()

	vise.SetLogPrefix("vise")
	vise.Use(mw.Logger())
	vise.Use(mw.Recover())

	apiEndpoint := vise.Group("/api")
	apiEndpoint.Post("/save", api.SaveFile)
	apiEndpoint.Get("/links", api.GetLinks)
	apiEndpoint.Get("/:token/links", api.GetTokenLinks)

	vise.Use(mw.Logger())
	vise.Use(mw.Recover())

	vise.Index("public/index.html")
	vise.Get("/download/:token/:file", website.DownloadFile)

	destroyer.Scan()

	log.Fatal(http.ListenAndServe(":8080", vise))
}
