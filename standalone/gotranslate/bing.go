// Translate text using Bing

// The curl was gotten by looking at the requests with Firefox and then cleaning up a bit
// https://www.bing.com/translator/?h_text=msn_ctxt&setlang=zh-cn

// Translate works seemingly without any authorization

package main

import (
	"encoding/json"

	"io/ioutil"
	"net/http"

	"net/url"

	"strings"
)

// [{"detectedLanguage":{"language":"en","score":1.0},"translations":[{"text":"Guten Abend","to":"de"}]}]
type Translation []struct {
	DetectedLanguage struct {
		Language string  `json:"language"`
		Score    float64 `json:"score"`
	} `json:"detectedLanguage"`
	Translations []struct {
		Text string `json:"text"`
		To   string `json:"to"`
	} `json:"translations"`
}

func Translate(text string, toLanguage string, proxyclient *http.Client) (string, error) {

	if text == "" {
		return "", nil
	}

	// TODO: Check if we are online and error out if we are not

	body := strings.NewReader(`&fromLang=auto-detect&text=` + url.QueryEscape(text) + `&to=` + toLanguage)
	req, err := http.NewRequest("POST", "https://www.bing.com/ttranslatev3?isVertical=1&IID=translator.5028.1", body)
	if err != nil {
		// handle err
		return "", err
	}
	// TODO: User Agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:60.0) Gecko/20100101 Firefox/60.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Referer", "https://www.bing.com/translator/?h_text=msn_ctxt&setlang=zh-cn")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Dnt", "1")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "SRCHD=AF=NOFORM; _EDGE_S=mkt=zh-cn&ui=zh-cn&F=1; MSTC=ST=1; _tarLang=default=de")
	req.Header.Set("Te", "Trailers")

	resp, err := proxyclient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// bodyString := string(bodyBytes)
	// fmt.Println(bodyString)

	var t Translation
	err = json.Unmarshal(bodyBytes, &t)
	if err != nil {
		return "", err
	}
	return t[0].Translations[0].Text, nil

}
