package stdcommands

import (
	"github.com/nicolito128/waffer/plugins/commands"
	holdingbooks "github.com/nicolito128/waffer/stdcommands/anime/holding-books"
	"github.com/nicolito128/waffer/stdcommands/development/dev"
	"github.com/nicolito128/waffer/stdcommands/development/github"
	"github.com/nicolito128/waffer/stdcommands/development/ping"
	"github.com/nicolito128/waffer/stdcommands/fun/dog"
	"github.com/nicolito128/waffer/stdcommands/fun/rps"
	"github.com/nicolito128/waffer/stdcommands/info/avatar"
	"github.com/nicolito128/waffer/stdcommands/info/commandlist"
	"github.com/nicolito128/waffer/stdcommands/info/help"
	"github.com/nicolito128/waffer/stdcommands/info/invite"
	"github.com/nicolito128/waffer/stdcommands/util/calc"
	"github.com/nicolito128/waffer/stdcommands/util/say"
)

type WafferCommand *commands.WafferCommand

// AddCommands load all the commands.
func AddCommands() {
	commands.AddRootCommands(
		ping.Command,
		dev.Command,
		avatar.Command,
		calc.Command,
		holdingbooks.Command,
		github.Command,
		help.Command,
		say.Command,
		dog.Command,
		invite.Command,
		rps.Command,
		commandlist.Command,
	)
}

// HasCommand return true if command list has the command name passed.
func HasCommand(commandName string) bool {
	if commands.CommandList[commandName] != nil {
		return true
	}

	return false
}

// GetCommandList return a map[string]:*WafferCommand
func GetCommandList() commands.CList {
	return commands.CommandList
}
