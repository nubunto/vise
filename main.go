package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"

	"github.com/nubunto/vise/api"
	"github.com/nubunto/vise/destroyer"
	"github.com/nubunto/vise/persistence"
	"github.com/nubunto/vise/uppath"
	"github.com/nubunto/vise/website"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Serve temporary and unimportant files"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "port, p",
			Value: 8080,
			Usage: "The port which to listen",
		},
	}
	app.Action = func(c *cli.Context) {
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

		website.LoadBox("public")
		vise.Get("/", website.Index)
		vise.Get("/download/:token/:file", website.DownloadFile)

		destroyer.Scan()

		port := strconv.Itoa(c.Int("port"))
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal(http.ListenAndServe(":"+port, vise))
	}
	app.Run(os.Args)
}
