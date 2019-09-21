package ossutil

import (
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func getObjectsFormResponse(lor oss.ListObjectsResult) string {
	var output string
	for _, object := range lor.Objects {
		output += object.Key + " --- "
	}
	return output
}

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
	fmt.Println("lor", getObjectsFormResponse(lor))

	url, err := bucket.SignURL("everything.png", oss.HTTPGet, 30)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println("url:", url)

	return "Hello World"
}
