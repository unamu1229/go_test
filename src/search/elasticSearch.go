package search

import (
	"io/ioutil"
	"net/http"
)

func Get() string {
	resp, _ := http.Get("http://localhost:9200/search_job/_search")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}