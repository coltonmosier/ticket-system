/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	// "github.com/alexeyco/simpletable"

	"github.com/spf13/cobra"
)

var ticket Ticket

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display the current tickets in the system.",
	Long:  `Displays a list of all current tickets in the system color coded.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

// list reads tickets from a file and prints them in a nice table
func list() {

}

func (t *Tickets) readTicket(fileName string) {
	bytes, err := os.ReadFile(fileName)
	Check(err)

	err = json.Unmarshal(bytes, t)
	fmt.Println(err)
	//Check(err)
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
