/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a current ticket in the system",
	Long: `Update a current ticket with the following information:
    ID: <ticket_id> 'ID to update'
    Category: <password_help | network_issue | other>
    Description: <description of the issue.>
    User: <username>
    Update Reason: <description of why the ticket was updated>`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
