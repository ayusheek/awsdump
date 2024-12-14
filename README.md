# awsdump

automates pulic aws bucket enumeration and efficient downloading to uncover exposed data in public S3 buckets.

## Installation

Using go:

```
go install -v github.com/ayushkr12/awsdump/cmd/awsdump@latest
```

or [download a release for your platform](https://github.com/ayusheek/awsdump/releases/).

## Usage

```console

                           __                    
  ____ __      ___________/ /_  ______ ___  ____ 
 / __  / | /| / / ___/ __  / / / / __  __ \/ __ \
/ /_/ /| |/ |/ (__  ) /_/ / /_/ / / / / / / /_/ /
\__,_/ |__/|__/____/\__,_/\__,_/_/ /_/ /_/ .___/ 
                                        /_/     

                                       v0.0.1

Usage:
      awsdump <bucketURL> [options]

Options:
      -o, --output     Specify an output directory to save files (default: awsdump)
      -t, --threads    Max threads to use while downloading (default: 5)
      -scrape          Enable scraping only without saving

Examples:
      awsdump https://aws.example.org
      awsdump https://aws.example.org -t 5 -o exampleBucket
      awsdump https://aws.example.org -scrape
```

## TODO

- [ ] Scrape & Download Truncated Buckets recursively
- [ ] Read URLs from stdin
- [ ] Custom & Random ua support
