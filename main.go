package main

import (
	"fmt"
	"github.com/wakuworks/chest/cmd"
	"os"
)

func main() {
	length := len(os.Args)
	if length >= 2 {
		switch os.Args[1] {
		case "usage":
			cmd.Usage()

		case "put":
			cmd.Put(os.Args[2:])

		case "take":
			cmd.Take(os.Args[2:])

		case "open":
			cmd.Open()

		case "list":
			cmd.List()

		case "close":
			cmd.Close()

		default:
			message := "Unrecongnized command line argument: "
			message += os.Args[1]
			message += " ( see: 'chest usage' )"
			fmt.Println(message)
		}
	} else {
		fmt.Println("@see 'chest usage'")
	}
}
