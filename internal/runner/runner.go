package runner

import (
	"flag"
	"fmt"
	"os"
)

func ShowHelp() {
	fmt.Println(banner)
	fmt.Println("Usage:")
	fmt.Println("      awsdump <bucketURL> [options]")
	fmt.Println("\nOptions:")
	fmt.Println("      -o, --output     Specify an output directory to save files (default: awsdump)")
	fmt.Println("      -t, --threads    Max threads to use while downloading (default: 5)")
	fmt.Println("      -scrape          Enable scraping only without saving")
	fmt.Println("\nExamples:")
	fmt.Println("      awsdump https://aws.example.org")
	fmt.Println("      awsdump https://aws.example.org -t 5 -o exampleBucket")
	fmt.Println("      awsdump https://aws.example.org -scrape")
	fmt.Println("")
	os.Exit(0)
}

func ParseArgs() (string, string, bool, int) {
	if len(os.Args) < 2 {
		ShowHelp()
	}
	bucketURL := os.Args[1]

	outputDir := flag.String("o", "awsdump", "")
	outputDirAlias := flag.String("output", "awsdump", "")
	booleanScrape := flag.Bool("scrape", false, "Enable scraping only without saving")
	maxDlThreads := flag.Int("t", 5, "")
	maxDlThreadsAlias := flag.Int("threads", 5, "")

	// Parse only the arguments after the URL.
	flag.CommandLine.Parse(os.Args[2:])

	// Synchronize values for -o and --output
	if *outputDir != "awsdump" {
		*outputDirAlias = *outputDir
	} else if *outputDirAlias != "awsdump" {
		*outputDir = *outputDirAlias
	}

	// Synchronize values for -t and --threads
	if *maxDlThreads != 1 {
		*maxDlThreadsAlias = *maxDlThreads
	} else if *maxDlThreadsAlias != 1 {
		*maxDlThreads = *maxDlThreadsAlias
	}

	return bucketURL, *outputDir, *booleanScrape, *maxDlThreads
}
