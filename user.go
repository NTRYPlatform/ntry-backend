package auth

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ntryapp/auth/eth"
)

// User is the model for the `user` table
type User struct {

	//TODO: Do we even need this, if we have the mapping?
	EthAddress string `db:"eth_address" json:"ethAddress" binding:"required"`

	SecondaryAddress string `db:"secondary_address" json:"secondaryAddress"`

	Password string `db:"password" json:"password" binding:"required"`

	EmailAddress string `db:"email_address" json:"email" binding:"required"`

	TelephoneNumber string `db:"telephone_number" json:"phone"`

	FirstName string `db:"first_name" json:"firstName"`

	LastName string `db:"last_name" json:"lastName"`

	Address string `db:"address" json:"address"`

	IsEmailVerified bool `db:"email_verified" json:"emailVerified"`

	RegTime time.Time `db:"reg_time" json:"regTime"`

	VerificationCode string `db:"verification_code" json:"verificationCode"`

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

type VerifyUserSignature struct {
	PubKey    string `json:"pubKey" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

//TODO
func RegisterUser(user User) (key string) {
	if UserExistsByUniqueField(&user) == true {
		// TODO: might want to throw exception for better client-side response
		log.Printf("User with either of these values already exists! %v\n", user)
		return
	}

	// create new eth secondary key
	address, key := eth.CreateAccount(user.Password)
	user.SecondaryAddress = address

	// verification
	rand := RandString(40)
	user.VerificationCode = rand
	user.RegTime = time.Now().UTC()
	//TODO: handle exceptions
	if InsertUser(user) {
		SendVerificationEmail(user.EmailAddress, rand)
	}
	return
}

func CompleteUserInfo(user *User) bool {
	err := UpdateUser(user)
	updated := false
	if err != nil {
		log.Printf("Error occurred while trying to update user: %s", err)
	} else {
		updated = true
	}
	return updated
}

//TODO: change name and add checks
func ValidateUser(user *LoginUser) *User {
	return LoginUserValidation(user)
}

func ValidateSecondaryAddress(user *VerifyUserSignature) bool {
	target := GetUserValidationCode(user)
	verified := VerifySignature(user.PubKey, user.Signature, target)
	if verified {
		log.Printf("User with ")
	}
	return verified
}

//TODO
func GetUserByAddress(addr string) *User {
	return nil
}

//TODO
func CheckForEthVerification() {

}
