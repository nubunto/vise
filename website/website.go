package website

import (
	"github.com/labstack/echo"
	"github.com/nubunto/vise/uppath"
	"io"
	"os"
	"path"
)

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
