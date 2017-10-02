package eth

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type ContractNotification struct {
	Contract    CarContract
	NotifyParty string
}

type ContractFields struct {
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Placeholder interface{} `json:"default"`
}

// CarContract is the model for the `user` table
type CarContract struct {
	CID int64 `db:"cid" json:"cid"  binding:"required"`

	Buyer string `db:"buyer" json:"buyer" required:"binding"`

	Seller string `db:"seller" json:"seller" required:"binding"`

	Year int `db:"year" json:"year" binding:"required"`

	Make string `db:"make" json:"make" binding:"required"`

	Model string `db:"model" json:"model" binding:"required"`

	VIN string `db:"vin" json:"vin" binding:"required"`

	Type string `db:"type" json:"type" binding:"required"`

	Color string `db:"color" json:"color,omitempty"`

	EngineNo string `db:"engine_no" json:"engineNo,omitempty"`

	Mileage int `db:"mileage" json:"mileage"`

	TotalPrice int `db:"total_price" json:"totalPrice"`

	DownPayment int `db:"down_payment" json:"downPayment"`

	RemainingPayment int `db:"remaining_payment" json:"remainingPayment"`

	CreationDate *time.Time `db:"creation_date" json:"creationDate"`

	RemainingPaymentDate *time.Time `db:"remaining_payment_date" json:"remainingPaymentDate"`

	ContentHash string `db:"content_hash" json:"hash"`

	Approved bool `db:"approved" json:"approved"`

	LastUpdateDate *time.Time `db:"last_updated_date" json:"lastUpdateDate"`
}

// GetContractFields
func GetContractFields() interface{} {
	t := time.Now().UTC()
	c := CarContract{Buyer: "Buyer's ID", Seller: "Seller's ID", Year: 2016,
		Make: "Tesla", Model: "Model X", VIN: "1HGBH41JXMN109186", Type: "Sedan",
		Color: "Gun Metal", EngineNo: "17100H0203611", Mileage: 23420, TotalPrice: 65450,
		DownPayment: 5000, RemainingPayment: 60450, RemainingPaymentDate: &t}

	v := reflect.TypeOf(c)
	cv := reflect.ValueOf(c)
	var f []ContractFields

	for i := 0; i < v.NumField(); i++ {
		reqd := true
		fi := v.Field(i)
		val := cv.Field(i).Interface()
		cf := ContractFields{Name: fi.Name}

		switch ft := val.(type) {
		case int:
			if ft != 0 {
				cf.Type = "number"
				cf.Placeholder = strconv.Itoa(ft)
			} else {
				reqd = false
			}
		case *time.Time:
			if ft != nil {
				cf.Type = "datetime"
				cf.Placeholder = ft.String()
			} else {
				reqd = false
			}
		case string:
			if len(ft) > 1 {
				cf.Type = "string"
				cf.Placeholder = string(ft)
			} else {
				reqd = false
			}
		default:
			fmt.Println("I don't know, ask stackoverflow.")
		}
		if reqd {
			f = append(f, cf)
		}
	}

	return f
}

func (c *CarContract) String() string {
	return fmt.Sprintf(`
		CID: %v
		Seller: %v
		Buyer: %v
		Year: %v
		Make: %v
		Model: %v
		VIN: %v
		Type: %v 
		Color: %v
		EngineNo: %v
		Mileage: %v
		TotalPrice: %v
		DownPayment: %v
		RemainingPayment: %v
		RemainingPaymentDate: %v
		`,
		c.CID,
		c.Seller,
		c.Buyer,
		c.Year,
		c.Make,
		c.Model,
		c.VIN,
		c.Type,
		c.Color,
		c.EngineNo,
		c.Mileage,
		c.TotalPrice,
		c.DownPayment,
		c.RemainingPayment,
		c.RemainingPaymentDate)
}

func (c *CarContract) Hash() string {
	h := sha256.New()
	h.Write([]byte(c.String()))
	return hex.EncodeToString(h.Sum(nil))
}
