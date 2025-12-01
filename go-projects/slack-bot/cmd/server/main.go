package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"slack-bot/internal/config"
)

// TODO -

func main() {
	credentials := config.LoadConfig()
	msg := map[string]string{
		"channel": credentials.SlackChannel,
		"text":    "Hello, Slack!",
	}
	body, _ := json.Marshal(msg)
	req, err := http.NewRequest("POST", "https://slack.com/api/chat.postMessage", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+credentials.SlackToken)
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
