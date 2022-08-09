// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AthletesColumns holds the columns for the "athletes" table.
	AthletesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "bio", Type: field.TypeString, Size: 2147483647},
		{Name: "first_name", Type: field.TypeString},
		{Name: "middle_name", Type: field.TypeString, Nullable: true},
		{Name: "last_name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// AthletesTable holds the schema information for the "athletes" table.
	AthletesTable = &schema.Table{
		Name:       "athletes",
		Columns:    AthletesColumns,
		PrimaryKey: []*schema.Column{AthletesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AthletesTable,
	}
)

func init() {
}
