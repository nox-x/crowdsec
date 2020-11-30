// Code generated by entc, DO NOT EDIT.

package decision

import (
	"time"
)

const (
	// Label holds the string label denoting the decision type in the database.
	Label = "decision"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUntil holds the string denoting the until field in the database.
	FieldUntil = "until"
	// FieldScenario holds the string denoting the scenario field in the database.
	FieldScenario = "scenario"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldStartIP holds the string denoting the start_ip field in the database.
	FieldStartIP = "start_ip"
	// FieldEndIP holds the string denoting the end_ip field in the database.
	FieldEndIP = "end_ip"
	// FieldScope holds the string denoting the scope field in the database.
	FieldScope = "scope"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldOrigin holds the string denoting the origin field in the database.
	FieldOrigin = "origin"
	// FieldSimulated holds the string denoting the simulated field in the database.
	FieldSimulated = "simulated"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"

	// Table holds the table name of the decision in the database.
	Table = "decisions"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "decisions"
	// OwnerInverseTable is the table name for the Alert entity.
	// It exists in this package in order to avoid circular dependency with the "alert" package.
	OwnerInverseTable = "alerts"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "alert_decisions"
)

// Columns holds all SQL columns for decision fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUntil,
	FieldScenario,
	FieldType,
	FieldStartIP,
	FieldEndIP,
	FieldScope,
	FieldValue,
	FieldOrigin,
	FieldSimulated,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Decision type.
var ForeignKeys = []string{
	"alert_decisions",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// DefaultSimulated holds the default value on creation for the simulated field.
	DefaultSimulated bool
)