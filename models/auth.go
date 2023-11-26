package models

import (
	"omni/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthModel struct {
	UM UserModel
}

type Claims struct {
	jwt.StandardClaims
}

type Auth struct {
	Token string
}

var jwtKey = []byte("my_secret_key")

func (a AuthModel) LoginUser(u User) (Auth, error) {
	user, err := a.UM.ListUserByEmail(u.Email)
	if err != nil {
		return Auth{}, err
	}

	errHash := utils.CheckPasswordHash(u.Password, user.Password)

	if !errHash {
		return Auth{}, err
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return Auth{}, err
	}

	at := Auth{
		Token: tokenString,
	}

	return at, nil
}

func ParseToken(tokenString string) (claims *Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, tokenErr := token.Claims.(*Claims)

	if !tokenErr {
		return nil, err
	}

	return claims, nil
}
