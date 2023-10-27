package main

import "fmt"

type ActionHandler interface{
	CommandAction() error
}

type CommandHandler struct{
	Args []string
	Command string
	Commands map[string]*Command
}

type Command struct{
	CommadName string
	CommandDesc string
}

func NewCommandHandler(args []string, command string, commands map[string]*Command) *CommandHandler{
	return &CommandHandler{
		Args: args,
		Command: command,
		Commands: commands,
	}
}

func (handler *CommandHandler) ValidateAndReturnCommand() (Command, error){
	found := handler.Commands[handler.Command]

	fmt.Println(found)

	return Command{}, nil
}

func (handler *CommandHandler) HandleCommand(command ActionHandler){
	command.CommandAction()
}