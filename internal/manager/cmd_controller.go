package manager

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"path/filepath"
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
	ticker := time.NewTicker(cc.Interval * time.Second)

	for {
		select {
		case <-ctx.Done():
			return

		case <-ticker.C:
			files, _ := filepath.Glob(cc.Path + "/*.json")

			var commands []*models.ApplicationCommand
			for _, file := range files {
				command, err := readFile(file)
				if err != nil {
					slog.Error("unable to read command", "error", err, "file", file)
				}
				commands = append(commands, command)
			}

			for _, command := range commands {
				err := registerCommand(cc.Session, command)
				if err != nil {
					slog.Error("failed to register command", "error", err)
					continue
				}
			}
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

func DeleteCommand(session *state.SessionState, id string) {
	uri := "/applications/" + session.ReadyData.Application.ID + "/commands/" + id
	_, err := session.NewRequest("DELETE", uri, nil)
	if err != nil {
		slog.Warn("failed to delete command", "error", err, "id", id)
		return
	}
}
