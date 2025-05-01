package client

import (
	"context"
	"net/url"

	"github.com/ready4god2513/desksdkgo/models"
)

// TicketPriorityService handles ticket-related operations
type TicketPriorityService struct {
	*Service[models.TicketPriorityResponse, models.TicketPrioritiesResponse]
}

// NewTicketPriorityService creates a new ticket service
func NewTicketPriorityService(client *Client) *TicketPriorityService {
	return &TicketPriorityService{
		Service: NewService[models.TicketPriorityResponse, models.TicketPrioritiesResponse](client, "ticketpriorities"),
	}
}

// Get retrieves a ticketpriority by ID
func (s *TicketPriorityService) Get(ctx context.Context, id int) (*models.TicketPriorityResponse, error) {
	return s.Service.Get(ctx, id)
}

// List retrieves a list of ticketpriorityes with optional filters
func (s *TicketPriorityService) List(ctx context.Context, params url.Values) (*models.TicketPrioritiesResponse, error) {
	return s.Service.List(ctx, params)
}

// Create creates a new ticketpriority
func (s *TicketPriorityService) Create(ctx context.Context, ticketpriority *models.TicketPriority) (*models.TicketPriorityResponse, error) {
	return s.Service.Create(ctx, &models.TicketPriorityResponse{TicketPriority: *ticketpriority})
}

// Update updates an existing ticketpriority
func (s *TicketPriorityService) Update(ctx context.Context, id int, ticketpriority *models.TicketPriority) (*models.TicketPriorityResponse, error) {
	return s.Service.Update(ctx, id, &models.TicketPriorityResponse{TicketPriority: *ticketpriority})
}
