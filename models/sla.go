package models

// SLA represents a SLA in the system
type SLA struct {
	BaseEntity
	Name             string      `json:"name"`
	Description      string      `json:"description"`
	DisplayOrder     int         `json:"displayOrder"`
	BusinessHour     *EntityRef  `json:"businesshours,omitempty"`
	Customers        []EntityRef `json:"customers"`
	Companies        []EntityRef `json:"companies"`
	Inboxes          []EntityRef `json:"inboxes"`
	TicketTypes      []EntityRef `json:"tickettypes"`
	TicketPriorities []EntityRef `json:"ticketpriorities"`
	Tags             []EntityRef `json:"tags"`
	Notifications    []EntityRef `json:"notifications"`
	Threads          []EntityRef `json:"threads"`
}

// SLAsResponse represents the response for a list of slas
type SLAsResponse struct {
	SLAs       []SLA        `json:"slas"`
	Included   IncludedData `json:"included"`
	Pagination Pagination   `json:"pagination"`
	Meta       Meta         `json:"meta"`
}

type SLAResponse struct {
	SLA      SLA          `json:"sla"`
	Included IncludedData `json:"included"`
}
