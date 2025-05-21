package models

type SLANotificationType string

const (
	SLANotificationTypeFirstResponse  SLANotificationType = "firstResponse"
	SLANotificationTypeReplyTime      SLANotificationType = "replyTime"
	SLANotificationTypeResolutionTime SLANotificationType = "resolutionTime"
)

var ValidSLANotificationTypes = []string{
	string(SLANotificationTypeFirstResponse),
	string(SLANotificationTypeReplyTime),
	string(SLANotificationTypeResolutionTime),
}

type SLAConditionOption string

const (
	SLAConditionOptionEqual    SLAConditionOption = "eq"
	SLAConditionOptionNotEqual SLAConditionOption = "ne"
)

// SLA represents a SLA in the system
type SLA struct {
	BaseEntity
	Name             string      `json:"name"`
	Description      string      `json:"description"`
	DisplayOrder     int         `json:"displayOrder"`
	Enabled          bool        `json:"enabled"`
	BusinessHour     *EntityRef  `json:"businesshours,omitempty"`
	Customers        []EntityRef `json:"slacustomers"`
	Companies        []EntityRef `json:"slacompanies"`
	Inboxes          []EntityRef `json:"slainboxes"`
	TicketTypes      []EntityRef `json:"slatickettypes"`
	TicketPriorities []EntityRef `json:"slaticketpriorities"`
	Tags             []EntityRef `json:"slatags"`
	Notifications    []EntityRef `json:"slanotifications"`
	Threads          []EntityRef `json:"threads"`
}

type SLANotification struct {
	BaseEntity
	Description        string              `json:"description"`
	Type               SLANotificationType `json:"type"`
	Duration           int                 `json:"duration"`
	User               *EntityRef          `json:"user,omitempty"`
	NotifyAssignedUser bool                `json:"notify_assigned_user"`
}

type SLATicketPriority struct {
	BaseEntity
	Hours          int        `json:"hours"`
	Minutes        int        `json:"minutes"`
	Description    string     `json:"description"`
	TicketPriority *EntityRef `json:"priority"`
}

type SLACustomer struct {
	BaseEntity
	Customer  *EntityRef         `json:"customer"`
	Condition SLAConditionOption `json:"conditionoption" db:"conditionoption"`
}

type SLACompany struct {
	BaseEntity
	Company   *EntityRef         `json:"company"`
	Condition SLAConditionOption `json:"conditionoption" db:"conditionoption"`
}

type SLAInbox struct {
	BaseEntity
	Inbox     *EntityRef         `json:"inbox"`
	Condition SLAConditionOption `json:"conditionoption" db:"conditionoption"`
}

type SLATag struct {
	BaseEntity
	Tag       *EntityRef         `json:"tag"`
	Condition SLAConditionOption `json:"conditionoption" db:"conditionoption"`
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
