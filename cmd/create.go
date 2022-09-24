/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

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

type Tickets []Ticket

var newTicket Tickets

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new ticket",
	Long: `Create a new ticket with the following information:
    Category: <password_help | network_issue | other>
    Description: <description of the issue.>
    User: <username>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
		newTicket.readTicket("C:/Go Projects/ticket-system/cmd/data/tickets.json")
		fmt.Println(newTicket)
		newTicket.createTicket()
		newTicket.writeTicket("C:/Go Projects/ticket-system/cmd/data/tickets.json")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func (t *Tickets) createTicket() {
	var user string
	var cat string
	scan := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	fmt.Scanln(&user)
	fmt.Print("Enter category [password_help, network_issue, other]: ")
	fmt.Scanln(&cat)
	fmt.Print("Enter description: ")
	desc, _ := scan.ReadString('\n')

	// remove white space before and after description string
	desc = strings.TrimSpace(desc)

	nt := Ticket{
		ID:          len(*t) + 1,
		User:        user,
		Category:    cat,
		Description: desc,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	*t = append(*t, nt)

}

// writeTicket takes the slice of ticket and marshals the data to print to a file for storage.
func (t *Tickets) writeTicket(fileName string) {
	for i, v := range *t {
		i++
		v.ID = i
	}
	res, err := json.Marshal(t)
	Check(err)
	os.WriteFile(fileName, res, 0644)
	Check(err)
}
