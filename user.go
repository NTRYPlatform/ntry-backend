package notary

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// User is the model for the `user` table
type User struct {
	UID string `db:"uid" json:"uid"  binding:"required"`

	EthAddress string `db:"eth_address" json:"ethAddress"`

	Password string `db:"password" json:"password" binding:"required"`

	EmailAddress string `db:"email_address" json:"email" binding:"required"`

	TelephoneNumber string `db:"telephone_number" json:"phone"`

	FirstName string `db:"first_name" json:"firstName"`

	LastName string `db:"last_name" json:"lastName"`

	Address string `db:"address" json:"address"`

	AccountVerified bool `db:"account_verified" json:"accountVerified"`

	RegTime time.Time `db:"reg_time" json:"regTime"`

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

// VerifyUser sets verification info
func VerifyUser(uid, address, txHash string) *User {
	return &User{EthAddress: address, EthAddressVerification: txHash, AccountVerified: true}
}

//TODO
func GetUserByAddress(uid string) (*User, error) {
	return nil, nil
}

func (u *User) OK() error {
	if len(u.EthAddress) < 32 {
		// return ErrRequied("eth address")
	}
	// if ..
	return nil
}

func (u *User) String() string {
	return fmt.Sprintf(`
	UID:     %v
	EthAddress:     %v
	Password:     %v
	EmailAddress:       %v
	TelephoneNumber:    %v
	FirstName: %v
	LastName  %v
	Address:    %v
	AccountVerified:     %v 
	RegTime:        %v
	EthAddressVerification:        %v
	`,
		u.UID,
		u.EthAddress,
		u.Password,
		u.EmailAddress,
		u.TelephoneNumber,
		u.FirstName,
		u.LastName,
		u.Address,
		u.AccountVerified,
		u.RegTime,
		u.EthAddressVerification)
}
