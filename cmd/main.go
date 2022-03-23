/*
1. Github Issue 만들기
2. Github Issue 바디 가져와서 파일로 저장하기
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var token string = os.Getenv("Github_Token")
var host string = "https://api.github.com"

type IssueParam struct {
	Title string `json:"title"` //오른쪽 내용을 붙여줘야 오른쪽 이름대로 json key 생성
	Body  string `json:"body"`
}

func createIssue(url string, c *http.Client, titleName string, bodyContent string, token string) {
	param := IssueParam{
		Title: titleName,
		Body:  bodyContent,
	}

	paramJson, _ := json.Marshal(param)
	fmt.Printf("%v", bytes.NewBuffer(paramJson))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(paramJson))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Accept", "application/vnd.github.v3.json")
	req.Header.Add("Authorization", "token "+token)

	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close() // resp.Body의 Response Struct가 io.ReadCloser이므로 io.Close() 해줘야함

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//h := resp.Header
	fmt.Printf("%s\n", string(data))
}

func getIssue(url string, c *http.Client, token string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	//req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Accept", "application/vnd.github.v3.html+json")
	// 오 accept 바꿔주면 html로 줄 수 있는건 html로 제공해줌
	req.Header.Add("Authorization", "token "+token)

	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close() // resp.Body의 Response Struct가 io.ReadCloser이므로 io.Close() 해줘야함

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//println(string(data))

	var stringArray []map[string]string
	json.Unmarshal(data, &stringArray)
	fmt.Println(stringArray[0])

}

func main() {
	client := &http.Client{}
	getApiPath := "/repos/Jaclyn6/Issue_Maker/issues"
	getIssueUrl := host + getApiPath
	//createApiPath := "/repos/Jaclyn6/Issue_Maker/issues"
	//createIssueUrl := host + createApiPath
	//createIssue(createIssueUrl, client, "test", "abcd\r\ncdef", token)
	getIssue(getIssueUrl, client, token)
}
