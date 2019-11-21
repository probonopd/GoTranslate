package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"time"

	"github.com/aus/proxyplease"
)
import "log"

func main() {

	toBeTranslated := os.Args[2]
	var proxyclient http.Client
	proxyclient = *HTTPClientWithProxy()
	translation, err := Translate(toBeTranslated, os.Args[1], &proxyclient)
	if err != nil {
		log.Fatal(err)
	} else {

		fmt.Println(translation)
	}
}

// We want to work this behind proxies, too
// proxyplease.NewDialContext(proxyplease.Proxy{})
func HTTPClientWithProxy() *http.Client {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			DialContext:           proxyplease.NewDialContext(proxyplease.Proxy{}),
			MaxIdleConns:          200,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   20 * time.Second,
			ExpectContinueTimeout: 20 * time.Second,
			TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		},
		Jar: jar,
	}
	return client
}
