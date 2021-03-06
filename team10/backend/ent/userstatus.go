// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team10/app/ent/userstatus"
)

// Userstatus is the model entity for the Userstatus schema.
type Userstatus struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Userstatus holds the value of the "Userstatus" field.
	Userstatus string `json:"Userstatus,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserstatusQuery when eager-loading is set.
	Edges UserstatusEdges `json:"edges"`
}

// UserstatusEdges holds the relations/edges for other nodes in the graph.
type UserstatusEdges struct {
	// User holds the value of the user edge.
	User []*User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e UserstatusEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Userstatus) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Userstatus
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Userstatus fields.
func (u *Userstatus) assignValues(values ...interface{}) error {
	if m, n := len(values), len(userstatus.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Userstatus", values[0])
	} else if value.Valid {
		u.Userstatus = value.String
	}
	return nil
}

// QueryUser queries the user edge of the Userstatus.
func (u *Userstatus) QueryUser() *UserQuery {
	return (&UserstatusClient{config: u.config}).QueryUser(u)
}

// Update returns a builder for updating this Userstatus.
// Note that, you need to call Userstatus.Unwrap() before calling this method, if this Userstatus
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *Userstatus) Update() *UserstatusUpdateOne {
	return (&UserstatusClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *Userstatus) Unwrap() *Userstatus {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: Userstatus is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *Userstatus) String() string {
	var builder strings.Builder
	builder.WriteString("Userstatus(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", Userstatus=")
	builder.WriteString(u.Userstatus)
	builder.WriteByte(')')
	return builder.String()
}

// Userstatuses is a parsable slice of Userstatus.
type Userstatuses []*Userstatus

func (u Userstatuses) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
