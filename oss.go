package ossutil

import (
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// Hello ...
func Hello() string {

	endpoint := os.Getenv("ENDPOINT")
	accessID := os.Getenv("ACCESS_ID")
	accessKey := os.Getenv("ACCESS_KEY")
	client, err := oss.New(endpoint, accessID, accessKey)
	fmt.Println(client)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	bucketName := "xma-rec"
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("bucket err", err)
		return ""
	}
	fmt.Println("bucket", bucket)
	lor, err := bucket.ListObjects()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println("lor", lor)

	return "Hello World"
}
