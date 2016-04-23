package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"compress/gzip"
)

var myCookieJar *cookiejar.Jar

// HTTPGet : request using GET method, will encode the get params into url
func HTTPGet(urlBase string, params map[string]string) (int, string, error) {
	url := urlBase
	if params != nil && len(params) > 0 {
		url += "?"
		for k, v := range params {
			url += (k + "=" + v + "&")
		}
		url = url[:len(url)-1]
	}
	return httpRequest("GET", url, nil)
}

// HTTPPost : request using POST method, will encode the post params into formBody
func HTTPPost(urlBase string, params map[string]string) (int, string, error) {
	val := url.Values{}
	if params != nil && len(params) > 0 {
		for k, v := range params {
			val.Set(k, v)
		}
	}
	formBody := ioutil.NopCloser(strings.NewReader(val.Encode()))
	return httpRequest("POST", urlBase, formBody)
}

// httpRequest : default charset UTF-8
func httpRequest(method string, urlStr string, postBody io.ReadCloser) (int, string, error) {
	if myCookieJar == nil {
		myCookieJar, _ = cookiejar.New(nil)
	}
	client := &http.Client{Jar: myCookieJar}
	if !((postBody == nil && method == "GET") || method == "POST") {
		panic("GET OR POST error")
	}
	req, err := http.NewRequest(method, urlStr, postBody)
	if err != nil {
		log.Println("NewRequest error")
		return -1, "", err
	}
	if debugLogging {
		log.Println("httpRequest ", method, " ", urlStr, " ")
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.112 Safari/537.36")
	req.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	req.Header.Set("Connection", "keep-alive")

	if postBody != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		if debugLogging {
			log.Println("do request error")
		}
		return -1, "", err
	}
	// ungzip or not
	var body []byte
	if res.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, _ := gzip.NewReader(res.Body)
		defer gzipReader.Close()
		body, _ = ioutil.ReadAll(gzipReader)
	} else {
		body, _ = ioutil.ReadAll(res.Body)
	}
	return res.StatusCode, (string(body)), nil
}
