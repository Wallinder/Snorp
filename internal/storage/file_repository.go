package storage

import (
	"os"
	"snorp/pkg/discord"
)

type Storage interface {
	SaveGuild(discord.Guild) error
	DeleteGuild(string) error
	SaveChannel(*discord.Channel, string) error
	DeleteChannel(string, string) error
	SaveMember(*discord.Member, string) error
	DeleteMember(string, string) error
	SaveRole(*discord.Role, string) error
	DeleteRole(string, string) error
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
	dirPath := fs.Path + "/guild_" + guild.ID
	if err := os.Mkdir(dirPath, fs.Permissions); err != nil && !os.IsExist(err) {
		return err
	}
	for _, channel := range guild.Channels {
		if err := fs.SaveChannel(channel, guild.ID); err != nil {
			return err
		}
	}
	for _, member := range guild.Members {
		if err := fs.SaveMember(member, guild.ID); err != nil {
			return err
		}
	}
	for _, role := range guild.Roles {
		if err := fs.SaveRole(role, guild.ID); err != nil {
			return err
		}
	}
	return nil
}

func (fs *FileStorage) DeleteGuild(guildID string) error {
	dirPath := fs.Path + "/guild_" + guildID
	return os.RemoveAll(dirPath)
}

func (fs *FileStorage) SaveChannel(channel *discord.Channel, guildID string) error {
	filePath := fs.Path + "/guild_" + guildID + "/channel_" + channel.ID
	return saveGob(filePath, channel)
}

func (fs *FileStorage) DeleteChannel(channelID string, guildID string) error {
	filePath := fs.Path + "/guild_" + guildID + "/channel_" + channelID
	return deleteGob(filePath)
}

func (fs *FileStorage) SaveMember(member *discord.Member, guildID string) error {
	filePath := fs.Path + "/guild_" + guildID + "/member_" + member.User.ID
	return saveGob(filePath, member)
}

func (fs *FileStorage) DeleteMember(userID string, guildID string) error {
	filePath := fs.Path + "/guild_" + guildID + "/channel_" + userID
	return deleteGob(filePath)
}

func (fs *FileStorage) SaveRole(role *discord.Role, guildID string) error {
	filePath := fs.Path + "/guild_" + guildID + "/role_" + role.ID
	return saveGob(filePath, role)
}

func (fs *FileStorage) DeleteRole(roleID string, guildID string) error {
	filePath := fs.Path + "/guild_" + guildID + "/role_" + roleID
	return deleteGob(filePath)
}
