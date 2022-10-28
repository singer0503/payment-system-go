package auth

import (
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("Here need a set of secret strings, replace him")

type JWTClaim struct {
	Unique_name string `json:"unique_name"`
	Role        string `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(userID string, role string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Unique_name: userID,
		Role:        role,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Local().Unix(),
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	// replace "Bearer " to "" in signedToken
	if strings.Index(signedToken, "Bearer ") != -1 {
		signedToken = signedToken[7:]
	}
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	return

}
