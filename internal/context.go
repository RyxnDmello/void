package internal

import (
	"github.com/ryxndmello/void/command"
)

type Context struct {
	path      string
	command   *command.Command
	flags     map[string]string
	arguments map[string]*command.Argument
}

func NewContext(path string, command *command.Command, flags map[string]string, arguments map[string]*command.Argument) *Context {
	if command == nil {
		panic("A Command Must Be Assigned To The Void Context")
	}

	return &Context{
		path:      path,
		command:   command,
		flags:     flags,
		arguments: arguments,
	}
}

func (ctx *Context) Execute() error {
	if preExecute := ctx.command.PreExecute(); preExecute != nil {
		err := preExecute(ctx.path, ctx.command, ctx.arguments, ctx.flags)

		if err != nil {
			return err
		}
	}

	execute := ctx.command.Execute()

	if execute == nil {
		panic("Command Execute Handler Must Be Provided")
	}

	err := execute(ctx.path, ctx.command, ctx.arguments, ctx.flags)

	if err != nil {
		return err
	}

	if postExecute := ctx.command.PostExecute(); postExecute != nil {
		err := postExecute(ctx.path, ctx.command, ctx.arguments, ctx.flags)

		if err != nil {
			return err
		}
	}

	return nil
}
