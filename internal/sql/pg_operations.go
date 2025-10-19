package sql

import (
	"context"
	"snorp/internal/api"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertGuild(ctx context.Context, pool *pgxpool.Pool, guild api.Guild) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	guildQuery := `INSERT INTO guilds (
			id, name, owner_id
		) 
		VALUES (
			@id, @name, @owner_id
		)
		ON CONFLICT (id) DO UPDATE SET
			name = @name,
			owner_id = @owner_id`

	guildArgs := pgx.NamedArgs{
		"id":       guild.ID,
		"name":     guild.Name,
		"owner_id": guild.OwnerID,
	}

	_, err = conn.Exec(ctx, guildQuery, guildArgs)
	if err != nil {
		return err
	}
	return nil
}

func InsertChannel(ctx context.Context, pool *pgxpool.Pool, channel api.GuildChannels) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	channelQuery := `INSERT INTO channels (
			id, guild_id, parent_id, name, type, topic
		) 
		VALUES (
			@id, @guild_id, @parent_id, @name, @type, @topic
		)
		ON CONFLICT (id) DO UPDATE SET
			guild_id = @guild_id,
			parent_id = @parent_id,
			name = @name,
			type = @type,
			topic = @topic`

	channelArgs := pgx.NamedArgs{
		"id":        channel.ID,
		"guild_id":  channel.GuildID,
		"parent_id": channel.ParentID,
		"name":      channel.Name,
		"type":      channel.Type,
		"topic":     channel.Topic,
	}

	_, err = conn.Exec(ctx, channelQuery, channelArgs)
	if err != nil {
		return err
	}
	return nil
}

func DeleteChannel(ctx context.Context, pool *pgxpool.Pool, channel api.GuildChannels) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	channelQuery := `DELETE FROM channels WHERE id = $1`

	_, err = conn.Exec(ctx, channelQuery, channel.ID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateChannel(ctx context.Context, pool *pgxpool.Pool, channel api.GuildChannels) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	channelQuery := `UPDATE channels SET
			guild_id = @guild_id,
			parent_id = @parent_id,
			name = @name,
			type = @type,
			topic = @topic`

	channelArgs := pgx.NamedArgs{
		"guild_id":  channel.GuildID,
		"parent_id": channel.ParentID,
		"name":      channel.Name,
		"type":      channel.Type,
		"topic":     channel.Topic,
	}

	_, err = conn.Exec(ctx, channelQuery, channelArgs)
	if err != nil {
		return err
	}
	return nil
}
