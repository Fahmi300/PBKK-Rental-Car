package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateToken(userID int, role string) string
	ValidateToken(token string) (*jwt.Token, error)
	GetUserIDByToken(token string) (int, error)
	IsUserAdmin(token string) (bool, error)
}

type jwtCustomClaim struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "Template",
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "Template"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(userID int, role string) string {
	claims := &jwtCustomClaim{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 120)),
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		log.Println(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (any, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}

func (j *jwtService) GetUserIDByToken(token string) (int, error) {
	t_Token, err := j.ValidateToken(token)
	if err != nil {
		return 0, err
	}
	claims := t_Token.Claims.(jwt.MapClaims)
	id := int(claims["user_id"].(float64)) // Convert to int
	return id, nil
}

func (j *jwtService) IsUserAdmin(token string) (bool, error) {
	t_Token, err := j.ValidateToken(token)
	if err != nil {
		return false, err
	}

	// Ambil klaim dan periksa role
	claims := t_Token.Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	// Kembalikan true jika role adalah admin, jika tidak kembalikan false
	return role == "admin", nil
}

