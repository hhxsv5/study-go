package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"time"
	"strings"
)

func get(url string) (string, error) {
	rsp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func post(url string, form map[string]string) (string, error) {
	var query string
	for key, value := range form {
		query += "&" + key + "=" + value
	}
	rsp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(query))

	if err != nil {
		return "", err
	}

	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func xget(url string) (string, error) {
	cli := new(http.Client)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Add("User-Agent", "GoAgent/1.0")

	req.Header.Add("Cookie", "a=1; b=2")

	rsp, err := cli.Do(req)

	defer rsp.Body.Close()

	//遍历cookie
	for _, c := range rsp.Cookies() {
		fmt.Println(c.Name, c.Value, c.Expires.Unix())
	}

	body, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {

	//form := map[string]string{"phone": "18866668888", "verifyCode": "abcd"}
	start := time.Now().UnixNano()
	body, err := xget("http://passport.qa.medlinker.com/titles")
	end := time.Now().UnixNano()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("cost:", (end-start)/1000/1000, "ms")
	fmt.Println(body)

}
