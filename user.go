package notary

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"

	jwt "github.com/dgrijalva/jwt-go"
)

// User is the model for the `user` table
type User struct {
	UID string `db:"uid" json:"uid"  binding:"required"`

	EthAddress string `db:"eth_address" json:"ethAddress"`

	Password string `db:"password" json:"password" binding:"required"`

	EmailAddress string `db:"email_address" json:"email" binding:"required"`

	TelephoneNumber string `db:"telephone_number,omitempty" json:"phone"`

	FirstName string `db:"first_name,omitempty" json:"firstName"`

	LastName string `db:"last_name,omitempty" json:"lastName"`

	Address string `db:"address,omitempty" json:"address"`

	AccountVerified bool `db:"account_verified" json:"accountVerified"`

	RegTime time.Time `db:"reg_time" json:"regTime"`

	EthAddressVerification string `db:"eth_verification,omitempty" json:"ethVerification"`
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

// OK validates LoginUser
func (u *LoginUser) OK() error {
	if len(u.Password) == 0 {
		return &ErrRequired{arg: "Password"}
	}
	if len(u.EmailAddress) == 0 {
		return &ErrRequired{arg: "Email Address"}
	}
	return nil
}

// OK validates User
func (u *User) OK() error {
	// mandatory values
	r := NewRegexUtil()
	if len(u.EmailAddress) == 0 {
		return &ErrRequired{arg: "Email Address"}
	} else if !r.MatchEmail(u.EmailAddress) {
		return &ErrInvalidValue{arg: "Email Address"}
	}
	if len(u.Password) == 0 {
		return &ErrRequired{arg: "Password"}
	}
	if len(u.UID) == 0 {
		return &ErrRequired{arg: "Password"}
	}
	// non-mandatory values
	if !(len(u.EthAddress) == 0) {
		defer func() error {
			if r := recover(); r != nil {
				return &ErrInvalidValue{arg: "Ethereum Address"}
			}
			return nil
		}()
		common.StringToAddress(u.EthAddress)
	}
	//TODO: add other non-mandatory values used later on?
	return nil
}

func (u *User) String() string {
	return fmt.Sprintf(`
	UID:     %v
	EthAddress:     %v
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
		u.EmailAddress,
		u.TelephoneNumber,
		u.FirstName,
		u.LastName,
		u.Address,
		u.AccountVerified,
		u.RegTime,
		u.EthAddressVerification)
}
