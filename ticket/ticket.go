package main

import "time"

type Ticket struct {
	ID           int
	User         string
	Category     string
	Description  string
	UpdateReason string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
