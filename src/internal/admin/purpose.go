package admin

import (
	"fmt"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"

	"github.com/very-important-unmutable-organization/equipment/internal/domain"
)

func displayPurposeTypeValue(value types.FieldModel) interface{} {
	// TODO: fast as fuck!!
	itemType := fmt.Sprintf("%s", value.Row["type"])
	if itemType == string(domain.Personal) {
		return "Personal"
	}
	if itemType == string(domain.General) {
		return "General"
	}
	if itemType == string(domain.Testing) {
		return "Testing"
	}
	return "-"
}

func GetPurposeTable(ctx *context.Context) table.Table {
	purposes := table.NewDefaultTable(table.Config{
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

	info := purposes.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int8).
		FieldFilterable().
		FieldSortable()
	info.AddField("Created at", "created_at", db.Timestamptz)
	info.AddField("Updated at", "updated_at", db.Timestamptz)
	info.AddField("Type", "type", db.Enum).
		FieldDisplay(displayPurposeTypeValue)
	info.AddField("ResponsibleEmployeeUID", "responsible_employee_uid", db.UUID)
	info.SetTable("purposes").
		SetTitle("Purpose").
		SetDescription("Equipment purposes")

	formList := purposes.GetForm()

	formList.AddField("Id", "id", db.Int8, form.Default).
		FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Created at", "created_at", db.Timestamptz, form.Datetime).
		FieldHide().
		FieldHideWhenUpdate().
		FieldNowWhenInsert()
	formList.AddField("Updated at", "updated_at", db.Timestamptz, form.Datetime).
		FieldHide().
		FieldNow()
	formList.AddField("Type", "type", db.Enum, form.SelectSingle).
		FieldPlaceholder("Personal").
		FieldOptions(types.FieldOptions{
			{Text: "Personal", Value: string(domain.Personal)},
			{Text: "General", Value: string(domain.General)},
			{Text: "Testing", Value: string(domain.Testing)},
		}).
		FieldDisplay(displayPurposeTypeValue)
	formList.AddField("ResponsibleEmployeeUID", "responsible_employee_uid", db.UUID, form.Text).
		FieldMust()

	formList.SetTable("purposes").
		SetTitle("Purpose").
		SetDescription("Equipment purposes")

	return purposes
}
