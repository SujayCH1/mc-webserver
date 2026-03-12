package service

import (
	"context"
	"errors"
	"mc-webserver/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	PlayerRepo *repository.PlayerRepository
}

func NewAuthService(repo *repository.PlayerRepository) *AuthService {
	return &AuthService{
		PlayerRepo: repo,
	}
}

func (s *AuthService) Register(
	ctx context.Context,
	username string,
	discordID string,
	discordUsername string,
	password string,
) error {

	// check if player exists
	_, err := s.PlayerRepo.GetByUsername(ctx, username)

	if err == nil {
		return errors.New("username already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	player := repository.Player{
		Username:        username,
		DiscordID:       discordID,
		DiscordUsername: discordUsername,
		PasswordHash:    string(hash),
	}

	return s.PlayerRepo.Create(ctx, player)
}

func (s *AuthService) VerifyPassword(
	ctx context.Context,
	username string,
	password string,
) (bool, error) {

	player, err := s.PlayerRepo.GetByUsername(ctx, username)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(player.PasswordHash),
		[]byte(password),
	)

	if err != nil {
		return false, nil
	}

	return true, nil
}
