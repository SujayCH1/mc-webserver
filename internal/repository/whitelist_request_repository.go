package repository

import (
	"context"
	"database/sql"
	"time"
)

type WhitelistRequest struct {
	ID              int
	Username        string
	DiscordID       string
	DiscordUsername string
	Message         string
	Status          string
	CreatedAt       time.Time
}

type WhitelistRequestRepository struct {
	DB *sql.DB
}

func NewWhitelistRequestRepository(db *sql.DB) *WhitelistRequestRepository {
	return &WhitelistRequestRepository{DB: db}
}

func (r *WhitelistRequestRepository) Create(ctx context.Context, req WhitelistRequest) error {

	query := `
	INSERT INTO whitelist_requests
	(username, discord_id, discord_username, message)
	VALUES ($1,$2,$3,$4)
	`

	_, err := r.DB.ExecContext(
		ctx,
		query,
		req.Username,
		req.DiscordID,
		req.DiscordUsername,
		req.Message,
	)

	return err
}

func (r *WhitelistRequestRepository) GetPending(ctx context.Context) ([]WhitelistRequest, error) {

	query := `
	SELECT id, username, discord_id, discord_username, message, status, created_at
	FROM whitelist_requests
	WHERE status='pending'
	`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []WhitelistRequest

	for rows.Next() {

		var rqs WhitelistRequest

		err := rows.Scan(
			&rqs.ID,
			&rqs.Username,
			&rqs.DiscordID,
			&rqs.DiscordUsername,
			&rqs.Message,
			&rqs.Status,
			&rqs.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		requests = append(requests, rqs)
	}

	return requests, nil
}
