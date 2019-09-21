package ossutil

import (
	"os"
	"testing"
)

func TestHello(t *testing.T) {

	endpoint := os.Getenv("ENDPOINT")
	accessID := os.Getenv("ACCESS_ID")
	accessKey := os.Getenv("ACCESS_KEY")
	bucketName := os.Getenv("BUCKET_NAME")
	objectName := os.Getenv("OBJECT_NAME")

	want := "http://xma-rec.oss"
	if got := GetObjectTimedURL(
		endpoint,
		accessID,
		accessKey,
		bucketName,
		objectName,
		30,
	); got[:len(want)] != want {
		t.Errorf("GetObjectTimedURL() = %q, want %q", got, want)
	}
}
