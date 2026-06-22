package storage

import (
	"os"
	"snorp/pkg/discord"
)

type Storage interface {
	SaveGuild(discord.Guild) error
	ReadGuild(discord.Guild) error
	DeleteGuild(discord.Guild) error
}

type FileStorage struct {
	Path        string
	Permissions os.FileMode
}

func NewStorage(path string, permissions uint32) (*FileStorage, error) {
	perm := os.FileMode(permissions)
	if err := os.Mkdir(path, perm); err != nil && !os.IsExist(err) {
		return nil, err
	}
	return &FileStorage{Path: path, Permissions: perm}, nil
}

func (fs *FileStorage) SaveGuild(guild discord.Guild) error {
	filePath := fs.Path + "/guild_" + guild.ID
	return saveGob(filePath, guild)
}

func (fs *FileStorage) ReadGuild(guild discord.Guild) error {
	filePath := fs.Path + "/guild_" + guild.ID
	return readGob(filePath, guild)
}

func (fs *FileStorage) DeleteGuild(guild discord.Guild) error {
	filePath := fs.Path + "/guild_" + guild.ID
	return deleteGob(filePath)
}
