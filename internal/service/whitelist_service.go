package service

import (
	"context"
	"mc-webserver/internal/repository"
)

type WhitelistService struct {
	PlayerRepo  *repository.PlayerRepository
	RequestRepo *repository.WhitelistRequestRepository
}

func NewWhitelistService(
	playerRepo *repository.PlayerRepository,
	requestRepo *repository.WhitelistRequestRepository,
) *WhitelistService {

	return &WhitelistService{
		PlayerRepo:  playerRepo,
		RequestRepo: requestRepo,
	}
}

func (s *WhitelistService) CreateRequest(
	ctx context.Context,
	username string,
	discordID string,
	discordUsername string,
	message string,
) error {

	req := repository.WhitelistRequest{
		Username:        username,
		DiscordID:       discordID,
		DiscordUsername: discordUsername,
		Message:         message,
	}

	return s.RequestRepo.Create(ctx, req)
}

func (s *WhitelistService) ApprovePlayer(
	ctx context.Context,
	username string,
) error {

	return s.PlayerRepo.SetWhitelist(ctx, username, true)
}
