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

func displayOriginTypeValue(value types.FieldModel) interface{} {
	// TODO: fast as fuck!!
	itemType := fmt.Sprintf("%s", value.Row["type"])
	if itemType == string(domain.CompanyProperty) {
		return "Company property"
	}
	if itemType == string(domain.EmployeeProperty) {
		return "Employee property"
	}
	return "-"
}

func GetOriginTable(ctx *context.Context) table.Table {
	origins := table.NewDefaultTable(table.Config{
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

	info := origins.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int8).
		FieldFilterable().
		FieldSortable()
	info.AddField("Created at", "created_at", db.Timestamptz)
	info.AddField("Updated at", "updated_at", db.Timestamptz)
	info.AddField("Origin type", "type", db.Enum).
		FieldDisplay(displayOriginTypeValue)
	info.AddField("EmployeeUID", "employee_uid", db.UUID)
	info.SetTable("origins").
		SetTitle("Origins").
		SetDescription("Equipment origins")

	formList := origins.GetForm()

	formList.AddField("Id", "id", db.Int8, form.Default).
		FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Created at", "created_at", db.Timestamptz, form.Datetime).
		FieldHide().
		FieldHideWhenUpdate().
		FieldNowWhenInsert()
	formList.AddField("Updated at", "updated_at", db.Timestamptz, form.Datetime).
		FieldHide().
		FieldNow()
	formList.AddField("Origin type", "type", db.Varchar, form.SelectSingle).
		FieldPlaceholder("-").
		FieldOptions(types.FieldOptions{
			{Text: "Company property", Value: string(domain.CompanyProperty)},
			{Text: "Employee property", Value: string(domain.EmployeeProperty)},
		}).
		FieldMust().
		FieldDisplay(displayOriginTypeValue)
	formList.AddField("EmployeeUID", "employee_uid", db.UUID, form.Text)

	formList.SetTable("origins").
		SetTitle("Origins").
		SetDescription("Equipment origins")

	return origins
}
