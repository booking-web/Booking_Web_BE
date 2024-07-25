package pkg

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func SaveImage(file multipart.File, fileName string, role string) error {
	path := fmt.Sprintf("./assets/%s/", role)

	// Create the file
	dst, err := os.Create(filepath.Join(path, fileName))
	if err != nil {
		return fmt.Errorf("error saving file: %v", err)
	}
	defer dst.Close()

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		return fmt.Errorf("error saving file: %v", err)
	}

	return nil
}
