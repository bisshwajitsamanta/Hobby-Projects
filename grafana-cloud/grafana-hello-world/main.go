package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	goapi "github.com/grafana/grafana-openapi-client-go/client"
)

func main() {
	cfg := &goapi.TransportConfig{
		Host:     "volvocars.grafana.net",
		BasePath: "/api",
		Schemes:  []string{"https"},
	}
	transport := client.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	auth := client.BearerToken(os.Getenv("API_ACCESS_TOKEN"))
	transport.DefaultAuthentication = auth
	apiClient := goapi.New(transport, nil, strfmt.Default)
	resp, err := apiClient.SignedInUser.GetSignedInUser(nil)
	if err != nil {
		fmt.Println("Error fetching signed-in user:", err)
		return
	}
	fmt.Println(resp.Payload.Login)
	b, _ := json.MarshalIndent(resp.Payload, "", "  ")
	fmt.Println(string(b))

}
