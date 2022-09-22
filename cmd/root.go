/*
Copyright © 2022 Colton Mosier <coltonmosier44@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ticket-service",
	Short: "Check and Create Tickets",
	Long: `ticket-service is used to submit IT Tickets with different options
    based on the catagory. Once submitted, the IT team will be able to read over
    the tickets to best fix any issues that have arrived.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}