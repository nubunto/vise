package uppath

import (
	"io"
	"mime/multipart"
	"os"
	"path"
)

var UploadedPath string

func EnsureDirectories(dirs ...string) error {
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return nil
}

func UploadFile(file multipart.File, fh *multipart.FileHeader, token string) error {
	directoryPath := path.Join(UploadedPath, token)
	err := os.MkdirAll(directoryPath, 0755)
	if err != nil {
		return err
	}

	filePath := path.Join(directoryPath, fh.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return err
	}
	return nil
}
