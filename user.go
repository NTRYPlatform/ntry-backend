package main

import jwt "github.com/dgrijalva/jwt-go"

// UserClaims is the model that will be a part of the JWT
type UserClaims struct {
	//TODO: add others...
	//TODO: Do we even need this, if we have the mapping?
	EthAddress string `json:"ethAddress" binding:"required"`
	//TODO: Doesn't need to be here. Can be replaced with key. Things We Have auth principle.
	Password string `json:"password" binding:"required"`

	EmailAddress string `json:"email" binding:"required"`

	TelephoneNumber string `json:"phone" binding:"required"`

	FullName string `json:"name" binding:"required"`

	// recommended practice to include this
	jwt.StandardClaims
}

//TODO
func RegisterUser() bool {
	return false
}
