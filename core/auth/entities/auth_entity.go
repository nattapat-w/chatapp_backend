package entities

import "github.com/golang-jwt/jwt/v4"

type UsersClaims struct {
	Id       int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
