package service

import (
	"CodeWarriors/internal/service/dtos"
	"database/sql"
)

const (
	createUser = `INSERT INTO users (role, name, password_hash) VALUES ($1,$2,$3)`
)

type AuthorizationService struct {
	db *sql.DB
}

func NewAuthorizationService(db *sql.DB) *AuthorizationService {
	return &AuthorizationService{db: db}
}

func (s *AuthorizationService) CreateUser(form dtos.RegisterForm) error {
	return nil
}
func (s *AuthorizationService) GenerateToken(form dtos.LoginForm) (string, error) {
	return "", nil
}
func (s *AuthorizationService) ParseToken(accessToken string) (int, string, error) {
	return 0, "", nil
}
