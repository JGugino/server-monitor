package cmd

import "errors"

type CommandAction interface{
	ExecuteAction(cmd *Command)
}

type CommandHandler struct{
	commands map[string]*Command
	args []string
	command string

}

type Command struct{
	CommandName string
	CommandDesc string
	CommandAction CommandAction
}

func NewCommandHandler(args []string, command string) *CommandHandler{
	return &CommandHandler{
		commands: make(map[string]*Command, 6),
		args: args,
		command: command,
	}
}

func (handler *CommandHandler) AddCommand(command *Command, ){
	handler.commands[command.CommandName] = command
}

func (handler *CommandHandler) ValidateAndReturnCommand() (*Command, error){
	found := handler.commands[handler.command]

	if found == nil{
		return nil, errors.New("no command found");
	}

	return found, nil
}

func (handler *CommandHandler) HandleCommand(command *Command){
	command.CommandAction.ExecuteAction(command)
}