package main

import (
	"github.com/gogit/clone"
	"log"
	"os"
)

func main() {
	mainAction := os.Args[1]

	switch mainAction {
	case "clone":
		clone.Exec(os.Args[2:])
	case "start":
		log.Println("Not implemented yet")
	case "end":
		log.Println("Not implemented yet")
	default:
		log.Println("Error while executing an action")
	}
}
