package cmd

import "fmt"

type HelpCommand struct{
	CommandName string
	CommandDescription string
}

func GetCommand() *HelpCommand{
	return &HelpCommand{
		CommandName: "help",
		CommandDescription: "Displays the help screen for the program",
	}
}

func (help *HelpCommand) CommandAction() error{
	fmt.Println("Helping")
	return nil
}