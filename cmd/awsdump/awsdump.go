package main

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/ayusheek/awsdump/internal/bucket"
	"github.com/ayusheek/awsdump/internal/download"
	"github.com/ayusheek/awsdump/internal/runner"
)

func scrapeBucket(bucketURL string) {
	keys, err := bucket.ExtractBucketKeys(bucketURL)
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, key := range keys {
		fmt.Printf("%s%s\n", bucketURL, key)
	}
}

func scrapeAndDownloadBucket(bucketURL, outputDir string, maxThreads int) {
	keys, err := bucket.ExtractBucketKeys(bucketURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	semaphore := make(chan struct{}, maxThreads)
	var wg sync.WaitGroup

	for _, key := range keys {
		wg.Add(1)
		go func(key string) {
			defer wg.Done()

			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			fileURL := fmt.Sprintf("%s%s", bucketURL, key)
			err := download.DownloadBucket(fileURL, filepath.Join(outputDir, key))
			if err != nil {
				fmt.Println("Error:", err)
			}
		}(key)
	}

	wg.Wait()
}

func main() {
	bucketURL, outputDir, booleanScrape, maxDlThreads := runner.ParseArgs()

	if booleanScrape {
		scrapeBucket(bucketURL)
		return
	} else {
		scrapeAndDownloadBucket(bucketURL, outputDir, maxDlThreads)
	}
}
