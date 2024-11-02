package service

import (
	"crypto/sha1"
	"e-commerce/internal/users/model"
	"e-commerce/internal/users/repository"
	"fmt"
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

	User_id int `json:"user_id"`
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(user *model.User) (int, error) {
	user.Password = GeneratePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *UserService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, GeneratePasswordHash(password))
	if err != nil {
		return "", nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString(signingKey)
}

func (s *UserService) GetUser(username, password string) (model.User, error) {
	return s.repo.GetUser(username, password)
}

func (s *UserService) SignOut(username, password string) error {
	user, err := s.repo.GetUser(username, GeneratePasswordHash(password))
	if err != nil {
		return err
	}

	return s.repo.DeleteUser(user.ID)
}

func GeneratePasswordHash(password string) string {
	// SHA-1 algoritmasi kullanmak icin yeni bir hash nesnesi olusturur
	hash := sha1.New()

	// Parolayi byte'lara donusturup hash nesnesine yazdirir
	hash.Write([]byte(password))

	// salt degerini ekleyip hash'in son halini alir ve stringe cevirir
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
