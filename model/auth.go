package model

import "github.com/golang-jwt/jwt/v5"

type AuthClaimJwt struct {
	jwt.RegisteredClaims
	UserId    int    `json:"user"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
	UserScopes []string `json:"user_scopes"`
}