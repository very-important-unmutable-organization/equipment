package domain

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Equipment struct {
	gorm.Model
	Category        Category `sql:"type:category"`
	Name            string   `sql:"type:varchar;not null"`
	Description     string
	SerialNumber    string   `sql:"not null"`
	TypeCode        int      `sql:"not null"`
	Type            ItemType `gorm:"foreignKey:TypeCode"`
	Status          Status   `sql:"type:item_state"`
	StateCode       int
	State           State           `gorm:"foreignKey:StateCode"`
	PurposeCode     int                    `sql:"not null"`
	Purpose         Purpose         `gorm:"foreignKey:PurposeCode"`
	PurchaseDate    time.Time              `sql:"not null"`
	Price           decimal.Decimal        `sql:"not null"`
	Currency        currency               `sql:"type:currency;default:ruble"`
	OriginCode      int                    `sql:"not null"`
	Origin          Origin          `gorm:"foreignKey:OriginCode"`
	Characteristics map[string]interface{} `sql:"type:jsonb"`
	Docs            []Document             `gorm:"foreignKey:ItemID"`
	Photos          []Photo                `gorm:"foreignKey:ItemID"`
}
