package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"strings"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/ready4god2513/desksdkgo/api"
	"github.com/ready4god2513/desksdkgo/client"
	"github.com/ready4god2513/desksdkgo/models"
	"github.com/ready4god2513/desksdkgo/util"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load environment variables from .env file
	util.LoadEnv()

	// Define flags with environment variable fallbacks
	apiKey := flag.String("api-key", util.GetEnv("DESK_API_KEY", ""), "Desk API key (can also be set via DESK_API_KEY env var)")
	baseURL := flag.String("base-url", util.GetEnv("DESK_BASE_URL", "https://mycompany.teamwork.com/desk/api/v2"), "Desk API base URL (can also be set via DESK_BASE_URL env var)")
	resource := flag.String("resource", util.GetEnv("DESK_RESOURCE", "tickets"), "Resource to interact with (tickets, customers, companies, users) (can also be set via DESK_RESOURCE env var)")
	action := flag.String("action", util.GetEnv("DESK_ACTION", "list"), "Action to perform (get, list, create, update) (can also be set via DESK_ACTION env var)")
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
		api.Call(ctx, c.Tickets, *action, *id, func() *models.TicketResponse {
			inboxes, err := c.Inboxes.List(ctx, nil)
			if err != nil {
				log.Fatalf("Failed to list inboxes: %v", err)
			}

			if len(inboxes.Inboxes) == 0 {
				log.Fatal("No inboxes found. Please create an inbox first.")
			}
			resp := &models.TicketResponse{Ticket: models.Ticket{
				Subject:           gofakeit.Sentence(5),
				PreviewText:       gofakeit.Paragraph(1, 2, 3, " "),
				OriginalRecipient: gofakeit.Email(),
				Inbox: models.EntityRef{
					ID: inboxes.Inboxes[0].ID,
				},
			}}
			if jsonData != nil {
				util.MergeJSONData(&resp.Ticket, jsonData)
			}
			return resp
		})
	case "customers":
		api.Call(ctx, c.Customers, *action, *id, func() *models.CustomerResponse {
			email := gofakeit.Email()
			resp := &models.CustomerResponse{
				Customer: models.Customer{
					FirstName: gofakeit.FirstName(),
					LastName:  gofakeit.LastName(),
					Email:     email,
				},
				Included: models.IncludedData{
					Contacts: []models.Contact{
						{
							BaseEntity: models.BaseEntity{
								Type: "email",
							},
							Value:  email,
							IsMain: true,
						},
					},
				},
			}
			if jsonData != nil {
				util.MergeJSONData(&resp.Customer, jsonData)
			}
			return resp
		})
	case "companies":
		api.Call(ctx, c.Companies, *action, *id, func() *models.CompanyResponse {
			resp := &models.CompanyResponse{
				Company: models.Company{
					Name:        gofakeit.Company(),
					Description: gofakeit.Paragraph(1, 2, 3, " "),
				},
				Included: models.IncludedData{
					Domains: []models.Domain{
						{
							Name: gofakeit.DomainName(),
						},
					},
				},
			}
			if jsonData != nil {
				util.MergeJSONData(&resp.Company, jsonData)
			}
			return resp
		})
	case "users":
		api.Call(ctx, c.Users, *action, *id, func() *models.UserResponse {
			resp := &models.UserResponse{User: models.User{
				FirstName: gofakeit.FirstName(),
				LastName:  gofakeit.LastName(),
				Email:     gofakeit.Email(),
			}}
			if jsonData != nil {
				util.MergeJSONData(&resp.User, jsonData)
			}
			return resp
		})
	case "tags":
		api.Call(ctx, c.Tags, *action, *id, func() *models.TagResponse {
			resp := &models.TagResponse{Tag: models.Tag{
				Name: gofakeit.Word(),
			}}
			if jsonData != nil {
				util.MergeJSONData(&resp.Tag, jsonData)
			}
			return resp
		})
	case "ticketstatuses":
		api.Call(ctx, c.TicketStatuses, *action, *id, func() *models.TicketStatusResponse {
			resp := &models.TicketStatusResponse{TicketStatus: models.TicketStatus{
				Name: gofakeit.Word(),
			}}
			if jsonData != nil {
				util.MergeJSONData(&resp.TicketStatus, jsonData)
			}
			return resp
		})
	case "tickettypes":
		api.Call(ctx, c.TicketTypes, *action, *id, func() *models.TicketTypeResponse {
			resp := &models.TicketTypeResponse{TicketType: models.TicketType{
				Name: gofakeit.Word(),
			}}
			if jsonData != nil {
				util.MergeJSONData(&resp.TicketType, jsonData)
			}
			return resp
		})
	case "ticketpriorities":
		api.Call(ctx, c.TicketPriorities, *action, *id, func() *models.TicketPriorityResponse {
			resp := &models.TicketPriorityResponse{TicketPriority: models.TicketPriority{
				Name:  gofakeit.Word(),
				Color: gofakeit.Color(),
			}}
			if jsonData != nil {
				util.MergeJSONData(&resp.TicketPriority, jsonData)
			}
			return resp
		})
	case "helpdocsites":
		api.Call(ctx, c.HelpDocSites, *action, *id, func() *models.HelpDocSiteResponse {
			resp := &models.HelpDocSiteResponse{HelpDocSite: models.HelpDocSite{
				Name: gofakeit.Company() + " Help Center",
			}}
			if jsonData != nil {
				util.MergeJSONData(&resp.HelpDocSite, jsonData)
			}
			return resp
		})
	case "helpdocarticles":
		api.Call(ctx, c.HelpDocArticles, *action, *id, func() *models.HelpDocArticleResponse {
			resp := &models.HelpDocArticleResponse{HelpDocArticle: models.HelpDocArticle{
				Title:    gofakeit.Sentence(5),
				Contents: gofakeit.Paragraph(3, 5, 10, "\n"),
			}}
			if jsonData != nil {
				util.MergeJSONData(&resp.HelpDocArticle, jsonData)
			}
			return resp
		})
	case "businesshours":
		api.Call(ctx, c.BusinessHours, *action, *id, func() *models.BusinessHourResponse {
			resp := &models.BusinessHourResponse{BusinessHour: models.BusinessHour{
				Name: gofakeit.Company() + " Business Hours",
			}}
			if jsonData != nil {
				util.MergeJSONData(&resp.BusinessHour, jsonData)
			}
			return resp
		})
	case "slas":
		api.Call(ctx, c.SLAs, *action, *id, func() *models.SLAResponse {
			priorities, err := c.TicketPriorities.List(ctx, nil)
			if err != nil {
				log.Fatalf("Failed to list ticketpriorities: %v", err)
			}

			if len(priorities.TicketPriorities) == 0 {
				log.Fatal("No ticketpriorities found. Please create a ticketpriority first.")
			}

			inboxes, err := c.Inboxes.List(ctx, nil)
			if err != nil {
				log.Fatalf("Failed to list inboxes: %v", err)
			}

			if len(inboxes.Inboxes) == 0 {
				log.Fatal("No inboxes found. Please create an inbox first.")
			}

			businesshours, err := c.BusinessHours.List(ctx, nil)
			if err != nil {
				log.Fatalf("Failed to list businesshours: %v", err)
			}

			if len(businesshours.BusinessHours) == 0 {
				log.Fatal("No businesshours found. Please create a businesshour first.")
			}

			resp := &models.SLAResponse{
				SLA: models.SLA{
					Name: gofakeit.Company() + " SLA Policy",
					BusinessHour: &models.EntityRef{
						ID: businesshours.BusinessHours[0].ID,
					},
				},
				Included: models.IncludedData{
					SLANotifications: []models.SLANotification{
						{
							Description:        "SLA Notification",
							Type:               models.SLANotificationTypeFirstResponse,
							Duration:           gofakeit.Number(1, 10),
							NotifyAssignedUser: true,
						},
					},
				},
			}

			for _, priority := range priorities.TicketPriorities {
				resp.Included.SLAPriorities = append(resp.Included.SLAPriorities, models.SLATicketPriority{
					Hours:       gofakeit.Number(1, 10),
					Minutes:     gofakeit.Number(1, 59),
					Description: "SLA for " + priority.Name,
					TicketPriority: &models.EntityRef{
						ID: priority.ID,
					},
				})
			}

			resp.Included.SLAPriorities = append(resp.Included.SLAPriorities, models.SLATicketPriority{
				Hours:       gofakeit.Number(1, 10),
				Minutes:     gofakeit.Number(1, 59),
				Description: "SLA for None",
			})

			for _, inbox := range inboxes.Inboxes {
				resp.SLA.Inboxes = append(resp.SLA.Inboxes, models.EntityRef{
					Meta: map[string]any{
						"inbox":           inbox.ID,
						"conditionoption": "eq",
					},
				})
			}

			if jsonData != nil {
				util.MergeJSONData(&resp.SLA, jsonData)
			}
			return resp
		})
	default:
		log.Fatalf("Unsupported resource: %s", *resource)
	}
}
