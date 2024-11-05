package bucket

// ListBucketResult represents the structure of the S3 bucket XML response
type ListBucketResult struct {
	Contents []struct {
		Key string `xml:"Key"`
	} `xml:"Contents"`
}
