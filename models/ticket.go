package models

import "time"

// Ticket related types
type Ticket struct {
	BaseEntity
	Subject               string      `json:"subject"`
	Readonly              bool        `json:"readonly"`
	MessageCount          int         `json:"messageCount"`
	PreviewText           string      `json:"previewText"`
	OriginalRecipient     string      `json:"originalRecipient"`
	ResponseTimeMins      int         `json:"responseTimeMins"`
	ResolutionTimeMins    int         `json:"resolutionTimeMins"`
	HappinessSurveySentAt time.Time   `json:"happinessSurveySentAt"`
	ImagesHidden          bool        `json:"imagesHidden"`
	IsRead                bool        `json:"isRead"`
	SpamScore             int         `json:"spam_score"`
	SpamRules             any         `json:"spam_rules"`
	Customer              EntityRef   `json:"customer"`
	Contact               EntityRef   `json:"contact"`
	Inbox                 EntityRef   `json:"inbox"`
	Agent                 EntityRef   `json:"agent"`
	Status                EntityRef   `json:"status"`
	Source                EntityRef   `json:"source"`
	Type                  EntityRef   `json:"type"`
	Messages              []EntityRef `json:"messages"`
	Timelogs              []EntityRef `json:"timelogs"`
	Activities            []EntityRef `json:"activities"`
	Suggestions           struct{}    `json:"suggestions"`
}

// Response types for tickets
type TicketsResponse struct {
	Tickets    []Ticket     `json:"tickets"`
	Included   IncludedData `json:"included"`
	Pagination Pagination   `json:"pagination"`
	Meta       Meta         `json:"meta"`
}

type TicketResponse struct {
	Ticket   Ticket       `json:"ticket"`
	Included IncludedData `json:"included"`
}
