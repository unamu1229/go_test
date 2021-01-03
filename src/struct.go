package main

import (
	"encoding/json"
	"fmt"
)

type GetBody struct {
	Query struct {
		ID string
	}
}

type GetJsonBody struct {
	Query struct {
		ID string `json:"_id"`
	} `json:"query"`
}

func main()  {

	// ネストしたstructの初期化
	hoge := GetBody{
		Query: struct{ ID string }{ID: "22"},
	}
	fmt.Println(hoge.Query.ID)

	// ネストしたstructでjsonエンコードの指定がある場合の初期化は
	// jsonからUnmarshalする方法でなければ初期化できなさそう
	b := []byte(`{"query":{"_id": "77"}}`)
	var gj GetJsonBody
	err := json.Unmarshal(b, &gj)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(gj.Query.ID)
}
