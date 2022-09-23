/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
    "encoding/json"
    "io"
    "bufio"

	"github.com/alexeyco/simpletable"

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
    table := simpletable.New()

    table.Header = &simpletable.Header{
        Cells: []*simpletable.Cell{
        }
    }
    
}

func readTicket(filename string) *Ticket {

    err := json.Unmarshal()
    if err != nil {
        panic(err)
    }
}
