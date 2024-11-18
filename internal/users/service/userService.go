package service

import (
	"crypto/sha1"
	"e-commerce/internal/users/model"
	"e-commerce/internal/users/repository"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	salt       = "sdnksnk34welfw23"
	signingKey = []byte("###%4544566656")
)

type UserService struct {
	repo *repository.UserRepository
}

type tokenClaims struct {
	jwt.StandardClaims

	User_id  int    `json:"user_id"`
	Username string `json:"username"`
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// Create a user
func (s *UserService) CreateUser(user *model.User) (int, error) {
	user.Password = GeneratePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

// Get user by username and password
func (s *UserService) GetUser(username, password string) (model.User, error) {
	return s.repo.GetUser(username, password)
}

// Get all users
func (s *UserService) GetAll() ([]model.User, error) {
	return s.repo.GetAll()
}

// Sign out
func (s *UserService) SignOut(username, password string) error {
	user, err := s.repo.GetUser(username, GeneratePasswordHash(password))
	if err != nil {
		return err
	}

	return s.repo.DeleteUser(user.ID)
}

// Generate password
func GeneratePasswordHash(password string) string {
	// SHA-1 algoritmasi kullanmak icin yeni bir hash nesnesi olusturur
	hash := sha1.New()

	// Parolayi byte'lara donusturup hash nesnesine yazdirir
	hash.Write([]byte(password))

	// salt degerini ekleyip hash'in son halini alir ve stringe cevirir
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

// Generate token
func (s *UserService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, GeneratePasswordHash(password))
	if err != nil {
		return "", nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
		user.Username,
	})

	return token.SignedString(signingKey)
}

// Validate token
func (s *UserService) ValidateToken(tokenString string) (string, error) {
	claims := &tokenClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return signingKey, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	return claims.Username, nil
}

func ValidatePhoneNumber(phone string) string {
	if len(phone) != 8 {
		return "Invalid phone number: must be exactly 8 digits long."
	}

	allowedPrefixes := []string{"61", "62", "63", "64", "65", "71"}
	isValidPrefix := false
	for _, prefix := range allowedPrefixes {
		if strings.HasPrefix(phone, prefix) {
			isValidPrefix = true
			break
		}
	}

	if !isValidPrefix {
		return "Invalid phone number: must start with 61, 62, 63, 64, 65, or 71."
	}

	for _, char := range phone {
		if char < '0' || char > '9' {
			return "Invalid phone number: must contain only digits."
		}
	}

	return "Valid phone number."
}
