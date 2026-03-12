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
	message string,
) error {

	player, err := s.PlayerRepo.GetByUsername(ctx, username)
	if err != nil {
		return err
	}

	req := repository.WhitelistRequest{
		Username:        player.Username,
		DiscordID:       player.DiscordID,
		DiscordUsername: player.DiscordUsername,
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
