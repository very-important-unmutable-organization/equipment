package domain

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type Equipment struct {
	gorm.Model
	Category        Category `sql:"type:category"`
	Name            string   `sql:"type:varchar;not null"`
	Description     string
	SerialNumber    string   `sql:"not null"`
	TypeCode        int      `sql:"not null"`
	Type            ItemType `sql:"foreignKey:TypeCode"`
	Status          Status   `sql:"type:item_state"`
	StateCode       int
	State           State                  `sql:"foreignKey:StateCode"`
	PurposeCode     int                    `sql:"not null"`
	Purpose         Purpose                `sql:"foreignKey:PurposeCode"`
	PurchaseDate    time.Time              `sql:"not null"`
	Price           decimal.Decimal        `sql:"not null"`
	Currency        currency               `sql:"type:currency;default:ruble"`
	OriginCode      int                    `sql:"not null"`
	Origin          Origin                 `sql:"foreignKey:OriginCode"`
	Characteristics map[string]interface{} `sql:"type:jsonb"`
	Docs            []Document             `gorm:"foreignKey:ItemID"`
	Photos          []Photo                `gorm:"foreignKey:ItemID"`
}
