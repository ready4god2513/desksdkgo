package client

import (
	"context"
	"net/url"

	"github.com/ready4god2513/desksdkgo/models"
)

// TicketStatusService handles ticket-related operations
type TicketStatusService struct {
	*Service[models.TicketStatusResponse, models.TicketStatusesResponse]
}

// NewTicketStatusService creates a new ticket service
func NewTicketStatusService(client *Client) *TicketStatusService {
	return &TicketStatusService{
		Service: NewService[models.TicketStatusResponse, models.TicketStatusesResponse](client, "ticketstatuses"),
	}
}

// Get retrieves a ticketstatus by ID
func (s *TicketStatusService) Get(ctx context.Context, id int) (*models.TicketStatusResponse, error) {
	return s.Service.Get(ctx, id)
}

// List retrieves a list of ticketstatuses with optional filters
func (s *TicketStatusService) List(ctx context.Context, params url.Values) (*models.TicketStatusesResponse, error) {
	return s.Service.List(ctx, params)
}

// Create creates a new ticketstatus
func (s *TicketStatusService) Create(ctx context.Context, ticketstatus *models.TicketStatus) (*models.TicketStatusResponse, error) {
	return s.Service.Create(ctx, &models.TicketStatusResponse{TicketStatus: *ticketstatus})
}

// Update updates an existing ticketstatus
func (s *TicketStatusService) Update(ctx context.Context, id int, ticketstatus *models.TicketStatus) (*models.TicketStatusResponse, error) {
	return s.Service.Update(ctx, id, &models.TicketStatusResponse{TicketStatus: *ticketstatus})
}
