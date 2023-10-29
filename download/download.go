package download

import (
	"crypto/tls"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"
)

/*
File name    : download.go
Author       : miaoyc
Create Date  : 2022/12/29 11:17
Update Date  : 2023/4/4 23:47
Description  :
*/

// NormalFileDownload 普通下载文件
func NormalFileDownload(url string) {
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

// ProxyFileDownload 支持设置代理和超时时间下载文件
func ProxyFileDownload(url, path, fileName, proxy string, timeout int) error {
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}
	downFile := path + "/" + fileName
	tr := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: true,
		DialContext: (&net.Dialer{
			Timeout: time.Duration(timeout) * time.Second,
		}).DialContext,
		Proxy: getProxy(proxy),
	}

	if proxy != "" {
		tr.DialTLSContext = (&net.Dialer{
			Timeout: time.Duration(timeout) * time.Second,
		}).DialContext
	}
	var httpClient = http.Client{Transport: tr}

	resp, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("code not 200")
	}
	defer resp.Body.Close()

	out, err := os.Create(downFile)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func getProxy(address string) func(*http.Request) (*url.URL, error) {
	if len(address) == 0 {
		return nil
	}
	proxyUrl, err := url.Parse(address)
	if err != nil {
		return nil
	}
	return http.ProxyURL(proxyUrl)
}
