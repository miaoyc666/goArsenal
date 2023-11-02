package http

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

/*
File name    : http.go
Author       : miaoyc
Create Date  : 2023/7/27 10:20
Update Date  : 2023/7/27 10:20
Description  : http相关方法
*/

type TransportParams struct {
	Proxy              string
	CaCertFile         string
	InsecureSkipVerify bool
	Timeout            int
}

func NewTransportParams() *TransportParams {
	return &TransportParams{InsecureSkipVerify: true, Timeout: 10}
}

func NewHttpClient(params TransportParams) *http.Client {
	var client *http.Client
	var tr *http.Transport

	// 自定义transport
	if params.Proxy != "" {
		tr = &http.Transport{
			Proxy:             GetProxy(params.Proxy),
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: params.InsecureSkipVerify},
			DisableKeepAlives: true,
			DialContext: (&net.Dialer{
				Timeout: time.Duration(params.Timeout) * time.Second,
			}).DialContext,
		}
	} else {
		tr = &http.Transport{
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: params.InsecureSkipVerify},
			DisableKeepAlives: true,
			DialContext: (&net.Dialer{
				Timeout: time.Duration(params.Timeout) * time.Second,
			}).DialContext,
		}
	}

	// 存在证书时，配置TLSClientConfig
	if params.CaCertFile != "" {
		caCert, err := LoadCACert(params.CaCertFile)
		if err != nil {
			return nil
		}
		tr.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
			RootCAs:            caCert}
	}

	client = &http.Client{
		Transport: tr,
	}
	return client
}

// LoadCACert 加载CA证书
func LoadCACert(caCertFile string) (*x509.CertPool, error) {
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

// GetProxy 获取proxy
func GetProxy(address string) func(*http.Request) (*url.URL, error) {
	if len(address) == 0 {
		return nil
	}
	proxyUrl, err := url.Parse(address)
	if err != nil {
		return nil
	}
	return http.ProxyURL(proxyUrl)
}
