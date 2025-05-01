package client

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/sirupsen/logrus"
)

// Client represents the Desk API client
type Client struct {
	baseURL    string
	apiKey     string
	logLevel   logrus.Level
	httpClient *http.Client

	// Services
	Tickets          *TicketService
	Users            *UserService
	Customers        *CustomerService
	Companies        *CompanyService
	Tags             *TagService
	TicketStatuses   *TicketStatusService
	TicketTypes      *TicketTypeService
	TicketPriorities *TicketPriorityService
	HelpDocSites     *HelpDocSiteService
	HelpDocArticles  *HelpDocArticleService
	BusinessHours    *BusinessHourService
	SLAs             *SLAService
}

// Config represents the client configuration
type Config struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

// Option is a function that configures a Client
type Option func(*Client)

// WithAPIKey sets the API key for the client
func WithAPIKey(apiKey string) Option {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithLogLevel sets the log level for the client
func WithLogLevel(level logrus.Level) Option {
	return func(c *Client) {
		c.logLevel = level
	}
}

// NewClient creates a new Desk.com API client
func NewClient(baseURL string, opts ...Option) *Client {
	client := &Client{
		baseURL: baseURL,
	}

	for _, opt := range opts {
		opt(client)
	}

	if client.httpClient == nil {
		client.httpClient = NewLoggingClient(client.logLevel)
	}

	// Initialize services
	client.Tickets = NewTicketService(client)
	client.Users = NewUserService(client)
	client.Customers = NewCustomerService(client)
	client.Companies = NewCompanyService(client)
	client.Tags = NewTagService(client)
	client.TicketStatuses = NewTicketStatusService(client)
	client.TicketTypes = NewTicketTypeService(client)
	client.HelpDocSites = NewHelpDocSiteService(client)
	client.HelpDocArticles = NewHelpDocArticleService(client)
	client.TicketPriorities = NewTicketPriorityService(client)
	client.BusinessHours = NewBusinessHourService(client)
	client.SLAs = NewSLAService(client)

	return client
}

// doRequest performs an HTTP request with the client's configuration
func (c *Client) doRequest(_ context.Context, req *http.Request) (*http.Response, error) {
	// Add API key if set
	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	// Add content type
	req.Header.Set("Content-Type", "application/json")

	// Add accept header
	req.Header.Set("Accept", "application/json")

	return c.httpClient.Do(req)
}

// ListOptions represents options for list operations
type ListOptions struct {
	Page    int
	PerPage int
	SortBy  string
	SortDir string
	Embed   string
	Fields  string
	Q       string
}

// Encode encodes the options into a query string
func (o *ListOptions) Encode() string {
	if o == nil {
		return ""
	}

	v := url.Values{}
	if o.Page > 0 {
		v.Set("page", strconv.Itoa(o.Page))
	}
	if o.PerPage > 0 {
		v.Set("per_page", strconv.Itoa(o.PerPage))
	}
	if o.SortBy != "" {
		v.Set("sort_by", o.SortBy)
	}
	if o.SortDir != "" {
		v.Set("sort_dir", o.SortDir)
	}
	if o.Embed != "" {
		v.Set("embed", o.Embed)
	}
	if o.Fields != "" {
		v.Set("fields", o.Fields)
	}
	if o.Q != "" {
		v.Set("q", o.Q)
	}

	return v.Encode()
}
