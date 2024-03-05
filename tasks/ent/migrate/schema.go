// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "price", Type: field.TypeInt64},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"todo", "in_progress", "done"}, Default: "todo"},
		{Name: "task_user", Type: field.TypeInt},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tasks_users_user",
				Columns:    []*schema.Column{TasksColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "role", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserLogsColumns holds the columns for the "user_logs" table.
	UserLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "idempotency_key", Type: field.TypeString, Unique: true},
	}
	// UserLogsTable holds the schema information for the "user_logs" table.
	UserLogsTable = &schema.Table{
		Name:       "user_logs",
		Columns:    UserLogsColumns,
		PrimaryKey: []*schema.Column{UserLogsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TasksTable,
		UsersTable,
		UserLogsTable,
	}
)

func init() {
	TasksTable.ForeignKeys[0].RefTable = UsersTable
}