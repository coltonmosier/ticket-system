package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new ticket",
	Long: `Create a new ticket with the following information:
    Category: <password_help | network_issue | other>
    Description: <description of the issue.>
    User: <username>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Create a new ticket:\n\n")
		newTicket.readTicket("data/tickets.json")
		newTicket.createTicket()
		newTicket.writeTicket("data/tickets.json")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func (t *Tickets) createTicket() {
	var user string
	var cat string
	scan := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	fmt.Scanln(&user)
	fmt.Print("Enter category [password_help, network_issue, other]: ")
	fmt.Scanln(&cat)
	fmt.Print("Enter description: ")
	desc, _ := scan.ReadString('\n')

	// remove white space before and after description string
	desc = strings.TrimSpace(desc)

	nt := Ticket{
		ID:          len(*t) + 1,
		User:        user,
		Category:    cat,
		Description: desc,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	*t = append(*t, nt)

}
