package search

import (
	"io/ioutil"
	"net/http"
)

func Get() string {
	resp, _ := http.Get("http://172.19.0.3:9200/search_job/_search")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}