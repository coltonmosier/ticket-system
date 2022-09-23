/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

type tickets []Ticket

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
		var newTicket tickets
		newTicket.createTicket()
		newTicket.writeTicket()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func (t *tickets) createTicket() {
	var user string
	var cat string
	var id int
	scan := bufio.NewReader(os.Stdin)
	fmt.Print("Enter ID: ")
	fmt.Scan(&id)
	fmt.Print("Enter username: ")
	fmt.Scanln(&user)
	fmt.Print("Enter category [password_help, network_issue, other]: ")
	fmt.Scanln(&cat)
	fmt.Print("Enter description: ")
	desc, _ := scan.ReadString('\n')

	newTicket := Ticket{
		ID:          id,
		User:        user,
		Category:    cat,
		Description: desc,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	*t = append(*t, newTicket)

}

func (t *tickets) writeTicket() {
	id := t.ID
	data, err := json.Marshal(t)
	if err != nil {
		fmt.Println("error marshalling data")
	}
	file := "data/ticket" + strconv.Itoa(id) + ".json"

	err = os.WriteFile(file, data, 0644)

}
