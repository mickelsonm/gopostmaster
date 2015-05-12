package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	API_KEY = "tt_MTUwMDEwMDE6Wm5BeklwR3U0RkN3dEUzcFJrUno5bHJXX1RZ"
)

type RateMessage struct {
	FromZip    string  `json:"from_zip"`
	ToZip      string  `json:"to_zip"`
	Weight     float32 `json:"weight"`
	Carrier    string  `json:"carrier"`
	Packaging  string  `json:"packaging"`
	Commercial bool    `json:"commercial"`
	Service    string  `json:"service"`
}

var (
	client = new(http.Client)
)

func doRequest(method string, endpoint string, data interface{}) (result []byte, err error) {
	js, err := json.Marshal(data)
	if err != nil {
		return
	}

	method = strings.ToUpper(method)
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(js))
	if err != nil {
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	userInfo := url.UserPassword(API_KEY, "")

	pwd, _ := userInfo.Password()
	req.SetBasicAuth(userInfo.Username(), pwd)

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	result, err = ioutil.ReadAll(resp.Body)
	return
}

func main() {
	endpoint := "https://api.postmaster.io/v1/rates"
	rm := RateMessage{
		FromZip: "54701",
		ToZip:   "54729",
		Weight:  10.0,
	}

	res, err := doRequest("post", endpoint, rm)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))
}
