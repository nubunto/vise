package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/labstack/echo/engine/standard"

	"github.com/nubunto/vise/api"
	"github.com/nubunto/vise/destroyer"
	"github.com/nubunto/vise/persistence"
	"github.com/nubunto/vise/uppath"
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
		cli.StringFlag{
			Name:  "upload-path, u",
			Value: "uploaded",
			Usage: "The directory on which to save uploaded data",
		},
	}
	app.Action = func(c *cli.Context) {
		uppath.UploadedPath = c.String("upload-path")
		err := uppath.EnsureDirectories("data", "uploaded")
		if err != nil {
			log.Fatal(err)
		}

		err = persistence.Init()
		if err != nil {
			log.Fatal(err)
		}
		defer persistence.Close()

		vise := echo.New()

		vise.SetLogPrefix("vise")
		vise.Use(mw.Logger())
		vise.Use(mw.Recover())

		apiEndpoint := vise.Group("/api")
		apiEndpoint.POST("/save", api.SaveFile)
		apiEndpoint.GET("/links", api.GetLinks)
		apiEndpoint.GET("/:token/links", api.GetTokenLinks)
		apiEndpoint.GET("/download/:file", api.DownloadFile)

		private := vise.Group("/private")
		private.Use(mw.BasicAuth(func(usr, pwd string) bool {
			return usr == os.Getenv("VISE_USER") && pwd == os.Getenv("VISE_PWD")
		}))
		private.GET("/stats", func(c echo.Context) error {
			return c.JSON(http.StatusOK, persistence.DbStats())
		})
		private.GET("/inspect", func(c echo.Context) error {
			d, err := persistence.Inspect()
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, d)
		})

		vise.Use(mw.Logger())
		vise.Use(mw.Recover())

		//assets := http.FileServer(assetFS())
		//vise.GET("/", assets)
		//vise.Get("/static/*", assets)

		destroyer.Scan()

		port := strconv.Itoa(c.Int("port"))
		if err != nil {
			log.Fatal(err)
		}
		//log.Fatal(http.ListenAndServe(":"+port, vise))
		vise.Run(standard.New(":" + port))

	}
	app.Run(os.Args)
}
