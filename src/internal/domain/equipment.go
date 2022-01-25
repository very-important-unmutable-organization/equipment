package domain

import (
	"time"

	"github.com/shopspring/decimal"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

var _ error = ErrorEquipmentTaken{}

type ErrorEquipmentTaken struct {
}

func (e ErrorEquipmentTaken) Error() string {
	return "Equipment is already taken"
}

var _ error = ErrorEquipmentFree{}

type ErrorEquipmentFree struct {
}

func (e ErrorEquipmentFree) Error() string {
	return "Equipment is already free"
}

type Equipment struct {
	gorm.Model
	Category        Category `gorm:"type:category;not null"`
	Name            string   `gorm:"type:varchar;not null"`
	Description     string
	SerialNumber    string          `gorm:"not null" json:"serial_number"`
	TypeCode        int             `gorm:"not null" json:"type_code"`
	Type            ItemType        `gorm:"foreignKey:TypeCode"`
	Status          Status          `gorm:"type:status;not null"`
	StateCode       null.Int        `json:"state_code"`
	State           State           `gorm:"foreignKey:StateCode"`
	PurposeCode     int             `gorm:"not null" json:"purpose_code"`
	Purpose         Purpose         `gorm:"foreignKey:PurposeCode"`
	PurchaseDate    time.Time       `gorm:"not null" json:"purchase_date"`
	Price           decimal.Decimal `gorm:"not null"`
	Currency        currency        `gorm:"type:currency;default:ruble"`
	OriginCode      int             `gorm:"not null" json:"origin_code"`
	Origin          Origin          `gorm:"foreignKey:OriginCode"`
	Characteristics JSONB           `gorm:"type:jsonb"`
	Docs            []Document      `gorm:"foreignKey:ItemID"`
	Photos          []Photo         `gorm:"foreignKey:ItemID"`
}
