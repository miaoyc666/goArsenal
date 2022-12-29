package download

import (
	"io"
	"net/http"
	"os"
	"path"
)

/*
File name    : download.go
Author       : miaoyc
Create Date  : 2022/12/29 11:17
Update Date  : 2022/12/29 11:17
Description  :
*/

func downloadFile(url string) {
	fileName := path.Base(url)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)

	}
	io.Copy(f, res.Body)
}
