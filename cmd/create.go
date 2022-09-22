/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/coltonmosier/ticket-service/ticket"
	"github.com/spf13/cobra"
)

var newTicket ticket.Ticket

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
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func (t *ticket.Ticket) createTicket() {

}
