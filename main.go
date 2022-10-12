package main

import (
	"fmt"
	"os"
)

var (
	accessKey = os.Getenv("QINIU_ACCESS_KEY")
	secretKey = os.Getenv("QINIU_SECRET_KEY")
)

func main() {
	fmt.Println(accessKey,secretKey)
}
