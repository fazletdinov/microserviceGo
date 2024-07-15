package grpcservice

import (
	"auth/internal/models"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func GenerateAccessToken(user *models.Users, pathSecret string, expiry int) (string, error) {
	privateKeyBytes, errPath := os.ReadFile(pathSecret)
	if errPath != nil {
		return "", errPath
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return "", errPath
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Minute * time.Duration(expiry)).Unix(),
	})

	signedAccessToken, _ := accessToken.SignedString(privateKey)

	return signedAccessToken, nil
}

func GenerateRefreshToken(user *models.Users, pathSecret string, expiry int) (string, error) {
	privateKeyBytes, errPath := os.ReadFile(pathSecret)
	if errPath != nil {
		return "", errPath
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return "", errPath
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * time.Duration(expiry)).Unix(),
	})

	signedRefreshToken, _ := refreshToken.SignedString(privateKey)

	return signedRefreshToken, nil
}

func ExtractUserIDFromToken(tokenString string, pathSecret string) (uuid.UUID, error) {
	publicKeyBytes, errPath := os.ReadFile(pathSecret)
	if errPath != nil {
		return uuid.Nil, errPath
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return uuid.Nil, errPath
	}

	parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("неожиданный метод подписи - %v", t.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		userIdValue, ok := claims["sub"].(string)
		if !ok {
			return uuid.Nil, fmt.Errorf("невозможно привести user_id к UUID")
		}
		return uuid.MustParse(userIdValue), nil
	} else {
		return uuid.Nil, fmt.Errorf("недействительный токен")
	}

}

func IsAuthorized(requestToken string, pathSecret string) (bool, error) {
	publicKeyBytes, errPath := os.ReadFile(pathSecret)
	if errPath != nil {
		return false, errPath
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return false, errPath
	}
	_, err = jwt.Parse(requestToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("неожиданный метод подписи - %v", t.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
