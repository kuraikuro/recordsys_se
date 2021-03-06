// Code generated by entc, DO NOT EDIT.

package registrar

const (
	// Label holds the string label denoting the registrar type in the database.
	Label = "registrar"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"

	// Table holds the table name of the registrar in the database.
	Table = "registrars"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "registrars"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for registrar fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Registrar type.
var ForeignKeys = []string{
	"user_id",
}
