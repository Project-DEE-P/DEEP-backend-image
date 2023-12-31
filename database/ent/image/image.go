// Code generated by ent, DO NOT EDIT.

package image

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the image type in the database.
	Label = "image"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldInstance holds the string denoting the instance field in the database.
	FieldInstance = "instance"
	// Table holds the table name of the image in the database.
	Table = "images"
)

// Columns holds all SQL columns for image fields.
var Columns = []string{
	FieldID,
	FieldInstance,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// InstanceValidator is a validator for the "instance" field. It is called by the builders before save.
	InstanceValidator func([]byte) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Image queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}
