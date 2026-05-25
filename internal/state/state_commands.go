package state

import (
	"encoding/json"
	"log/slog"
	"os"
	"path/filepath"
	"snorp/pkg/discord"
)

func (s *SessionState) setCommands() {
	files, err := filepath.Glob(s.CommandsDir + "/*.json")
	if err != nil {
		LogAndExit("unable to read directory", err, 1)
	}

	var commands []discord.ApplicationCommand
	for _, file := range files {
		command, err := readFile(file)
		if err != nil {
			LogAndExit("unable to read command", err, 1)
		}
		slog.Info("registered command", "file", file, "command", command.Name)
		commands = append(commands, command)
	}

	currentCommands, err := s.Discord.BulkOverwriteCommands(commands)
	if err != nil {
		LogAndExit("failed to overwrite commands", err, 1)
	}
	s.Commands = currentCommands
}

func readFile(file string) (discord.ApplicationCommand, error) {
	var command discord.ApplicationCommand
	content, err := os.ReadFile(file)
	if err != nil {
		return command, err
	}

	err = json.Unmarshal(content, &command)
	return command, err
}
