package http

import (
	"crypto/tls"
	"net/http"
	"net/url"
)

/*
File name    : http.go
Author       : miaoyc
Create Date  : 2023/7/27 10:20
Update Date  : 2023/7/27 10:20
Description  : http相关方法
*/

func NewHttpClient(proxy string, insecureSkipVerify_ bool) *http.Client {
	var client *http.Client
	var tr *http.Transport
	if proxy != "" {
		proxyUrl, _ := url.Parse(proxy)
		// 创建一个http客户端并设置代理Transport
		tr = &http.Transport{
			Proxy:           http.ProxyURL(proxyUrl),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureSkipVerify_},
		}
	} else {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureSkipVerify_},
		}
	}
	client = &http.Client{
		Transport: tr,
	}
	return client
}
