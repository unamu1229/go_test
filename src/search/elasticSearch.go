package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)



type GetResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			ID     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				Pref       string `json:"pref"`
				Employment string `json:"employment"`
				Line       struct {
					Name    string   `json:"name"`
					Station []string `json:"station"`
				} `json:"line"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type GetBody struct {
	Query struct {
		Term struct {
			ID string `json:"_id"`
		} `json:"term"`
	} `json:"query"`
}

func Get(id int) string {
	url := "http://elasticsearch:9200/search_job/_search"
	//url := "http://127.0.0.1:9200/search_job/_search"
	var getBody = GetBody{}
	getBody.Query.Term.ID = strconv.Itoa(id)
	v, _ := json.Marshal(getBody)
	req, _ := http.NewRequest("GET", url, bytes.NewBuffer(v))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	fmt.Println(string(body))
	var getResponse GetResponse
	unmarshalErr := json.Unmarshal(body, &getResponse)
	if unmarshalErr != nil {
		fmt.Println(unmarshalErr)
	}

	for _, hit := range getResponse.Hits.Hits {
		fmt.Println("id", hit.ID)
		fmt.Println("pref", hit.Source.Pref)
		fmt.Println("employment", hit.Source.Employment)
		fmt.Println("employment", hit.Source.Line.Name)
		for _, station := range hit.Source.Line.Station {
			fmt.Println("station", station)
		}

	}

	fmt.Println(getResponse.Hits.Hits)

	return "ok"
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