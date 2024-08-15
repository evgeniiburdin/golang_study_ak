package jwt

import (
	"auth-service/internal/entity"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type JWTServicer interface {
	CreateJWTs(user entity.User) (accessToken string, refreshToken string, err error)
	ReadJWT(accessToken string) (claims jwt.MapClaims, err error)
}

type JWTService struct{}

func (s *JWTService) CreateJWTs(user entity.User) (accessToken string, refreshToken string, err error) {

	log.Printf("creating new jwts for user: %+v", user)

	accessPayload := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  (time.Now().Local().Add(time.Minute * 15)).Unix(),
		"sub":  user.Email,
		"typ":  "access",
		"role": user.Role,
	})

	refreshPayload := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  (time.Now().Local().Add(time.Hour)).Unix(),
		"sub":  user.Email,
		"typ":  "refresh",
		"role": user.Role,
	})

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("pkg - jwt - CreateJWTs - godotenv: %v", err)
	}
	jwtSigningKey := os.Getenv("JWT_SIGNING_KEY")
	if jwtSigningKey == "" {
		log.Fatal("pkg - jwt - CreateJWTs - env var JWT_SIGNING_KEY not set")
	}

	accessToken, err = accessPayload.SignedString([]byte(jwtSigningKey))
	if err != nil {
		return "", "", fmt.Errorf("pkg - jwt - CreateJWTs - error signing accessJWT: %v", err)
	}

	refreshToken, err = refreshPayload.SignedString([]byte(jwtSigningKey))
	if err != nil {
		return "", "", fmt.Errorf("pkg - jwt - CreateJWTs - error signing refreshJWT: %v", err)
	}

	return accessToken, refreshToken, nil
}

func (s *JWTService) ReadJWT(accessToken string) (claims jwt.MapClaims, err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("pkg - jwt - ValidateJWT - godotenv: %v", err)
	}
	jwtSigningKey := os.Getenv("JWT_SIGNING_KEY")
	if jwtSigningKey == "" {
		log.Fatal("pkg - jwt - ValidateJWT - env var JWT_SIGNING_KEY not set")
	}

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		jwtKey := []byte(jwtSigningKey)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("pkg - jwt - ValidateJWT - unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("pkg - jwt - ValidateJWT - error parsing accessJWT: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("pkg - jwt - ValidateJWT - invalid accessJWT")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("pkg - jwt - ValidateJWT - invalid accessJWT claims")
	}

	return claims, nil
}
