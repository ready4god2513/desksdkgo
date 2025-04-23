package models

// Common pagination structure used across responses
type Pagination struct {
	Records      int  `json:"records"`
	PageSize     int  `json:"pageSize"`
	Pages        int  `json:"pages"`
	Page         int  `json:"page"`
	HasMorePages bool `json:"hasMorePages"`
}

// Common page metadata structure
type PageMeta struct {
	Count      int  `json:"count"`
	PageSize   int  `json:"pageSize"`
	PageOffset int  `json:"pageOffset"`
	Pages      int  `json:"pages"`
	HasMore    bool `json:"hasMore"`
}

// Common included data structure
type IncludedData struct {
	Companies        []Company        `json:"companies"`
	Contacts         []Contact        `json:"contacts"`
	Customers        []Customer       `json:"customers"`
	Inboxes          []Inbox          `json:"inboxes"`
	Messages         []Message        `json:"messages"`
	Tags             []Tag            `json:"tags"`
	Ticketactivities []TicketActivity `json:"ticketactivities"`
	Ticketsources    []TicketSource   `json:"ticketsources"`
	Ticketstatuses   []TicketStatus   `json:"ticketstatuses"`
	Tickettypes      []TicketType     `json:"tickettypes"`
	Timelogs         []TimeLog        `json:"timelogs"`
	Users            []User           `json:"users"`
}

// Common meta structure
type Meta struct {
	Page PageMeta `json:"page"`
}
