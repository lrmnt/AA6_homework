// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "user_role", Type: field.TypeInt},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_roles_role",
				Columns:    []*schema.Column{UsersColumns[3]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		RolesTable,
		UsersTable,
	}
)

func init() {
	UsersTable.ForeignKeys[0].RefTable = RolesTable
}
