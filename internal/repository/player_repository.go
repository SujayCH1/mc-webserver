package repository

import (
	"context"
	"database/sql"
	"time"
)

type Player struct {
	ID              int
	Username        string
	DiscordID       string
	DiscordUsername string
	PasswordHash    string
	Whitelisted     bool
	Banned          bool
	CreatedAt       time.Time
}

type PlayerRepository struct {
	DB *sql.DB
}

func NewPlayerRepository(db *sql.DB) *PlayerRepository {
	return &PlayerRepository{
		DB: db,
	}
}

func (r *PlayerRepository) Create(ctx context.Context, p Player) error {

	query := `
	INSERT INTO players
	(username, discord_id, discord_username, password_hash)
	VALUES ($1,$2,$3,$4)
	`

	_, err := r.DB.ExecContext(
		ctx,
		query,
		p.Username,
		p.DiscordID,
		p.DiscordUsername,
		p.PasswordHash,
	)

	return err
}

func (r *PlayerRepository) GetByUsername(ctx context.Context, username string) (*Player, error) {

	query := `
	SELECT id, username, discord_id, discord_username, password_hash,
	       whitelisted, banned, created_at
	FROM players
	WHERE username=$1
	`

	row := r.DB.QueryRowContext(ctx, query, username)

	var p Player

	err := row.Scan(
		&p.ID,
		&p.Username,
		&p.DiscordID,
		&p.DiscordUsername,
		&p.PasswordHash,
		&p.Whitelisted,
		&p.Banned,
		&p.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *PlayerRepository) SetWhitelist(ctx context.Context, username string, value bool) error {

	query := `
	UPDATE players
	SET whitelisted=$1
	WHERE username=$2
	`

	_, err := r.DB.ExecContext(ctx, query, value, username)

	return err
}
