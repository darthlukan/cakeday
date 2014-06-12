package cakeday

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	version string = "1.0"
	baseUrl string = "http://www.reddit.com/user/"
	suffix  string = "/about.json"
	data    map[string]interface{}
)

func request(url string) (map[string]interface{}, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &data); err != nil {
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

	timestamp := time.Unix(int64(resp["data"].(map[string]interface{})["created_utc"].(float64)), 0)
	final := fmt.Sprintf("Reddit Cake Day for %v is: %v", username, timestamp.Format("1 Jan"))
	return final, err

}
