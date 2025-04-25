package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/ready4god2513/desksdkgo/client"
	"github.com/ready4god2513/desksdkgo/models"
)

func main() {
	// Define flags with environment variable fallbacks
	apiKey := flag.String("api-key", getEnv("DESK_API_KEY", ""), "Desk API key (can also be set via DESK_API_KEY env var)")
	baseURL := flag.String("base-url", getEnv("DESK_BASE_URL", "https://api.desk.com/v2"), "Desk API base URL (can also be set via DESK_BASE_URL env var)")
	resource := flag.String("resource", getEnv("DESK_RESOURCE", "tickets"), "Resource to interact with (tickets, customers, companies, users) (can also be set via DESK_RESOURCE env var)")
	action := flag.String("action", getEnv("DESK_ACTION", "list"), "Action to perform (get, list, create, update) (can also be set via DESK_ACTION env var)")
	id := flag.Int("id", 0, "Resource ID for get/update actions")
	flag.Parse()

	// Validate required flags
	if *apiKey == "" {
		log.Fatal("API key is required. Set it via --api-key flag or DESK_API_KEY environment variable")
	}

	// Create client
	c := client.NewClient(*baseURL, client.WithAPIKey(*apiKey))

	// Create context
	ctx := context.Background()

	// Execute action based on resource and action
	switch strings.ToLower(*resource) {
	case "tickets":
		handleTickets(ctx, c, *action, *id)
	case "customers":
		handleCustomers(ctx, c, *action, *id)
	case "companies":
		handleCompanies(ctx, c, *action, *id)
	case "users":
		handleUsers(ctx, c, *action, *id)
	case "tags":
		handleTags(ctx, c, *action, *id)
	case "ticketstatuses":
		handleTicketStatuses(ctx, c, *action, *id)
	case "tickettypes":
		handleTicketTypes(ctx, c, *action, *id)
	case "helpdocsites":
		handleHelpDocSites(ctx, c, *action, *id)
	case "helpdocarticles":
		handleHelpDocArticles(ctx, c, *action, *id)
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

// handleTickets handles ticket-related actions
func handleTickets(ctx context.Context, c *client.Client, action string, id int) {
	service := c.Tickets

	switch strings.ToLower(action) {
	case "get":
		if id == 0 {
			log.Fatal("ID is required for get action")
		}
		ticket, err := service.Get(ctx, id)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(ticket)

	case "list":
		tickets, err := service.List(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(tickets)

	case "create":
		// Example ticket creation
		ticket := &models.Ticket{
			Subject: "Test Ticket",
		}
		created, err := service.Create(ctx, ticket)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(created)

	case "update":
		if id == 0 {
			log.Fatal("ID is required for update action")
		}
		// Example ticket update
		ticket := &models.Ticket{
			Subject: "Updated Ticket",
		}
		updated, err := service.Update(ctx, id, ticket)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(updated)

	default:
		log.Fatalf("Unsupported action: %s", action)
	}
}

// handleCustomers handles customer-related actions
func handleCustomers(ctx context.Context, c *client.Client, action string, id int) {
	service := c.Customers

	switch strings.ToLower(action) {
	case "get":
		if id == 0 {
			log.Fatal("ID is required for get action")
		}
		customer, err := service.Get(ctx, id)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(customer)

	case "list":
		customers, err := service.List(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(customers)

	case "create":
		// Example customer creation
		customer := &models.Customer{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
		}
		created, err := service.Create(ctx, customer)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(created)

	case "update":
		if id == 0 {
			log.Fatal("ID is required for update action")
		}
		// Example customer update
		customer := &models.Customer{
			FirstName: "Jane",
		}
		updated, err := service.Update(ctx, id, customer)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(updated)

	default:
		log.Fatalf("Unsupported action: %s", action)
	}
}

// handleCompanies handles company-related actions
func handleCompanies(ctx context.Context, c *client.Client, action string, id int) {
	service := c.Companies

	switch strings.ToLower(action) {
	case "get":
		if id == 0 {
			log.Fatal("ID is required for get action")
		}
		company, err := service.Get(ctx, id)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(company)

	case "list":
		companies, err := service.List(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(companies)

	case "create":
		// Example company creation
		company := &models.Company{
			Name: "Example Corp",
		}
		created, err := service.Create(ctx, company)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(created)

	case "update":
		if id == 0 {
			log.Fatal("ID is required for update action")
		}
		// Example company update
		company := &models.Company{
			Name: "Updated Corp",
		}
		updated, err := service.Update(ctx, id, company)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(updated)

	default:
		log.Fatalf("Unsupported action: %s", action)
	}
}

// handleUsers handles user-related actions
func handleUsers(ctx context.Context, c *client.Client, action string, id int) {
	service := c.Users

	switch strings.ToLower(action) {
	case "get":
		if id == 0 {
			log.Fatal("ID is required for get action")
		}
		user, err := service.Get(ctx, id)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(user)

	case "list":
		users, err := service.List(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(users)

	case "create":
		// Example user creation
		user := &models.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
		}
		created, err := service.Create(ctx, user)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(created)

	case "update":
		if id == 0 {
			log.Fatal("ID is required for update action")
		}
		// Example user update
		user := &models.User{
			FirstName: "Jane",
		}
		updated, err := service.Update(ctx, id, user)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(updated)

	default:
		log.Fatalf("Unsupported action: %s", action)
	}
}

// handleTags handles tag-related actions
func handleTags(ctx context.Context, c *client.Client, action string, id int) {
	service := c.Tags

	switch strings.ToLower(action) {
	case "get":
		if id == 0 {
			log.Fatal("ID is required for get action")
		}
		tag, err := service.Get(ctx, id)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(tag)

	case "list":
		tags, err := service.List(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(tags)

	case "create":
		// Example tag creation
		tag := &models.Tag{
			Name: "Example Tag",
		}
		created, err := service.Create(ctx, tag)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(created)

	case "update":
		if id == 0 {
			log.Fatal("ID is required for update action")
		}
		// Example tag update
		tag := &models.Tag{
			Name: "Updated Tag",
		}
		updated, err := service.Update(ctx, id, tag)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(updated)

	default:
		log.Fatalf("Unsupported action: %s", action)
	}
}

// handleTicketStatuses handles ticket status-related actions
func handleTicketStatuses(ctx context.Context, c *client.Client, action string, id int) {
	service := c.TicketStatuses

	switch strings.ToLower(action) {
	case "get":
		if id == 0 {
			log.Fatal("ID is required for get action")
		}
		status, err := service.Get(ctx, id)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(status)

	case "list":
		statuses, err := service.List(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(statuses)

	case "create":
		// Example ticket status creation
		status := &models.TicketStatus{
			Name: "Example Status",
		}
		created, err := service.Create(ctx, status)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(created)

	case "update":
		if id == 0 {
			log.Fatal("ID is required for update action")
		}
		// Example ticket status update
		status := &models.TicketStatus{
			Name: "Updated Status",
		}
		updated, err := service.Update(ctx, id, status)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(updated)

	default:
		log.Fatalf("Unsupported action: %s", action)
	}
}

// handleTicketTypes handles ticket type-related actions
func handleTicketTypes(ctx context.Context, c *client.Client, action string, id int) {
	service := c.TicketTypes

	switch strings.ToLower(action) {
	case "get":
		if id == 0 {
			log.Fatal("ID is required for get action")
		}
		ticketType, err := service.Get(ctx, id)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(ticketType)

	case "list":
		ticketTypes, err := service.List(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(ticketTypes)

	case "create":
		// Example ticket type creation
		ticketType := &models.TicketType{
			Name: "Example Type",
		}
		created, err := service.Create(ctx, ticketType)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(created)

	case "update":
		if id == 0 {
			log.Fatal("ID is required for update action")
		}
		// Example ticket type update
		ticketType := &models.TicketType{
			Name: "Updated Type",
		}
		updated, err := service.Update(ctx, id, ticketType)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(updated)

	default:
		log.Fatalf("Unsupported action: %s", action)
	}
}

// handleHelpDocSites handles help doc site-related actions
func handleHelpDocSites(ctx context.Context, c *client.Client, action string, id int) {
	service := c.HelpDocSites

	switch strings.ToLower(action) {
	case "get":
		if id == 0 {
			log.Fatal("ID is required for get action")
		}
		site, err := service.Get(ctx, id)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(site)

	case "list":
		sites, err := service.List(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(sites)

	case "create":
		// Example help doc site creation
		site := &models.HelpDocSite{
			Name: "Example Site",
		}
		created, err := service.Create(ctx, site)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(created)

	case "update":
		if id == 0 {
			log.Fatal("ID is required for update action")
		}
		// Example help doc site update
		site := &models.HelpDocSite{
			Name: "Updated Site",
		}
		updated, err := service.Update(ctx, id, site)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(updated)

	default:
		log.Fatalf("Unsupported action: %s", action)
	}
}

// handleHelpDocArticles handles help doc article-related actions
func handleHelpDocArticles(ctx context.Context, c *client.Client, action string, id int) {
	service := c.HelpDocArticles

	switch strings.ToLower(action) {
	case "get":
		if id == 0 {
			log.Fatal("ID is required for get action")
		}
		article, err := service.Get(ctx, id)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(article)

	case "list":
		articles, err := service.List(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(articles)

	case "create":
		// Example help doc article creation
		article := &models.HelpDocArticle{
			Title:    "Example Article",
			Contents: "This is an example help doc article.",
			Status:   "published",
		}
		created, err := service.Create(ctx, article)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(created)

	case "update":
		if id == 0 {
			log.Fatal("ID is required for update action")
		}
		// Example help doc article update
		article := &models.HelpDocArticle{
			Title:    "Updated Article",
			Contents: "This is an updated help doc article.",
			Status:   "published",
		}
		updated, err := service.Update(ctx, id, article)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(updated)

	default:
		log.Fatalf("Unsupported action: %s", action)
	}
}
