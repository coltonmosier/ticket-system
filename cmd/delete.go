package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a ticket",
	Long: `Delete a ticket that is currently in the system requiring the following
    only if the ticket exists.
    ID: <ticket_id>
    Reason: <Description of the purpose of the deletion.`,
	Run: func(cmd *cobra.Command, args []string) {
		newTicket.readTicket("data/tickets.json")
		newTicket.delete(args)
		newTicket.writeTicket("data/tickets.json")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func (t *Tickets) delete(args []string) {
	temp := *t
	var id int
	if len(args) == 0 {
		fmt.Println("What ticket would you like to delete?")
		fmt.Print("ID: ")
		fmt.Scan(&id)
		// id, _ := strconv.Atoi(id)
		*t = append(temp[:id-1], temp[id:]...)
	} else if len(args) == 1 {
		id, _ = strconv.Atoi(args[0])
		if id < 1 || id > len(temp) {
			fmt.Printf("%s", red(fmt.Sprintf("Ticket #%v does not exist\n", args[0])))
		}
		*t = append(temp[:id-1], temp[id:]...)
	} else {
		fmt.Printf("%v\n", red("Can only delete one Ticket at a time."))
	}
}
