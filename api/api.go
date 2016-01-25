package api

import (
	"io"
	"net/http"
	"os"
	"path"
	"strconv"

	"code.google.com/p/go-uuid/uuid"
	"github.com/labstack/echo"

	"github.com/nubunto/vise/persistence"
	"github.com/nubunto/vise/persistence/types"
	"github.com/nubunto/vise/uppath"
)

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
		return err
	}
	defer file.Close()

	err = uppath.UploadFile(file, fh, token)
	if err != nil {
		return err
	}

	err = persistence.UpdateFiles(types.FileInfo{fh.Filename, days}, token)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, FileUploadResponse{ResponseOK, token})
}

func uriBuilder(c *echo.Context) func(string, string) string {
	return func(token, filename string) string {
		return c.Echo().URI(DownloadFile, token, filename)
	}
}

func GetLinks(c *echo.Context) error {
	matchAll := func(string) bool {
		return true
	}
	linksResponse, err := getTokenLinks(matchAll, uriBuilder(c))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, linksResponse)
}

func GetTokenLinks(c *echo.Context) error {
	userToken := c.P(0)
	matchSpecific := func(b string) bool {
		return b == userToken
	}
	linksResponse, err := getTokenLinks(matchSpecific, uriBuilder(c))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, linksResponse)
}

func getTokenLinks(match func(string) bool, uriBuilder func(string, string) string) (*LinksResponse, error) {
	links, err := persistence.GetLinks(match, uriBuilder)
	if err != nil {
		return nil, err
	}
	return &LinksResponse{ResponseOK, links}, nil
}

func DownloadFile(c *echo.Context) error {
	token, file := c.P(0), c.P(1)
	filePath := path.Join(uppath.UploadedPath, token, file)

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
}
