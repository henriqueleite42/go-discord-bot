package utils

import "monetizzer/src/types"

var SlashCommands []types.SlashCommand

func GetSlashCommands() []types.SlashCommand {
	return SlashCommands
}

func RegisterSlashCommand(command types.SlashCommand) {
	SlashCommands = append(SlashCommands, command)
}
