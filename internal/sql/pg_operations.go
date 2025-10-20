package sql

import (
	"context"
	"fmt"
	"log"
	"slices"
	"snorp/internal/api"
	"snorp/internal/state"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func DeleteStaleGuilds(ctx context.Context, pool *pgxpool.Pool, guilds []state.UnavailableGuild) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Release()

	guildIDs, err := GetTableIDs(ctx, conn, "guilds")
	if err != nil {
		log.Println(err)
		return
	}

	var readyGuilds []string
	for _, guild := range guilds {
		readyGuilds = append(readyGuilds, guild.ID)
	}

	for _, dbIDs := range guildIDs {
		if !slices.Contains(readyGuilds, dbIDs) {
			_, err = conn.Exec(ctx, `DELETE FROM guilds WHERE id = $1`, dbIDs)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func InsertGuild(ctx context.Context, pool *pgxpool.Pool, guild api.Guild) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println(err)
		return
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
		log.Println(err)
		return
	}
}

func InsertChannel(ctx context.Context, pool *pgxpool.Pool, channel api.GuildChannels) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println(err)
		return
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
		log.Println(err)
		return
	}
}

func DeleteGuild(ctx context.Context, pool *pgxpool.Pool, id string) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Release()

	query := `DELETE FROM guilds WHERE id = $1`

	_, err = conn.Exec(ctx, query, id)
	if err != nil {
		log.Println(err)
		return
	}
}

func UpdateGuild(ctx context.Context, pool *pgxpool.Pool, guild api.Guild) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Release()

	guildQuery := `UPDATE guilds SET
			name = @name,
			owner_id @owner_id`

	guildArgs := pgx.NamedArgs{
		"name":     guild.Name,
		"owner_id": guild.OwnerID,
	}

	_, err = conn.Exec(ctx, guildQuery, guildArgs)
	if err != nil {
		log.Println(err)
		return
	}
}

func DeleteChannel(ctx context.Context, pool *pgxpool.Pool, channel api.GuildChannels) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Release()

	query := `DELETE FROM channels WHERE id = $1`

	_, err = conn.Exec(ctx, query, channel.ID)
	if err != nil {
		log.Println(err)
	}
}

func UpdateChannel(ctx context.Context, pool *pgxpool.Pool, channel api.GuildChannels) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println(err)
		return
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
		log.Println(err)
		return
	}
}

func GetTableIDs(ctx context.Context, conn *pgxpool.Conn, table string) ([]string, error) {
	var ids []string

	query := fmt.Sprintf(`SELECT id FROM %s`, table)

	rows, err := conn.Query(ctx, query)
	if err != nil {
		return ids, err
	}

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}

	if rows.Err() != nil {
		return ids, err
	}

	return ids, nil
}
