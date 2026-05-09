package manager

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"snorp/internal/models"
	"snorp/internal/state"
	"time"
)

type CommandController struct {
	Session  *state.SessionState
	Path     string
	Interval time.Duration
}

func (cc *CommandController) start(ctx context.Context) {
	ticker := time.NewTicker(cc.Interval)

	for {
		files, err := filepath.Glob(cc.Path + "/*.json")
		if err != nil {
			slog.Error("unable to read directory", "error", err, "dir", cc.Path)
			continue
		}

		var commands []*models.ApplicationCommand
		for _, file := range files {
			command, err := readFile(file)
			if err != nil {
				slog.Error("unable to read command", "error", err, "file", file)
				continue
			}
			commands = append(commands, command)
		}

		for _, oldCmd := range cc.Session.Commands {
			if slices.Contains(commands, oldCmd) {
				continue
			}
			slog.Info("deleting old command", "name", oldCmd.Name)
			deleteCommand(cc.Session, oldCmd.ID)
		}

		for _, newCmd := range commands {
			if slices.Contains(cc.Session.Commands, newCmd) {
				continue
			}
			slog.Info("creating command", "name", newCmd.Name)

			err := registerCommand(cc.Session, newCmd)
			if err != nil {
				slog.Error("failed to register command", "error", err)
				continue
			}
			cc.Session.Commands = append(cc.Session.Commands, newCmd)
		}

		select {
		case <-ctx.Done():
			return

		case <-ticker.C:
			continue
		}
	}
}

func readFile(file string) (*models.ApplicationCommand, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var command *models.ApplicationCommand

	err = json.Unmarshal(content, &command)
	return command, err
}

func registerCommand(session *state.SessionState, command *models.ApplicationCommand) error {
	uri := "/applications/" + session.ReadyData.Application.ID + "/commands"

	body, err := json.Marshal(command)
	if err != nil {
		return err
	}
	_, err = session.NewRequest("POST", uri, bytes.NewBuffer(body))
	return err
}

func deleteCommand(session *state.SessionState, id string) {
	uri := "/applications/" + session.ReadyData.Application.ID + "/commands/" + id
	_, err := session.NewRequest("DELETE", uri, nil)
	if err != nil {
		slog.Warn("failed to delete command", "error", err, "id", id)
		return
	}
}
