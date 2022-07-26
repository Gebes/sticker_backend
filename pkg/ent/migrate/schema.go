// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// StickersColumns holds the columns for the "stickers" table.
	StickersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "location_description", Type: field.TypeString},
		{Name: "latitude", Type: field.TypeFloat64},
		{Name: "longitude", Type: field.TypeFloat64},
		{Name: "edition", Type: field.TypeEnum, Enums: []string{"original", "original_with_earring", "winter", "train", "mail", "holiday", "other"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "sticker_owner", Type: field.TypeString},
	}
	// StickersTable holds the schema information for the "stickers" table.
	StickersTable = &schema.Table{
		Name:       "stickers",
		Columns:    StickersColumns,
		PrimaryKey: []*schema.Column{StickersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "stickers_users_owner",
				Columns:    []*schema.Column{StickersColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		StickersTable,
		UsersTable,
	}
)

func init() {
	StickersTable.ForeignKeys[0].RefTable = UsersTable
}
