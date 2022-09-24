package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
)

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe a ticket by ID",
	Long: `describe will take in multiple ID's as an integer values and
    describe each ticket that corresponds to that ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		newTicket.readTicket("data/tickets.json")
		newTicket.describeTicket(args)
	},
}

func init() {
	rootCmd.AddCommand(describeCmd)
}

func (t *Tickets) describeTicket(args []string) {
	temp := *t
	var id int
	if len(args) == 0 {
		fmt.Println("What ticket would you like to describe?")
		fmt.Print("ID: ")
		fmt.Scan(&id)
		table := buildTable(temp[id-1])
		table.Println()
	} else if len(args) > 0 {
		// fmt.Printf("%s", red("Too many ID's inputted\n"))
		for i := range args {
			id, _ = strconv.Atoi(args[i])
			if id <= 0 || id > len(temp) {
				fmt.Printf("%s", red(fmt.Sprintf("Ticket #%v does not exist\n", args[i])))
				continue
			}
			table := buildTable(temp[id-1])
			table.Println()
		}
		return
	}
}

func buildTable(t Ticket) *simpletable.Table {

	idValue := blue("ID")
	use := blue("User")
	cat := blue("Category")
	comp := green("Completed")
	if t.Completed == false {
		comp = red("Completed")
	}
	desc := blue("Description")
	uReas := blue("Update Reason")
	cTime := yellow("Created Time")
	uTime := yellow("Updated Time")

	table := simpletable.New()
	// create table header
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Category"},
			{Align: simpletable.AlignCenter, Text: "Value"},
		},
	}

	r := []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: idValue},
		{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%d", t.ID)},
	}
	table.Body.Cells = append(table.Body.Cells, r)

	r = []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: comp},
		{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%t", t.Completed)},
	}
	table.Body.Cells = append(table.Body.Cells, r)

	r = []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: use},
		{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%v", t.User)},
	}
	table.Body.Cells = append(table.Body.Cells, r)

	r = []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: cat},
		{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%v", t.Category)},
	}
	table.Body.Cells = append(table.Body.Cells, r)

	r = []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: desc},
		{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%v", t.Description)},
	}
	table.Body.Cells = append(table.Body.Cells, r)

	r = []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: uReas},
		{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%v", t.UpdateReason)},
	}
	table.Body.Cells = append(table.Body.Cells, r)

	r = []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: cTime},
		{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%s", t.CreatedAt.Format(time.RFC1123))},
	}
	table.Body.Cells = append(table.Body.Cells, r)

	r = []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: uTime},
		{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%s", t.UpdatedAt.Format(time.RFC1123))},
	}
	table.Body.Cells = append(table.Body.Cells, r)

	table.SetStyle(simpletable.StyleUnicode)

	return table
}
