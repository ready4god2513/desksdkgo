package main

import (
	"context"
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

	// Execute action based on resource and action
	switch strings.ToLower(*resource) {
	case "tickets":
		cli.Call(ctx, c.Tickets, *action, *id, func() *models.Ticket {
			return &models.Ticket{Subject: "Test Ticket"}
		})
	case "customers":
		cli.Call(ctx, c.Customers, *action, *id, func() *models.Customer {
			return &models.Customer{FirstName: "Test", LastName: "Customer"}
		})
	case "companies":
		cli.Call(ctx, c.Companies, *action, *id, func() *models.Company {
			return &models.Company{Name: "Test Company"}
		})
	case "users":
		cli.Call(ctx, c.Users, *action, *id, func() *models.User {
			return &models.User{FirstName: "Test", LastName: "User"}
		})
	case "tags":
		cli.Call(ctx, c.Tags, *action, *id, func() *models.Tag {
			return &models.Tag{Name: "Test Tag"}
		})
	case "ticketstatuses":
		cli.Call(ctx, c.TicketStatuses, *action, *id, func() *models.TicketStatus {
			return &models.TicketStatus{Name: "Test Status"}
		})
	case "tickettypes":
		cli.Call(ctx, c.TicketTypes, *action, *id, func() *models.TicketType {
			return &models.TicketType{Name: "Test Type"}
		})
	case "ticketpriorities":
		cli.Call(ctx, c.TicketPriorities, *action, *id, func() *models.TicketPriority {
			return &models.TicketPriority{Name: "Test Priority"}
		})
	case "helpdocsites":
		cli.Call(ctx, c.HelpDocSites, *action, *id, func() *models.HelpDocSite {
			return &models.HelpDocSite{Name: "Test Site"}
		})
	case "helpdocarticles":
		cli.Call(ctx, c.HelpDocArticles, *action, *id, func() *models.HelpDocArticle {
			return &models.HelpDocArticle{Title: "Test Article"}
		})
	case "businesshours":
		cli.Call(ctx, c.BusinessHours, *action, *id, func() *models.BusinessHour {
			return &models.BusinessHour{Name: "Test Business Hour"}
		})
	case "slas":
		cli.Call(ctx, c.SLAs, *action, *id, func() *models.SLA {
			return &models.SLA{Name: "Test SLA"}
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
