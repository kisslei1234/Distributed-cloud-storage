package main

import (
	"fmt"
	"net/http"

	"example.com/m/v2/handle"
)

func main() {
	http.HandleFunc("/file/upload", handle.UploadHandler)
	http.HandleFunc("/file/upload/suc", handle.UploadSucHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server,err:%s", err.Error())
	}
}
