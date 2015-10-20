package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

import (
	"github.com/couchbase/gocb"
	"github.com/rudijs/urban/lib"
)

// bucket reference - reuse as bucket reference in the application
var bucket *gocb.Bucket

func main() {

	// Connect to Cluster
	cluster, err := gocb.Connect("couchbase://127.0.0.1")
	if err != nil {
		fmt.Println("ERRROR CONNECTING TO CLUSTER")
		panic(err)
	}

	// Open Bucket
	bucket, err = cluster.OpenBucket("default", "")
	if err != nil {
		fmt.Println("ERRROR OPENING BUCKET")
		panic(err)
	}

	fmt.Println("==> Enter keyword(s):")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		word := strings.Replace(scanner.Text(), " ", "+", -1)

		meaning, err := lib.GetURL(lib.URL + word)

		if err != nil {
			fmt.Println("Error fetching from urbandictionary.com")
			panic(err)
		}

		cas, err2 := lib.Upsert(bucket, word, meaning)
		if err2 != nil {
			fmt.Println(err2)
		}
		fmt.Printf("Inserted document CAS is `%08x`\n", cas)

		fmt.Println("==> Enter keyword(s):")
	}
}
