package admin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

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
	info.AddField("Deleted at", "deleted_at", db.Timestamptz)
	info.AddField("Category", "category", db.Varchar)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Description", "description", db.Text)
	info.AddField("Serial number", "serial_number", db.Text)
	info.AddField("Type code", "type_code", db.Int8)
	info.AddField("Status", "status", db.Varchar)
	info.AddField("State code", "state_code", db.Int8)
	info.AddField("Purpose code", "purpose_code", db.Int8)
	info.AddField("Purchase date", "purchase_date", db.Timestamptz)
	info.AddField("Price", "price", db.Text)
	info.AddField("Currency", "currency", db.Varchar)
	info.AddField("Origin code", "origin_code", db.Int8)
	info.AddField("Characteristics", "characteristics", db.Varchar)

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
	formList.AddField("Deleted at", "deleted_at", db.Timestamptz, form.Datetime).
		FieldHide().
		FieldNowWhenInsert()
	formList.AddField("Category", "category", db.Enum, form.SelectSingle).
		FieldMust().
		FieldPlaceholder("penis").
		FieldOptions(types.FieldOptions{
			{Text: "Furniture", Value: "furniture"},
			{Text: "Office Equipment", Value: "office_equipment"},
			{Text: "Personal Equipment", Value: "personal_equipment"},
		})
	formList.AddField("Name", "name", db.Varchar, form.Text).
		FieldMust()
	formList.AddField("Description", "description", db.Text, form.RichText)
	formList.AddField("Serial number", "serial_number", db.Text, form.RichText)
	formList.AddField("Type code", "type_code", db.Int8, form.Number).
		FieldMust()
	formList.AddField("Status", "status", db.Enum, form.SelectSingle).
		FieldMust().
		FieldOptions(types.FieldOptions{
			{Text: "Free", Value: "free"},
			{Text: "Taken", Value: "taken"},
		})
	formList.AddField("State code", "state_code", db.Int8, form.Number)
	formList.AddField("Purpose code", "purpose_code", db.Int8, form.Number).
		FieldMust()
	formList.AddField("Purchase date", "purchase_date", db.Timestamptz, form.Datetime).
		FieldMust()
	formList.AddField("Price", "price", db.Text, form.Text).
		FieldMust()
	formList.AddField("Currency", "currency", db.Enum, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "₽", Value: "ruble"},
			{Text: "$", Value: "usd"},
			{Text: "€", Value: "euro"},
			{Text: "£", Value: "pound"},
		})
	formList.AddField("Origin code", "origin_code", db.Int8, form.Number).
		FieldMust()
	formList.AddField("Characteristics", "characteristics", db.JSON, form.Text)

	//formList.SetTabGroups(types.
	//	NewTabGroups("id", "created_at", "updated_at", "deleted_at").
	//	AddGroup("name", "description", "serial_number", "status", "state_code").
	//	AddGroup("price", "currency"),
	//)

	formList.SetTable("equipment").SetTitle("Equipment").SetDescription("Equipment")

	return equipment
}
