package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	channel := os.Getenv("SLACK_CHANNEL")
	msg := map[string]string{
		"channel": channel,
		"text":    "Hello, Slack!",
	}
	body, _ := json.Marshal(msg)
	req, err := http.NewRequest("POST", "https://slack.com/api/chat.postMessage", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println("response Body:", result)
	fmt.Println("response Status:", resp.Status)
}
