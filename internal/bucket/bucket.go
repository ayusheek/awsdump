package bucket

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// ExtractBucketKeys fetches and extracts keys from the given bucket URL
func ExtractBucketKeys(bucketURL string) ([]string, error) {

	if !strings.Contains(bucketURL, "http://") && !strings.Contains(bucketURL, "https://") {
		return nil, fmt.Errorf("%s is not a valid URL", bucketURL)
	}

	data, err := getBucket(bucketURL)
	if err != nil {
		return nil, err
	}

	var result ListBucketResult
	if err := xml.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("The URL %s does not contain a valid AWS bucket XML data", bucketURL)
	}

	if len(result.Contents) == 0 {
		return nil, fmt.Errorf("No content found in the bucket")
	}

	var keys []string
	for _, content := range result.Contents {
		keys = append(keys, content.Key)
	}

	return keys, nil
}

// getBucket retrieves the XML data from the specified bucket URL
func getBucket(bucketURL string) ([]byte, error) {
	resp, err := http.Get(bucketURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get bucket: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
