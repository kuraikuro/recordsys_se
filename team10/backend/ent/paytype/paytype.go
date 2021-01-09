// Code generated by entc, DO NOT EDIT.

package paytype

const (
	// Label holds the string label denoting the paytype type in the database.
	Label = "paytype"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPaytype holds the string denoting the paytype field in the database.
	FieldPaytype = "paytype"

	// EdgeBills holds the string denoting the bills edge name in mutations.
	EdgeBills = "bills"

	// Table holds the table name of the paytype in the database.
	Table = "paytypes"
	// BillsTable is the table the holds the bills relation/edge.
	BillsTable = "bills"
	// BillsInverseTable is the table name for the Bill entity.
	// It exists in this package in order to avoid circular dependency with the "bill" package.
	BillsInverseTable = "bills"
	// BillsColumn is the table column denoting the bills relation/edge.
	BillsColumn = "paytype_id"
)

// Columns holds all SQL columns for paytype fields.
var Columns = []string{
	FieldID,
	FieldPaytype,
}

var (
	// PaytypeValidator is a validator for the "paytype" field. It is called by the builders before save.
	PaytypeValidator func(string) error
)
