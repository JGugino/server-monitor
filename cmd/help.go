package cmd

import "fmt"

type HelpCmd struct{}

func (help *HelpCmd) ExecuteAction(cmd *Command){
	fmt.Println(cmd.CommandName)
	fmt.Println(cmd.CommandDesc)
}