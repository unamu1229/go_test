package search

import (
	"io/ioutil"
	"net/http"
)

func Get() string {
	url := "http://172.54.0.54:9200/search_job/_search"
	//url := "http://127.0.0.1:9200/search_job/_search"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}