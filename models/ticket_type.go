package models

// TicketType related types
type TicketType struct {
	BaseEntity
	Name                    string      `json:"name"`
	DisplayOrder            int         `json:"displayOrder"`
	Inboxes                 []EntityRef `json:"inboxes"`
	Default                 bool        `json:"default"`
	EnabledForFutureInboxes bool        `json:"enabledForFutureInboxes"`
}
