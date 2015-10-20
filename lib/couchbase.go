package lib

import "github.com/couchbase/gocb"

// Upsert a document into couchbase
func Upsert(myBucket *gocb.Bucket, key string, doc string) (*gocb.Cas, error) {

	cas, err := myBucket.Upsert(key, &doc, 0)
	if err != nil {
		return nil, err
	}

	return &cas, nil
}
