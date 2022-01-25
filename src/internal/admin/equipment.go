package admin

import (
	"database/sql"
	"fmt"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"

	"github.com/very-important-unmutable-organization/equipment/internal/domain"
)

func displayStatusValue(value types.FieldModel) interface{} {
	// TODO: fast as fuck!!
	itemType := fmt.Sprintf("%s", value.Row["status"])
	if itemType == string(domain.Free) {
		return "Free"
	}
	if itemType == string(domain.Taken) {
		return "Taken"
	}
	return "-"
}

func displayCategoryValue(value types.FieldModel) interface{} {
	// TODO: fast as fuck!!
	itemType := fmt.Sprintf("%s", value.Row["category"])
	if itemType == string(domain.Furniture) {
		return "Furniture"
	}
	if itemType == string(domain.OfficeEquipment) {
		return "Office Equipment"
	}
	if itemType == string(domain.PersonalEquipment) {
		return "Personal Equipment"
	}
	return "-"
}

func displayCurrencyValue(value types.FieldModel) interface{} {
	// TODO: fast as fuck!!
	itemType := fmt.Sprintf("%s", value.Row["currency"])
	if itemType == string(domain.Ruble) {
		return "₽"
	}
	if itemType == string(domain.USD) {
		return "$"
	}
	if itemType == string(domain.Pound) {
		return "£"
	}
	if itemType == string(domain.Euro) {
		return "€"
	}
	return "-"
}

func displayJSONBValue(value types.FieldModel) interface{} {
	array, ok := value.Row["characteristics"].([]byte)

	if !ok {
		return "-"
	}
	return string(array[:])
}

func displayStateCode(value types.FieldModel) interface{} {
	fmt.Printf("value is %#v\n", value)
	if value.Value == "0" {
		return ""
	}
	return value.Value
}

func GetEquipmentTable(ctx *context.Context) table.Table {
	equipment := table.NewDefaultTable(table.Config{
		Driver:     db.DriverPostgresql,
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: true,
		Connection: table.DefaultConnectionName,
		PrimaryKey: table.PrimaryKey{
			Type: db.Int,
			Name: table.DefaultPrimaryKeyName,
		},
	})

	info := equipment.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int8).
		FieldFilterable().
		FieldSortable()
	info.AddField("Created at", "created_at", db.Timestamptz)
	info.AddField("Updated at", "updated_at", db.Timestamptz)
	info.AddField("Category", "category", db.Enum).
		FieldDisplay(displayCategoryValue)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Description", "description", db.Text)
	info.AddField("Serial number", "serial_number", db.Text)
	info.AddField("Type code", "type_code", db.Int8)
	info.AddField("Status", "status", db.Varchar).
		FieldDisplay(displayStatusValue)
	info.AddField("State code", "state_code", db.Int8).
		FieldDisplay(displayStateCode)
	info.AddField("Purpose code", "purpose_code", db.Int8)
	info.AddField("Purchase date", "purchase_date", db.Timestamptz)
	info.AddField("Price", "price", db.Text)
	info.AddField("Currency", "currency", db.Varchar).
		FieldDisplay(displayCurrencyValue)
	info.AddField("Origin code", "origin_code", db.Int8)
	info.AddField("Characteristics", "characteristics", db.Varchar).
		FieldDisplay(displayJSONBValue)

	info.SetTable("equipment").
		SetTitle("Equipment").
		SetDescription("Equipment")

	formList := equipment.GetForm()

	formList.AddField("Id", "id", db.Int8, form.Default).
		FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Created at", "created_at", db.Timestamptz, form.Datetime).
		FieldHide().
		FieldHideWhenUpdate().
		FieldNowWhenInsert()
	formList.AddField("Updated at", "updated_at", db.Timestamptz, form.Datetime).
		FieldHide().
		FieldNow()
	formList.AddField("Category", "category", db.Enum, form.SelectSingle).
		FieldDisplay(displayCategoryValue).
		FieldMust().
		FieldOptions(types.FieldOptions{
			{Text: "Furniture", Value: string(domain.Furniture)},
			{Text: "Office Equipment", Value: string(domain.OfficeEquipment)},
			{Text: "Personal Equipment", Value: string(domain.PersonalEquipment)},
		})
	formList.AddField("Name", "name", db.Varchar, form.Text).
		FieldMust()
	formList.AddField("Description", "description", db.Text, form.RichText)
	formList.AddField("Serial number", "serial_number", db.Text, form.Text).
		FieldMust()
	formList.AddField("Type code", "type_code", db.Int8, form.Number).
		FieldMust()
	formList.AddField("Status", "status", db.Enum, form.SelectSingle).
		FieldMust().
		FieldOptions(types.FieldOptions{
			{Text: "Free", Value: string(domain.Free)},
			{Text: "Taken", Value: string(domain.Taken)},
		}).
		FieldDisplay(displayStatusValue)
	formList.AddField("State code", "state_code", db.Int, form.Text).
		FieldDefault("").
		FieldDisplay(displayStateCode).
		FieldPostFilterFn(nullableValueFilterFn).
		FieldDefault("")
	formList.AddField("Purpose code", "purpose_code", db.Int8, form.Number).
		FieldMust()
	formList.AddField("Purchase date", "purchase_date", db.Timestamptz, form.Datetime).
		FieldMust()
	formList.AddField("Price", "price", db.Text, form.Text).
		FieldMust()
	formList.AddField("Currency", "currency", db.Enum, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "₽", Value: string(domain.Ruble)},
			{Text: "$", Value: string(domain.USD)},
			{Text: "€", Value: string(domain.Euro)},
			{Text: "£", Value: string(domain.Pound)},
		}).
		FieldDisplay(displayCurrencyValue).
		FieldPlaceholder("₽")
	formList.AddField("Origin code", "origin_code", db.Int8, form.Number).
		FieldMust()
	formList.AddField("Characteristics", "characteristics", db.JSON, form.Text).
		FieldDisplay(displayJSONBValue).FieldPostFilterFn(nullableValueFilterFn)

	//formList.SetTabGroups(types.
	//	NewTabGroups("id", "created_at", "updated_at", "deleted_at").
	//	AddGroup("name", "description", "serial_number", "status", "state_code").
	//	AddGroup("price", "currency"),
	//)

	formList.SetTable("equipment").SetTitle("Equipment").SetDescription("Equipment")

	return equipment
}

func nullableValueFilterFn(value types.PostFieldModel) interface{} {
	if value.Value.Value() == "null" || value.Value.Value() == "" {
		return sql.NullString{}
	}
	return value.Value
}
