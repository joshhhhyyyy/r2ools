# r2ools &nbsp; &nbsp; &nbsp;[![Latest Release](https://img.shields.io/github/release/joshhhhyyyy/r2ools.svg)](https://github.com/joshhhhyyyy/r2ools/releases)      [![Go ReportCard](https://goreportcard.com/badge/joshhhhyyyy/r2ools)](https://goreportcard.com/report/joshhhhyyyy/r2ools)
![r2ools](https://socialify.git.ci/joshhhhyyyy/r2ools/image?font=Source%20Code%20Pro&language=1&name=1&owner=1&pattern=Overlapping%20Hexagons&theme=Dark)

## What is this?
**A wrapper for the official aws cli tool, with cloudflare r2 in mind**

**This tool was designed for use by myself only. Support may not be provided for any issues faced.**

Err handling by **[Sentry](sentry.io)** 

Made with 😖 , 😓 &amp; 🤮

## Installation
### Dependencies
- aws cli tool
### Via Go
```go get github.com/joshhhhyyyy/r2ools```

```go install github.com/joshhhhyyyy/r2ools```

```export PATH=$PATH:$(go env GOPATH)/bin``` (Add gopath to path)

### Via apt
```echo "deb [trusted=yes] https://apt.joseos.com/ ./" | sudo tee /etc/apt/sources.list.d/joseos.list```

```sudo apt update```

```sudo apt install r2ools```

## Usage
**An environment variable named ```endpoint``` that contains your s3 api endpoint url**

**r2** ```ls``` or ```list``` ```<bucket-name>```: lists all files in the bucket, lists all buckets if no name is supplied   

**r2** ```put <bucket> <path to file>```: uploads file

**r2** ```get <bucket> <file to get> <output path>```: downloads file

**r2** ```sign <bucket> <file> <expiresin>```: generate presigned url for file. Expiresin default: 3600s (1h)

	
