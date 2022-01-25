package admin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetDocumentTable(ctx *context.Context) table.Table {
	documents := table.NewDefaultTable(table.Config{
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

	info := documents.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int8).
		FieldFilterable().
		FieldSortable()
	info.AddField("Created at", "created_at", db.Timestamptz)
	info.AddField("Updated at", "updated_at", db.Timestamptz)
	info.AddField("Item ID", "item_id", db.Int8)
	info.AddField("Location", "location", db.Varchar)

	info.SetTable("documents").
		SetTitle("Documents").
		SetDescription("Documents related to equipment")

	formList := documents.GetForm()

	formList.AddField("Id", "id", db.Int8, form.Default).
		FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Created at", "created_at", db.Timestamptz, form.Datetime).
		FieldHide().
		FieldHideWhenUpdate().
		FieldNowWhenInsert()
	formList.AddField("Updated at", "updated_at", db.Timestamptz, form.Datetime).
		FieldHide().
		FieldNow()
	formList.AddField("Item ID", "item_id", db.Int8, form.Text)
	formList.AddField("Location", "location", db.Varchar, form.Text)

	formList.SetTable("documents").
		SetTitle("Documents").
		SetDescription("Documents related to equipment")

	return documents
}
