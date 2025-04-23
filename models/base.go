package models

import "time"

// Base types for common fields
type BaseEntity struct {
	ID        int       `json:"id"`
	Type      any       `json:"type"` // Can be string or object
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedBy UserRef   `json:"createdBy"`
	UpdatedBy UserRef   `json:"updatedBy"`
	State     string    `json:"state"`
}

type UserRef struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

type EntityRef struct {
	ID   int            `json:"id"`
	Type string         `json:"type"`
	Meta map[string]any `json:"meta"`
}
