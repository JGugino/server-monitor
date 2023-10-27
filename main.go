package main

import (
	"os"

	"github.com/JGugino/server-monitor/cmd"
)

func main(){
	commands := make(map[string]*Command, 6)

	helpCmd := cmd.HelpCommand{CommandName: "help", CommandDescription: "Displays the help screen for the program"}

	commandHandler := NewCommandHandler(os.Args[1:], os.Args[1:2][0], commands)

	commandHandler.HandleCommand(&helpCmd)
}