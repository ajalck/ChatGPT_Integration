package main

import (
	"log"
	"net/http"
	"os"

	chatgpt "github.com/ajalck/ChatGPT_Integration/chatGPTapi/api"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Failed to connect environment variables")
	}
}

func main() {
	apikey := os.Getenv("APIkey")
	apiurl := os.Getenv("APIurl")

	client := chatgpt.CreateClient(apikey, apiurl)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		client.TryChatGPT()
	})
	log.Println("Server started successfully")
	err := http.ListenAndServe("localhost:5000", nil)
	if err != nil {
		log.Fatalln("Failed to start server")
	}
}
