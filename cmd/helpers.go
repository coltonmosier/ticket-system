package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// consts for colorizing output
const (
	ColorDefault = "\x1b[39m"

	ColorRed    = "\x1b[91m"
	ColorGreen  = "\x1b[32m"
	ColorBlue   = "\x1b[94m"
	ColorGray   = "\x1b[90m"
	ColorYellow = "\x1b[93m"
)

// Ticket is the struct that stores the data for a ticket
type Ticket struct {
	ID           int       `json:"id"`
	User         string    `json:"user"`
	Category     string    `json:"category"`
	Description  string    `json:"description"`
	UpdateReason string    `json:"updateReason"`
	Completed    bool      `json:"completed"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// Tickets is a slice of struct Ticket
type Tickets []Ticket

// variabel of Tickets to use everywhere in program
var newTicket Tickets

// writeTicket takes the slice of ticket and marshals the data to print to a file for storage.
func (t *Tickets) writeTicket(fileName string) {
	temp := *t
	for i := range temp {
		id := i + 1
		temp[i].ID = id
	}
	res, err := json.Marshal(t)
	Check(err)
	os.WriteFile(fileName, res, 0644)
	Check(err)
}

// readTicket reads in data from a .json file that is passed as an argument,
// Unmarshals the data and stores it into the slice of ticket struct
func (t *Tickets) readTicket(fileName string) {
	bytes, err := os.ReadFile(fileName)
	Check(err)
	if len(bytes) == 0 {
		return
	}

	err = json.Unmarshal(bytes, t)
	Check(err)
}

// Check checks for error witht he passed in error argument
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// These funcs are used to colorize the text for output

func red(s string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, s, ColorDefault)
}

func yellow(s string) string {
	return fmt.Sprintf("%s%s%s", ColorYellow, s, ColorDefault)
}

func green(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)
}

func blue(s string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, s, ColorDefault)
}

func gray(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGray, s, ColorDefault)
}
