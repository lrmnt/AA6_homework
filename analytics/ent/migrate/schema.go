// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BillingOperationsColumns holds the columns for the "billing_operations" table.
	BillingOperationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "amount", Type: field.TypeInt64},
		{Name: "timestamp", Type: field.TypeInt64},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"payout", "debt_move"}},
		{Name: "billing_operations_user", Type: field.TypeInt},
	}
	// BillingOperationsTable holds the schema information for the "billing_operations" table.
	BillingOperationsTable = &schema.Table{
		Name:       "billing_operations",
		Columns:    BillingOperationsColumns,
		PrimaryKey: []*schema.Column{BillingOperationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "billing_operations_users_user",
				Columns:    []*schema.Column{BillingOperationsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "cost", Type: field.TypeInt64},
		{Name: "timestamp", Type: field.TypeInt64},
		{Name: "jira_id", Type: field.TypeString},
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
				Columns:    []*schema.Column{TasksColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TaskLogsColumns holds the columns for the "task_logs" table.
	TaskLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "idempotency_key", Type: field.TypeString, Unique: true},
	}
	// TaskLogsTable holds the schema information for the "task_logs" table.
	TaskLogsTable = &schema.Table{
		Name:       "task_logs",
		Columns:    TaskLogsColumns,
		PrimaryKey: []*schema.Column{TaskLogsColumns[0]},
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
		BillingOperationsTable,
		TasksTable,
		TaskLogsTable,
		UsersTable,
		UserLogsTable,
	}
)

func init() {
	BillingOperationsTable.ForeignKeys[0].RefTable = UsersTable
	TasksTable.ForeignKeys[0].RefTable = UsersTable
}