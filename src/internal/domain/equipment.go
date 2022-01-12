package domain

import (
	"time"

	"github.com/shopspring/decimal"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type Equipment struct {
	gorm.Model
	Category        Category `gorm:"type:category;not null"`
	Name            string   `gorm:"type:varchar;not null"`
	Description     string
	SerialNumber    string   `gorm:"not null"`
	TypeCode        int      `gorm:"not null"`
	Type            ItemType `gorm:"foreignKey:TypeCode"`
	Status          Status   `gorm:"type:status"`
	StateCode       int
	StateCode       null.Int
	State           State           `gorm:"foreignKey:StateCode"`
	PurposeCode     int             `gorm:"not null"`
	Purpose         Purpose         `gorm:"foreignKey:PurposeCode"`
	PurchaseDate    time.Time       `gorm:"not null"`
	Price           decimal.Decimal `gorm:"not null"`
	Currency        currency        `gorm:"type:currency;default:ruble"`
	OriginCode      int             `gorm:"not null"`
	Origin          Origin          `gorm:"foreignKey:OriginCode"`
	Characteristics JSONB           `gorm:"type:jsonb"`
	Docs            []Document      `gorm:"foreignKey:ItemID"`
	Photos          []Photo         `gorm:"foreignKey:ItemID"`
}
