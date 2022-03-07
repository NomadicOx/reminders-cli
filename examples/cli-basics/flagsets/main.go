package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You need to provide a command")
		os.Exit(2)
	}

	cmd := os.Args[1]
	switch cmd {
	case "greet":
		greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
		msgFlag := greetCmd.String("msg", "CLI BASICS - REMINDERS CLI", "the message for greet command")
		err := greetCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("hello and welcome: %s", *msgFlag)
	case "help":
		fmt.Println("some help message")
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
	}
}
