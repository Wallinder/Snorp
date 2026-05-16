package state

import (
	"bytes"
	"encoding/json"
	"io"
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

	err = s.bulkOverwriteCommands(commands)
	if err != nil {
		LogAndExit("failed to overwrite commands", err, 1)
	}
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

func (s *SessionState) bulkOverwriteCommands(commands []discord.ApplicationCommand) error {
	uri := "/applications/" + s.ReadyData.Application.ID + "/commands"

	body, err := json.Marshal(commands)
	if err != nil {
		return err
	}

	response, err := s.NewDiscordRequest("PUT", uri, bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &s.Commands)
}
