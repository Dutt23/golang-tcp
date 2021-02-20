package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	resp, err := http.Get("http://www.google.com/robots.txt")
	defer resp.Body.Close()
	inspectResp(resp)
	resp, err = http.Head("http://www.google.com/robots.txt")
	form := url.Values{}
	form.Add("foo", "test")
	fmt.Println(form.Encode())
	resp, err = http.PostForm("https://www.google.com/robots.txt", form)
	fmt.Println(err)
	resp, err = request("PUT", "https://www.google.com/robots.txt", form)
}

func request(httpType string, url string, body url.Values) (*http.Response, error) {
	var requestBody io.Reader
	var client http.Client
	if body != nil {
		requestBody = strings.NewReader(body.Encode())
	}
	req, err := http.NewRequest(httpType, url, requestBody)
	resp, err := client.Do(req)
	return resp, err
}

func inspectResp(resp *http.Response) {
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("Reading Body")
	fmt.Println(string(body))
}
