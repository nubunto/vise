package api

import (
	"io"
	"net/http"
	"strconv"

	"code.google.com/p/go-uuid/uuid"
	"github.com/labstack/echo"

	"github.com/nubunto/vise/persistence"
	"github.com/nubunto/vise/persistence/types"
	"github.com/nubunto/vise/uppath"
)

func SaveFile(c *echo.Context) error {
	req := c.Request()
	req.ParseMultipartForm(16 << 20)

	days, err := strconv.Atoi(c.Form("days"))
	if err != nil {
		return err
	}

	srcFile, fh, err := req.FormFile("file")
	if err != nil {
		return err
	}
	defer srcFile.Close()

	userToken := c.Form("user-token")
	if len(userToken) == 0 {
		userToken = uuid.New()
	}
	fileToken := uuid.New()
	boltFile := types.File{
		UserToken:     userToken,
		FileToken:     fileToken,
		DaysAvailable: days,
		Filename:      fh.Filename,
	}
	err = uppath.UploadFile(boltFile, srcFile)
	if err != nil {
		return err
	}

	err = persistence.Save(boltFile)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, FileUploadResponse{ResponseOK, userToken})
}

func uriBuilder(c *echo.Context) func(string) string {
	return func(token string) string {
		return c.Echo().URI(DownloadFile, token)
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

func getTokenLinks(match func(string) bool, uriBuilder func(string) string) (*LinksResponse, error) {
	links, err := persistence.GetLinks(match, uriBuilder)
	if err != nil {
		return nil, err
	}
	return &LinksResponse{ResponseOK, links}, nil
}

func DownloadFile(c *echo.Context) error {
	fileToken := c.P(0)

	src, err := persistence.FindFile([]byte(fileToken))
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
