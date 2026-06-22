package storage

import (
	"os"
	"snorp/pkg/discord"
)

type Storage interface {
	SaveGuild(discord.Guild)
	SaveUser(discord.Guild)
}

type FileStorage struct {
	Path        string
	Permissions os.FileMode
}

func NewStorage(path string, permissions uint32) (*FileStorage, error) {
	perm := os.FileMode(permissions)

	if err := os.Mkdir(path, perm); err != nil || os.IsExist(err) {
		return nil, err
	}
	return &FileStorage{Path: path, Permissions: perm}, nil
}

func (fs *FileStorage) SaveGuild(guild discord.Guild) {
}

func (fs *FileStorage) SaveUser(guild discord.Guild) {
}
