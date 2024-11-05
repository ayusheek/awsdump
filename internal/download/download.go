package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// DownloadBucket downloads a file from the specified URL and saves it to the provided file path
func DownloadBucket(url, destinationPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download %s: %v", url, err)
	}
	defer resp.Body.Close()

	// Get the directory path and ensure it exists
	dirPath := filepath.Dir(destinationPath)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directories for path %s: %v", dirPath, err)
	}

	// Create the file
	out, err := os.Create(destinationPath)
	if err != nil {
		return fmt.Errorf("failed to create file at %s: %v", destinationPath, err)
	}
	defer out.Close()

	// Write the response body to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %v", destinationPath, err)
	}

	fmt.Printf("Downloaded: %s\n", url)
	return nil
}
