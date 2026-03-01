package files

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func SaveFile(fileInfo *multipart.FileHeader) (*string, error) {
	src, err := fileInfo.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	extension := filepath.Ext(fileInfo.Filename)
	newUUID := uuid.New().String()

	secureFilename := fmt.Sprintf("%s%s", newUUID, extension)

	dstPath := filepath.Join("uploads", secureFilename)

	dst, err := os.Create(dstPath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

	return &secureFilename, nil
}

func GetFileInfoFromExtraction(key string, extractionFiles map[string][]*multipart.FileHeader) (*multipart.FileHeader, error) {
	files, ok := extractionFiles[key]

	if !ok || len(files) == 0 {
		return nil, nil
	}

	return files[0], nil
}

func RemoveFile(fileName string) error {
	return os.Remove(filepath.Join("uploads", fileName))
}
