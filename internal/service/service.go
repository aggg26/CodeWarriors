package service

import (
	"CodeWarriors/internal/service/dtos"
	"database/sql"
)

type Authorization interface {
	CreateUser(form dtos.RegisterForm) error
	GenerateToken(form dtos.LoginForm) (string, error)
	ParseToken(accessToken string) (int, string, error)
}

type Service struct {
	Authorization
}

func NewService(db *sql.DB) *Service {
	return &Service{}
}
