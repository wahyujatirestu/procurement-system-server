package models

import "github.com/golang-jwt/jwt/v5"

type JwtPayloadClaim struct {
	jwt.RegisteredClaims
	UserId uint   `json:"user_id"`
	Role   string `json:"role"`
}