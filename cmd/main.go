package main

import (
	"fmt"
)

func main() {
	running := true
	input := "a"

	for running {
		fmt.Printf("C: create a new ticket.\nR: read a ticket.\nU: update a ticket.\nD: delete a ticket.\nQ: quit\n")
		fmt.Print(">")
		fmt.Scan(&input)
		switch {
		case input == "Q", input == "q":
			// exit func
			fmt.Println("Goodbye")
			return
		case input == "C", input == "c":
			// create new ticket
			// ticket.Create()
			fmt.Println("Creating")
		case input == "R", input == "r":
			// read a ticket from db
			// ticket.Read()
			fmt.Println("Reading")
		case input == "u", input == "u":
			// update a ticket if exists in db
			// ticket.Update()
			fmt.Println("Updating")
		case input == "D", input == "d":
			// delete a ticket if exists in db
			// ticket.Delete()
			fmt.Println("Deleting")
		default:
			// handle invalid input
		}
	}
}
