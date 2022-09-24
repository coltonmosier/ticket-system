package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

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
		newTicket.readTicket("data/tickets.json")
		newTicket.update(args)
		newTicket.writeTicket("data/tickets.json")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

// update will take in the args passed in when calling to update each ticket by
// ID, only update information if input is not blank
func (t *Tickets) update(args []string) {
	scan := bufio.NewReader(os.Stdin)
	var id int
	temp := *t
	if len(args) == 0 {

		fmt.Println("What ticket would you like to delete?")
		fmt.Print("ID: ")
		fmt.Scan(&id)
	} else if len(args) > 1 {
		fmt.Printf("%s", red("Too many ID's inputted\n"))
		return
	} else {
		id, _ = strconv.Atoi(args[0])
	}

	if id < 0 || id > len(temp) {
		fmt.Printf("%s", red("Ticket does not exist\n"))
		return
	}

	id -= 1

	// Ask what to update
	running := true
	for running {
		var answer string
		fmt.Printf("What would you like to update for ticket %d?\n\n", temp[id].ID)
		fmt.Println("U: user, C: category, D: description, F: completed, Q: exit")
		fmt.Print("Choice: ")
		fmt.Scan(&answer)
		answer = strings.ToLower(answer)

		switch {
		case answer == "u":
			var user string
			fmt.Print("New user: ")
			fmt.Scanln(&user)
			if user != "" {
				temp[id].User = user
			}
		case answer == "c":
			var cat string
			fmt.Print("Enter new category [password_help, network_issue, other]: ")
			fmt.Scanln(&cat)
			if cat != "" {
				temp[id].Category = cat
			}
		case answer == "d":
			fmt.Print("Enter new description: ")
			desc, _ := scan.ReadString('\n')
			desc = strings.TrimSpace(desc)
			if desc != "" {
				temp[id].Description = desc
			}
		case answer == "f":
			var response string
			var comp bool
			fmt.Print("Completed?(Y/N): ")
			fmt.Scanln(&response)
			response = strings.ToLower(response)
			if response == "y" {
				comp = true
			} else {
				comp = false
			}

			if comp != temp[id].Completed {
				temp[id].Completed = comp
			}
		case answer == "q":
			fmt.Println("Updates complete")
			running = false
		default:
			fmt.Printf("%v", red("is an invalid option\n"))
		}
	}

	// Get user's reason for update
	fmt.Print("Update Reason: ")
	uReas, _ := scan.ReadString('\n')
	uReas = strings.TrimSpace(uReas)
	temp[id].UpdateReason = uReas

	temp[id].UpdatedAt = time.Now()

}
