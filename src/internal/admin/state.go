package admin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetStateTable(ctx *context.Context) table.Table {
	states := table.NewDefaultTable(table.Config{
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

	info := states.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int8).
		FieldFilterable().
		FieldSortable()
	info.AddField("Created at", "created_at", db.Timestamptz)
	info.AddField("Updated at", "updated_at", db.Timestamptz)
	info.AddField("Name", "name", db.Varchar)

	info.SetTable("states").
		SetTitle("States").
		SetDescription("Item states")

	formList := states.GetForm()

	formList.AddField("Id", "id", db.Int8, form.Default).
		FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Created at", "created_at", db.Timestamptz, form.Datetime).
		FieldHide().
		FieldHideWhenUpdate().
		FieldNowWhenInsert()
	formList.AddField("Updated at", "updated_at", db.Timestamptz, form.Datetime).
		FieldHide().
		FieldNowWhenUpdate()
	formList.AddField("Name", "name", db.Varchar, form.Text)

	//formList.SetTabGroups(types.
	//	NewTabGroups("id", "created_at", "updated_at", "deleted_at").
	//	AddGroup("name", "description", "serial_number", "status", "state_code").
	//	AddGroup("price", "currency"),
	//)

	formList.SetTable("states").
		SetTitle("States").
		SetDescription("Item states")

	return states
}
