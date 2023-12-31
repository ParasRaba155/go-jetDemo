//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Users = newUsersTable("public", "users", "")

type usersTable struct {
	postgres.Table

	// Columns
	Email        postgres.ColumnString
	FullName     postgres.ColumnString
	Intro        postgres.ColumnString
	Profile      postgres.ColumnString
	Role         postgres.ColumnString
	PasswordHash postgres.ColumnString
	RegisteredAt postgres.ColumnTimestamp
	UpdatedAt    postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type UsersTable struct {
	usersTable

	EXCLUDED usersTable
}

// AS creates new UsersTable with assigned alias
func (a UsersTable) AS(alias string) *UsersTable {
	return newUsersTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new UsersTable with assigned schema name
func (a UsersTable) FromSchema(schemaName string) *UsersTable {
	return newUsersTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new UsersTable with assigned table prefix
func (a UsersTable) WithPrefix(prefix string) *UsersTable {
	return newUsersTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new UsersTable with assigned table suffix
func (a UsersTable) WithSuffix(suffix string) *UsersTable {
	return newUsersTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newUsersTable(schemaName, tableName, alias string) *UsersTable {
	return &UsersTable{
		usersTable: newUsersTableImpl(schemaName, tableName, alias),
		EXCLUDED:   newUsersTableImpl("", "excluded", ""),
	}
}

func newUsersTableImpl(schemaName, tableName, alias string) usersTable {
	var (
		EmailColumn        = postgres.StringColumn("email")
		FullNameColumn     = postgres.StringColumn("full_name")
		IntroColumn        = postgres.StringColumn("intro")
		ProfileColumn      = postgres.StringColumn("profile")
		RoleColumn         = postgres.StringColumn("role")
		PasswordHashColumn = postgres.StringColumn("password_hash")
		RegisteredAtColumn = postgres.TimestampColumn("registered_at")
		UpdatedAtColumn    = postgres.TimestampColumn("updated_at")
		allColumns         = postgres.ColumnList{EmailColumn, FullNameColumn, IntroColumn, ProfileColumn, RoleColumn, PasswordHashColumn, RegisteredAtColumn, UpdatedAtColumn}
		mutableColumns     = postgres.ColumnList{FullNameColumn, IntroColumn, ProfileColumn, RoleColumn, PasswordHashColumn, RegisteredAtColumn, UpdatedAtColumn}
	)

	return usersTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		Email:        EmailColumn,
		FullName:     FullNameColumn,
		Intro:        IntroColumn,
		Profile:      ProfileColumn,
		Role:         RoleColumn,
		PasswordHash: PasswordHashColumn,
		RegisteredAt: RegisteredAtColumn,
		UpdatedAt:    UpdatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
