package models

import (
	"time"
)

// struct holding ticket information
type Ticket struct {
	Priority    int
	ID          int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
