package notary

import (
	"fmt"
	"reflect"
	"time"
)

// CarContract is the model for the `user` table
type CarContract struct {
	CID string `db:"cid" json:"cid"  binding:"required"`

	Buyer string `json:"buyer" required:"binding"`

	Seller string `json:"seller" required:"binding"`

	Year int8 `db:"year" json:"year" binding:"required"`

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

	LastUpdateDate *time.Time `db:"last_update_date" json:"lastUpdateDate"`
}

type CarContractUsers struct {
	CID string `db:"cid" json:"cid"  binding:"required"`

	Buyer string `db:"buyer" json:"buyer" required:"binding"`

	Seller string `db:"seller" json:"seller" required:"binding"`
}

func (c *CarContract) GetContractFields() interface{} {

	v := reflect.ValueOf(c)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	fmt.Println(values)
	return values
}
