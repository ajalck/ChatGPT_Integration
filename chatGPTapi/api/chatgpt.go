package chatgpt

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Client struct {
	APIurl     string
	APIkey     string
	HTTPclient http.Client
}

func CreateClient(APIkey, APIurl string) *Client {
	return &Client{
		APIkey: APIkey,
		APIurl: APIurl,
	}
}
func (c *Client) TryChatGPT() {
	fmt.Println("hello everyone")
	req, err := http.NewRequest("POST", c.APIurl, strings.NewReader(`{
		"model":"gpt-3.5-turbo",
		"prompt":"How to make an HTTP request in Golang?",
		"max_tokens":40
	}`))
	if err != nil {
		log.Fatalln("create request error:", err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.APIkey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("send request error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("read and parse response error:", err)
		return
	}
	log.Println("body: ", string(body))
}
