package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func SaveUploadedFile(file multipart.File, handler *multipart.FileHeader, uploadDir string) (string, error) {
	os.MkdirAll(uploadDir, os.ModePerm)

	ext := filepath.Ext(handler.Filename)

	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	filePath := filepath.Join(uploadDir, filename)

	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("/%s/%s", uploadDir, filename), nil
}
