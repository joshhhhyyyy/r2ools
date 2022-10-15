package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/getsentry/sentry-go"
)

// Set-up err handler
func check(e error) {
	if e != nil {
		sentry.CaptureException(e)
		panic(e)
	}
}

func main() {
	// Initiallise Sentry
	sentryerr := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://0aed6c09e1c94dfd90a03f10593e328e@o1153157.ingest.sentry.io/4503987011387392",
		TracesSampleRate: 1.0,
	})
	if sentryerr != nil {
		log.Fatalf("sentry.Init: %s", sentryerr)
		panic(sentryerr)
	}

	// r2 help
	if os.Args[1] == "help" {
		fmt.Println("Available Commands:")
		fmt.Println("ls or list <bucket-name>: lists all files in the bucket, lists all buckets if no name is supplied")
		fmt.Println("put <bucket> <path to file>: uploads file")
		fmt.Println("get <bucket> <file to get> <output path>: downloads file")
		fmt.Println("sign <bucket> <file> <expiresin>: generate presigned url for file. Expiresin defaults to 3600 seconds (1 hour)")
	}

	// Check if endpoint url is provided
	endpointurl := os.Getenv("endpoint")
	if endpointurl == "" {
		panic("no endpoint url provided")
	}

	// r2 list (bucket)
	if os.Args[1] == "ls" {
		if os.Args[2] != "" {
			listBucketCmd, err := exec.Command("aws", "s3api", "list-objects-v2", "--endpoint-url", endpointurl, "--bucket", os.Args[2]).Output()
			log.Println("list of files in bucket ", os.Args[2], ":")
			log.Println(listBucketCmd)
			check(err)
		} else {
			listBuckets, err := exec.Command("aws", "s3api", "list-buckets", "--endpoint-url", endpointurl).Output()
			log.Println("list of buckets: ")
			log.Println(listBuckets)
			check(err)
		}
	} else if os.Args[1] == "list" {
		if os.Args[2] != "" {
			listBucketCmd, err := exec.Command("aws", "s3api", "list-objects-v2", "--endpoint-url", endpointurl, "--bucket", os.Args[2]).Output()
			log.Println("list of files in bucket ", os.Args[2], ":")
			log.Println(listBucketCmd)
			check(err)
		} else {
			listBuckets, err := exec.Command("aws", "s3api", "list-buckets", "--endpoint-url", endpointurl).Output()
			log.Println("list of buckets: ")
			log.Println(listBuckets)
			check(err)
		}
	}

	// r2 put
	if os.Args[1] == "put" {
		if os.Args[2] != "" {
			filetoputarray := []string{"s3://", os.Args[2], "/"}
			putCmd, err := exec.Command("aws", "s3", "cp", "--endpoint-url", endpointurl, os.Args[3], strings.Join(filetoputarray, "")).Output()
			log.Println(putCmd)
			check(err)
		} else {
			panic("no file provided")
		}
	}

	// r2 get
	if os.Args[1] == "get" {
		if os.Args[2] != "" {
			filetoputarray := []string{"s3://", os.Args[2], "/", os.Args[3]}
			if os.Args[4] != "" {
				putCmd, err := exec.Command("aws", "s3", "cp", "--endpoint-url", endpointurl, strings.Join(filetoputarray, ""), os.Args[4]).Output()
				log.Println(putCmd)
				check(err)
			} else {
				putCmd, err := exec.Command("aws", "s3", "cp", "--endpoint-url", endpointurl, strings.Join(filetoputarray, ""), ".").Output()
				log.Println(putCmd)
				check(err)
			}
		} else {
			panic("no file provided")
		}
	}

	// r2 sign (bucket) (file) (expiresin)
	if os.Args[1] == "sign" {
		if os.Args[2] != "" && os.Args[3] != "" {
			filetosignarray := []string{"s3://", os.Args[2], "/", os.Args[3]}
			var expiresin string
			rawexpiresin := os.Args[4]
			parsedexpiresin := rawexpiresin[len(rawexpiresin)-1:]
			if rawexpiresin != "" {
				if parsedexpiresin == "s" {
					expiresin = rawexpiresin[:len(rawexpiresin)-1]
				} else if parsedexpiresin == "m" {
					timeexpiresin := rawexpiresin[:len(rawexpiresin)-1]
					intVar, err := strconv.Atoi(timeexpiresin)
					check(err)
					expiresin = strconv.Itoa(intVar * 60)
				} else if parsedexpiresin == "h" {
					timeexpiresin := rawexpiresin[:len(rawexpiresin)-1]
					intVar, err := strconv.Atoi(timeexpiresin)
					check(err)
					expiresin = strconv.Itoa(intVar * 3600)
				} else if parsedexpiresin == "d" {
					timeexpiresin := rawexpiresin[:len(rawexpiresin)-1]
					intVar, err := strconv.Atoi(timeexpiresin)
					check(err)
					expiresin = strconv.Itoa(intVar * 86400)
				} else {
					panic("invalid time format. Use 's' for seconds, 'm' for minutes and 'd' for days")
				}
			} else {
				expiresin = "3600"
			}
			presignCmd, err := exec.Command("aws", "s3", "presign", "--endpoint-url", endpointurl, strings.Join(filetosignarray, ""), "--expires-in", expiresin).Output()
			log.Println(presignCmd)
			check(err)
		} else {
			panic("no file provided")
		}
	}
}
