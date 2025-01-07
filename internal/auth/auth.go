package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func HashPassword(password string) (string, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	return string(p), err
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err
}

func MakeJWT(userID uuid.UUID, username, tokenSecret string) (string, error) {
	claims := CustomClaims{
		username,
		jwt.RegisteredClaims{
			Issuer:    "htmxTodos",
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour)),
			Subject:   userID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(tokenSecret))
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	claims := CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	tokenUserID, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, err
	}

	userId, err := uuid.Parse(tokenUserID)
	if err != nil {
		return uuid.Nil, err
	}

	return userId, nil
}
