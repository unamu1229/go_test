package search

import (
	"bytes"
	"encoding/json"

	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Get() string {
	url := "http://172.54.0.54:9200/search_job/_search"
	//url := "http://127.0.0.1:9200/search_job/_search"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

type Job struct {
	Pref string `json:"pref"`
	Employment string `json:"employment"`
	Line []Line `json:"line"`
}

type Line struct {
	Name string `json:"name"`
	Station []Station `json:"station"`
}

type Station struct {
	Name string `json:"name"`
}

func Make(id int) {
	url := "http://172.54.0.54:9200/search_job/_doc/"
	//url := "http://127.0.0.1:9200/search_job/_doc/"
	url = url + strconv.Itoa(id)
	job := Job{
		"東京" + strconv.Itoa(id),
		"full-time",
		[]Line{
			{"山手線",
					[]Station {
						{"東京"},
					},
			},
		},
	}
	v, _ := json.Marshal(job)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(v))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	_, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
}