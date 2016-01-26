package uppath

import (
	"io"
	"mime/multipart"
	"os"
	"path"

	"github.com/nubunto/vise/persistence/types"
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

func UploadFile(file types.File, src multipart.File) error {
	directoryPath := path.Join(UploadedPath, file.FileToken)
	err := os.MkdirAll(directoryPath, 0755)
	if err != nil {
		return err
	}

	filePath := path.Join(directoryPath, file.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}
