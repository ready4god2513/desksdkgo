package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/ready4god2513/desksdkgo/cli"
	"github.com/ready4god2513/desksdkgo/client"
	"github.com/ready4god2513/desksdkgo/models"
	"github.com/sirupsen/logrus"
)

func main() {
	// Define flags with environment variable fallbacks
	apiKey := flag.String("api-key", getEnv("DESK_API_KEY", ""), "Desk API key (can also be set via DESK_API_KEY env var)")
	baseURL := flag.String("base-url", getEnv("DESK_BASE_URL", "https://mycompany.teamwork.com/desk/api/v2"), "Desk API base URL (can also be set via DESK_BASE_URL env var)")
	resource := flag.String("resource", getEnv("DESK_RESOURCE", "tickets"), "Resource to interact with (tickets, customers, companies, users) (can also be set via DESK_RESOURCE env var)")
	action := flag.String("action", getEnv("DESK_ACTION", "list"), "Action to perform (get, list, create, update) (can also be set via DESK_ACTION env var)")
	id := flag.Int("id", 0, "Resource ID for get/update actions")
	debug := flag.Bool("debug", false, "Enable debug logging")
	data := flag.String("data", "", "JSON data to merge with default values for create/update actions")
	flag.Parse()

	// Validate required flags
	if *apiKey == "" {
		log.Fatal("API key is required. Set it via --api-key flag or DESK_API_KEY environment variable")
	}

	// Create client
	opts := []client.Option{}
	if *debug {
		opts = append(opts, client.WithLogLevel(logrus.DebugLevel))
	}
	opts = append(opts, client.WithAPIKey(*apiKey))

	c := client.NewClient(*baseURL, opts...)

	// Create context
	ctx := context.Background()

	// Parse JSON data if provided
	var jsonData map[string]interface{}
	if *data != "" {
		if err := json.Unmarshal([]byte(*data), &jsonData); err != nil {
			log.Fatalf("Failed to parse JSON data: %v", err)
		}
	}

	// Execute action based on resource and action
	switch strings.ToLower(*resource) {
	case "tickets":
		cli.Call(ctx, c.Tickets, *action, *id, func() *models.TicketResponse {
			resp := &models.TicketResponse{Ticket: models.Ticket{Subject: "Test Ticket"}}
			if jsonData != nil {
				mergeJSONData(&resp.Ticket, jsonData)
			}
			return resp
		})
	case "customers":
		cli.Call(ctx, c.Customers, *action, *id, func() *models.CustomerResponse {
			resp := &models.CustomerResponse{Customer: models.Customer{FirstName: "Test", LastName: "Customer"}}
			if jsonData != nil {
				mergeJSONData(&resp.Customer, jsonData)
			}
			return resp
		})
	case "companies":
		cli.Call(ctx, c.Companies, *action, *id, func() *models.CompanyResponse {
			resp := &models.CompanyResponse{Company: models.Company{Name: "Test Company"}}
			if jsonData != nil {
				mergeJSONData(&resp.Company, jsonData)
			}
			return resp
		})
	case "users":
		cli.Call(ctx, c.Users, *action, *id, func() *models.UserResponse {
			resp := &models.UserResponse{User: models.User{FirstName: "Test", LastName: "User"}}
			if jsonData != nil {
				mergeJSONData(&resp.User, jsonData)
			}
			return resp
		})
	case "tags":
		cli.Call(ctx, c.Tags, *action, *id, func() *models.TagResponse {
			resp := &models.TagResponse{Tag: models.Tag{Name: "Test Tag"}}
			if jsonData != nil {
				mergeJSONData(&resp.Tag, jsonData)
			}
			return resp
		})
	case "ticketstatuses":
		cli.Call(ctx, c.TicketStatuses, *action, *id, func() *models.TicketStatusResponse {
			resp := &models.TicketStatusResponse{TicketStatus: models.TicketStatus{Name: "Test Status"}}
			if jsonData != nil {
				mergeJSONData(&resp.TicketStatus, jsonData)
			}
			return resp
		})
	case "tickettypes":
		cli.Call(ctx, c.TicketTypes, *action, *id, func() *models.TicketTypeResponse {
			resp := &models.TicketTypeResponse{TicketType: models.TicketType{Name: "Test Type"}}
			if jsonData != nil {
				mergeJSONData(&resp.TicketType, jsonData)
			}
			return resp
		})
	case "ticketpriorities":
		cli.Call(ctx, c.TicketPriorities, *action, *id, func() *models.TicketPriorityResponse {
			resp := &models.TicketPriorityResponse{TicketPriority: models.TicketPriority{Name: "Test Priority"}}
			if jsonData != nil {
				mergeJSONData(&resp.TicketPriority, jsonData)
			}
			return resp
		})
	case "helpdocsites":
		cli.Call(ctx, c.HelpDocSites, *action, *id, func() *models.HelpDocSiteResponse {
			resp := &models.HelpDocSiteResponse{HelpDocSite: models.HelpDocSite{Name: "Test Site"}}
			if jsonData != nil {
				mergeJSONData(&resp.HelpDocSite, jsonData)
			}
			return resp
		})
	case "helpdocarticles":
		cli.Call(ctx, c.HelpDocArticles, *action, *id, func() *models.HelpDocArticleResponse {
			resp := &models.HelpDocArticleResponse{HelpDocArticle: models.HelpDocArticle{Title: "Test Article"}}
			if jsonData != nil {
				mergeJSONData(&resp.HelpDocArticle, jsonData)
			}
			return resp
		})
	case "businesshours":
		cli.Call(ctx, c.BusinessHours, *action, *id, func() *models.BusinessHourResponse {
			resp := &models.BusinessHourResponse{BusinessHour: models.BusinessHour{Name: "Test Business Hour"}}
			if jsonData != nil {
				mergeJSONData(&resp.BusinessHour, jsonData)
			}
			return resp
		})
	case "slas":
		cli.Call(ctx, c.SLAs, *action, *id, func() *models.SLAResponse {
			resp := &models.SLAResponse{SLA: models.SLA{Name: "Test SLA"}}
			if jsonData != nil {
				mergeJSONData(&resp.SLA, jsonData)
			}
			return resp
		})
	default:
		log.Fatalf("Unsupported resource: %s", *resource)
	}
}

// getEnv returns the value of the environment variable or the default value if not set
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// mergeJSONData merges JSON data into a struct
func mergeJSONData(target interface{}, data map[string]interface{}) {
	// Convert target to JSON
	jsonData, err := json.Marshal(target)
	if err != nil {
		log.Fatalf("Failed to marshal target: %v", err)
	}

	// Unmarshal into map
	var targetMap map[string]interface{}
	if err := json.Unmarshal(jsonData, &targetMap); err != nil {
		log.Fatalf("Failed to unmarshal target: %v", err)
	}

	// Merge data
	for k, v := range data {
		targetMap[k] = v
	}

	// Convert back to struct
	jsonData, err = json.Marshal(targetMap)
	if err != nil {
		log.Fatalf("Failed to marshal merged data: %v", err)
	}

	if err := json.Unmarshal(jsonData, target); err != nil {
		log.Fatalf("Failed to unmarshal merged data: %v", err)
	}
}
