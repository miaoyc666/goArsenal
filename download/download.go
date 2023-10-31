package download

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"

	miaoycHttp "github.com/miaoyc666/goArsenal/http"
)

/*
File name    : download.go
Author       : miaoyc
Create Date  : 2022/12/29 11:17
Update Date  : 2023/10/30 5:47
Description  : */

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
	_, err = io.Copy(f, res.Body)
	if err != nil {
		return
	}
}

// loadCACert 加载CA证书
func loadCACert(caCertFile string) (*x509.CertPool, error) {
	// 读取CA证书文件
	caCert, err := ioutil.ReadFile(caCertFile)
	if err != nil {
		return nil, err
	}
	// 创建CertPool并添加CA证书
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	return caCertPool, nil
}

// ProxyFileDownload 支持设置代理和超时时间下载文件，代理类型支持
//
//	@Description:
//	@param url 下载地址
//	@param path 下载文件保持路径
//	@param fileName 下载文件名称
//	@param params 下载参数
//	@return error
func ProxyFileDownload(url, path, fileName string, params miaoycHttp.TransportParams) error {
	// 预创建目录
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}
	downFile := path + "/" + fileName

	// 自定义transport
	tr := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: params.InsecureSkipVerify},
		DisableKeepAlives: true,
		DialContext: (&net.Dialer{
			Timeout: time.Duration(params.Timeout) * time.Second,
		}).DialContext,
		Proxy: getProxy(params.Proxy),
	}

	// 存在证书时，配置TLSClientConfig
	if params.CaCertFile != "" {
		caCert, err := loadCACert(params.CaCertFile)
		if err != nil {
			return err
		}
		tr.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
			RootCAs:            caCert}
	}

	var httpClient = http.Client{Transport: tr}
	resp, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("code not 200")
	}

	out, err := os.Create(downFile)
	if err != nil {
		return err
	}

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
