package handle

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//返回上传html页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internel server error")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data,err :%s\n", err.Error())
			return
		}
		defer file.Close()
		newFile, err2 := os.Create("/tmp/" + head.Filename)
		if err2 != nil {
			fmt.Printf("Failed to get data,err :%s\n", err2.Error())
			return
		}
		defer newFile.Close()
		_, err3 := io.Copy(newFile, file)
		if err3 != nil {
			fmt.Printf("Failed to get data,err :%s\n", err3.Error())
			return
		}
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload succ")
}
