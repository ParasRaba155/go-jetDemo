//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package enum

import "github.com/go-jet/jet/v2/postgres"

var Role = &struct {
	User  postgres.StringExpression
	Admin postgres.StringExpression
}{
	User:  postgres.NewEnumValue("user"),
	Admin: postgres.NewEnumValue("admin"),
}
