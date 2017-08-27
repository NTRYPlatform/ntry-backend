package main

import (
	"log"

	jwt "github.com/dgrijalva/jwt-go"
)

// User is the model for the `user` table
type User struct {

	//TODO: Do we even need this, if we have the mapping?
	EthAddress string `db:"eth_address" json:"ethAddress" binding:"required"`

	PubKey string `db:"pub_key"`

	Password string `db:"password" json:"password" binding:"required"`

	EmailAddress string `db:"email_address" json:"email" binding:"required"`

	TelephoneNumber string `db:"telephone_number" json:"phone"`

	FirstName string `db:"first_name" json:"firstName"`

	LastName string `db:"last_name" json:"lastName"`

	Address string `db:"address" json:"address"`

	IsEmailVerified bool `db:"email_verified" json:"emailVerified"`

	EthAddressVerification string `db:"eth_verification" json:"ethVerification"`
}

// UserJWT the custom JWT token
type UserJWT struct {
	User User
	jwt.StandardClaims
}

type LoginUser struct {
	EmailAddress string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
}

//TODO
func RegisterUser(user User) bool {
	if UserExistsByUniqueField(&user) == true {
		log.Println("User with either of these values already exists! %v", user)
		return false
	}
	//TODO: trigger email verification
	return InsertUser(user)
}

//TODO
func ValidateUser(user *LoginUser) *User {
	return LoginUserValidation(user)
}

//TODO
func GetUserByAddress(addr string) *User {
	return nil
}

//TODO
func CheckForEthVerification() {

}
