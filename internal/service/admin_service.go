package service

import (
	"context"
	"mc-webserver/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
	AdminRepo *repository.AdminRepository
}

func NewAdminService(repo *repository.AdminRepository) *AdminService {
	return &AdminService{
		AdminRepo: repo,
	}
}

func (s *AdminService) Login(
	ctx context.Context,
	username string,
	password string,
) (bool, error) {

	admin, err := s.AdminRepo.GetByUsername(ctx, username)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(admin.PasswordHash),
		[]byte(password),
	)

	if err != nil {
		return false, nil
	}

	return true, nil
}
