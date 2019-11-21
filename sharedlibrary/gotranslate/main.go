/*
https://medium.com/learning-the-go-programming-language/calling-go-functions-from-other-languages-4c7d8bcc69bf
https://github.com/vladimirvivien/go-cshared-examples

Starting with version 1.5, the Go compiler introduced support for several
build modes via the -buildmode flag. Known as the Go Execution Modes,
these build modes extended the go tool to compile Go packages into several formats
including Go archives, Go shared libraries, C archives, C shared libraries,
and (introduced in 1.8) Go dynamic plugins.

* The package must be amain package
* The source must import the pseudo-package “C”
* Use the //export comment to annotate functions you wish to make accessible to other languages
* It must returnC.CString(...) if it wants to return a string
* An empty main function must be declared
* The package is compiled using the -buildmode=c-shared build flag to create the shared object binary:
  go build -o libgotranslate.so -buildmode=c-shared *.go
* The Qt .cpp file needs: #include "gotranslate.h"
* FIXME: How can I make it to accept: #include "gotranslate/gotranslate.h"
* The Qt .pro file needs: LIBS += -L/home/me/GoTranslate/bridged/gotranslate -lgotranslate # libgotranslate.so

*/

package main

import (
	"C"
	"crypto/tls"
	"fmt"

	"net/http"
	"net/http/cookiejar"

	"time"

	"github.com/aus/proxyplease"
)
import "log"

func main() {
	// An empty main function must be declared
}

//export DoTranslate
func DoTranslate(toLang string, toBeTranslated string) *C.char {
	fmt.Println("toLang", toLang)
	fmt.Println("toBeTranslated", toBeTranslated)
	var proxyclient http.Client
	proxyclient = *HTTPClientWithProxy()
	translation, err := Translate(toBeTranslated, toLang, &proxyclient)
	if err != nil {
		log.Fatal(err)
	}
	return C.CString(translation)
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
