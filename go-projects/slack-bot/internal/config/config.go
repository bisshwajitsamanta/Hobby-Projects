// Package config - loads env vars, tokens, etc
package config

import "os"

type Config struct {
	SlackToken   string
	SlackChannel string
}

// LoadConfig loads configuration from environment variables

func LoadConfig() Config {

	token := os.Getenv("SLACK_TOKEN")
	channel := os.Getenv("SLACK_CHANNEL")

	return Config{
		SlackToken:   token,
		SlackChannel: channel,
	}
}
