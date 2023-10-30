package main

import (
	"os"

	"github.com/JGugino/server-monitor/cmd"
)

func main(){
	if len(os.Args) <= 1{
		panic("Not enough arguments")
	}

	commandHandler := cmd.NewCommandHandler(os.Args[2:], os.Args[1:2][0])

	helpCmd := &cmd.HelpCmd{}

	commandHandler.AddCommand(&cmd.Command{CommandName: "help", CommandDesc: "Shows a help menu for all available commands.", CommandAction: helpCmd})

	currentCmd, err := commandHandler.ValidateAndReturnCommand();

	if err != nil{
		panic(err)
	}

	commandHandler.HandleCommand(currentCmd)
}