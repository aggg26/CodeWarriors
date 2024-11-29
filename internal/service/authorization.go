package service

import (
	"CodeWarriors/internal/models"
	"CodeWarriors/internal/service/dtos"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	createUser      = `INSERT INTO users (role, name, password) VALUES ($1,$2,$3)`
	getPasswordHash = `SELECT id, name, password FROM users WHERE name=$1`
)

// Спрятать в конфигах либо в env файл
const secretKey = "jfhjdshjf32387"

type AuthorizationService struct {
	db *sql.DB
}

func NewAuthorizationService(db *sql.DB) *AuthorizationService {
	return &AuthorizationService{db: db}
}

func (s *AuthorizationService) CreateUser(form dtos.RegisterForm) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to generate passwordhash")
	}
	if _, err := s.db.Exec(createUser, form.Role, form.Username, passwordHash); err != nil {
		return errors.New("failed to create new user")
	}
	return nil
}

func (s *AuthorizationService) GenerateToken(form dtos.LoginForm) (string, error) {
	var user models.User
	if err := s.db.QueryRow(getPasswordHash, form.Username).Scan(&user.ID, &user.Name, &user.PasswordHash); err != nil {
		return "", fmt.Errorf("error: not founded user with username %s", form.Username)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(form.Password)); err != nil {
		return "", errors.New("error: wrong username or password")
	}
	claims := CustomClaims{
		user.ID,
		user.Name,
		user.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", errors.New("error: failed to generate token")
	}
	return tokenString, nil
}
func (s *AuthorizationService) ParseToken(accessToken string) (int, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, "", err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return 0, "", errors.New("token claims are not of type (*CustomClaims)")
	}
	return claims.UserID, claims.Role, nil
}

type CustomClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"user_role"`
	jwt.RegisteredClaims
}
