// Code generated by ent, DO NOT EDIT.

package vector

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the vector type in the database.
	Label = "vector"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVector holds the string denoting the vector field in the database.
	FieldVector = "vector"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the vector in the database.
	Table = "vectors"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "vectors"
	// OwnerInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	OwnerInverseTable = "products"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "product_vectors"
)

// Columns holds all SQL columns for vector fields.
var Columns = []string{
	FieldID,
	FieldVector,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "vectors"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"product_vectors",
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

// OrderOption defines the ordering options for the Vector queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}