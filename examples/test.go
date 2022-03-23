package main

import (
	bytes2 "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IssueParam struct {
	Title string
	Body  string
}

func simpleGet() {
	resp, err := http.Get("https://naver.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close() // resp.Body의 Response Struct가 io.ReadCloser이므로 io.Close() 해줘야함
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	h := resp.Header
	fmt.Printf("%s\n %s\n", string(data), h)

}

func main() {
	simpleGet()
	param := IssueParam{
		Title: "abc",
		Body:  "Def",
	}

	paramJson, _ := json.Marshal(param)
	param := bytes2.NewBuffer(paramJson)

}
