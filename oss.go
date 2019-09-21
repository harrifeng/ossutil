package ossutil

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// Hello ...
func Hello() string {

	endpoint := ""
	accessID := ""
	accessKey := ""
	client, err := oss.New(endpoint, accessID, accessKey)
	fmt.Println(client)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return "Hello World"
}
