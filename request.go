package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {

	resp, err := http.Get("http://www.google.com/robots.txt")
	resp, err = http.Head("http://www.google.com/robots.txt")
	form := url.Values{}
	form.Add("foo", "test")
	resp, err = http.PostForm("https://www.google.com/robots.txt", form)
	defer resp.Body.Close()
	fmt.Println(err)
	// fmt.Println(r1.Body)

	// body, err := ioutil.ReadAll(r1.Body)
	// if err != nil {
	// 	log.Panicln(err)
	// }
	// fmt.Println("Reading Body")
	// fmt.Println(string(body))
}
