package notary

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/NTRYPlatform/ntry-backend/eth"
	"github.com/ethereum/go-ethereum/common"
)

// User is the model for the `user` table
type User struct {
	UID string `db:"uid" json:"uid"  binding:"required"`

	EthAddress string `db:"eth_address" json:"ethAddress"`

	//TODO: Need to make sure pwd isn't ever serialized, maybe custom marshaller?
	Password string `db:"password" json:"password,omitempty" binding:"required"`

	EmailAddress string `db:"email_address" json:"email" binding:"required"`

	TelephoneNumber string `db:"telephone_number" json:"phone"`

	FirstName string `db:"first_name" json:"firstName"`

	LastName string `db:"last_name" json:"lastName"`

	Country string `db:"country" json:"country"`

	City string `db:"city" json:"city"`

	State string `db:"state" json:"state"`

	Avatar string `db:"avatar" json:"-"`

	Address string `db:"address" json:"address"`

	AccountVerified bool `db:"account_verified" json:"accountVerified"`

	RegTime *time.Time `db:"reg_time" json:"regTime"`

	EthAddressVerification string `db:"eth_verification" json:"ethVerification,omitempty"`
}

// UserContact the custom JWT token
type UserContact struct {
	PUid string `db:"p_uid"`
	SUid string `db:"s_uid"`
}

type LoginUser struct {
	EmailAddress string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
}

type ForgottenPassword struct {
	UID          string `db:"uid" json:"-"  binding:"required"`
	Timestamp    int64  `db:"timestamp" json:"-"`
	TempPassword string `db:"temp_password" json:"-" binding:"required"`
}

type ChangePasswordUser struct {
	EmailAddress     string `json:"email" binding:"required"`
	PreviousPassword string `json:"prevPassword" binding:"required"`
	NewPassword      string `json:"newPassword" binding:"required"`
}

type ForgottenPasswordWithUser struct {
	User              `db:",inline"`
	ForgottenPassword `db:",inline"`
}

// VerifyUser sets verification info
func VerifyUser(uid, address, txHash string) *User {
	return &User{UID: uid, EthAddress: address, EthAddressVerification: txHash, AccountVerified: true}
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

func (u *ChangePasswordUser) OK() error {
	if len(u.PreviousPassword) == 0 {
		return &ErrRequired{arg: "Prev Password Required"}
	}
	if len(u.EmailAddress) == 0 {
		return &ErrRequired{arg: "Email Address"}
	}
	if len(u.NewPassword) == 0 {
		return &ErrRequired{arg: "New Password Required"}
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
		return &ErrRequired{arg: "UID"}
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

func (u *User) MarshalJSON() ([]byte, error) {
	type user User // prevent recursion
	x := user(*u)
	x.Password = ""
	return json.Marshal(x)
}

type UserContractsWrapper struct {
	url string              `json:"etherscanURL"`
	con []eth.UserContracts `json:"contracts"`
}
