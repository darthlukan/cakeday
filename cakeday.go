package cakeday

import (
	"fmt"
	"github.com/stretchr/objx"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	version string = "1.0"
	baseUrl string = "http://www.reddit.com/user/"
	suffix  string = "/about.json"
)

func request(url string) (objx.Map, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	data, err := objx.FromJSON(string(body))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Get(username string) (string, error) {
	url := fmt.Sprintf("%v%v%v", baseUrl, username, suffix)
	resp, err := request(url)
	if err != nil {
		return "Not found!", err
	}
	timestamp := time.Unix(int64(resp.Get("data.created_utc").Float64()), 0)
	final := fmt.Sprintf("Reddit Cake Day for %v is: %v %v", username, timestamp.Day(), timestamp.Month().String())
	fmt.Println(final)
	return final, err

}
