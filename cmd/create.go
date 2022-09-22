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

type ticket struct {
	ID           int
	User         string
	Category     string
	Description  string
	UpdateReason string
	Completed    bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type tickets []ticket

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
		newTicket := ticket{}
		createTicket(&newTicket)
		writeTicket(newTicket)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func createTicket(t *ticket) {
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

	newTicket := ticket{
		ID:          id,
		User:        user,
		Category:    cat,
		Description: desc,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	*t = newTicket

}

func writeTicket(t ticket) {
	id := t.ID
	data, err := json.Marshal(t)
	if err != nil {
		fmt.Println("error marshalling data")
	}
	file := "data/ticket" + strconv.Itoa(id) + ".json"

	err = os.WriteFile(file, data, 0644)

}
