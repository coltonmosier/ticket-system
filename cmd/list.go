package cmd

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	"time"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display the current tickets in the system.",
	Long:  `Displays a list of all current tickets in the system color coded.`,
	Run: func(cmd *cobra.Command, args []string) {
		// must read ticket before we can list it
		newTicket.readTicket("data/tickets.json")
		newTicket.list()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

// list loops through the slice of ticke struct to print it legibly in a table
func (t *Tickets) list() {
	table := simpletable.New()
	// create table header
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "Completed"},
			{Align: simpletable.AlignCenter, Text: "User"},
			{Align: simpletable.AlignCenter, Text: "Category"},
			{Align: simpletable.AlignCenter, Text: "Created At"},
			{Align: simpletable.AlignCenter, Text: "Updated At"},
		},
	}

	// fill cells for table
	for _, v := range *t {
		complete := red(fmt.Sprintf("%t", v.Completed))
		if v.Completed {
			complete = blue(fmt.Sprintf("%t", v.Completed))
		}
		r := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", v.ID)},
			{Align: simpletable.AlignCenter, Text: complete},
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%s", v.User)},
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%s", v.Category)},
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%s", v.CreatedAt.Format(time.RFC1123))},
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%s", v.UpdatedAt.Format(time.RFC1123))},
		}

		table.Body.Cells = append(table.Body.Cells, r)

	}

	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}
