package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	goapi "github.com/grafana/grafana-openapi-client-go/client"
	"github.com/grafana/grafana-openapi-client-go/client/search"
	"github.com/grafana/grafana-openapi-client-go/models"
)

func InitGrafana() *goapi.GrafanaHTTPAPI {
	cfg := &goapi.TransportConfig{
		Host:     "volvocars.grafana.net",
		BasePath: "/api",
		Schemes:  []string{"https"},
	}
	transport := client.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	auth := client.BearerToken(os.Getenv("API_ACCESS_TOKEN"))
	transport.DefaultAuthentication = auth
	apiClient := goapi.New(transport, nil, strfmt.Default)
	return apiClient
}

func ListPublicDashboards() (string, error) {
	apiClientResp := InitGrafana()
	if apiClientResp == nil {
		return "", fmt.Errorf("error Initializing Grafana Client")
	}
	dashboardResp, err := apiClientResp.Dashboards.ListPublicDashboards(nil)
	if err != nil {
		fmt.Println("Error Getting Response", err)
		return "", err
	}
	jsonOutput, _ := json.MarshalIndent(dashboardResp.Payload, "", "  ")
	return string(jsonOutput), nil
}

func ListPrivateDashboards() ([]*models.Hit, error) {
	apiClientResp := InitGrafana()
	if apiClientResp == nil {
		return nil, fmt.Errorf("error Initializing Grafana Client")
	}
	params := search.NewSearchParams()
	dashboardType := "dash-db"
	params.Type = &dashboardType
	dashboardResp, err := apiClientResp.Search.Search(params)
	if err != nil {
		fmt.Println("Error Getting Response", err)
		return nil, err
	}
	return dashboardResp.Payload, nil
}

func main() {
	fmt.Println(ListPublicDashboards())
	dashboardResp, err := ListPrivateDashboards()
	if err != nil {
		fmt.Println("Error Getting Response", err)
		return
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', 0)
	_, _ = fmt.Fprintln(w, "TITLE\tUID\tFOLDER\t")
	for _, dashboard := range dashboardResp {
		_, _ = fmt.Fprintf(w, "%s\t%s\t%s\t\n", dashboard.Title, dashboard.UID, dashboard.FolderTitle)
	}
	_ = w.Flush()

}
