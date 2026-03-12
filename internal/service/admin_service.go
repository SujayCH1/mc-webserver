package service

import (
	"context"
	"mc-webserver/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

var adminJWTSecret = []byte("ADMIN_SECRET_KEY")

func (s *AdminService) GenerateAdminToken(username string) (string, error) {

	claims := jwt.MapClaims{
		"username": username,
		"role":     "admin",
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(adminJWTSecret)
}
