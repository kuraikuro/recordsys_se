// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team10/app/ent/financier"
	"github.com/team10/app/ent/medicalrecordstaff"
	"github.com/team10/app/ent/nurse"
	"github.com/team10/app/ent/patientrights"
	"github.com/team10/app/ent/registrar"
	"github.com/team10/app/ent/user"
	"github.com/team10/app/ent/userstatus"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges         UserEdges `json:"edges"`
	user_id       *int
	userstatus_id *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Financier holds the value of the financier edge.
	Financier *Financier
	// Historytaking holds the value of the historytaking edge.
	Historytaking *Nurse
	// UserPatientrights holds the value of the UserPatientrights edge.
	UserPatientrights *Patientrights
	// Medicalrecordstaff holds the value of the medicalrecordstaff edge.
	Medicalrecordstaff *Medicalrecordstaff
	// User2registrar holds the value of the user2registrar edge.
	User2registrar *Registrar
	// Userstatus holds the value of the userstatus edge.
	Userstatus *Userstatus
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// FinancierOrErr returns the Financier value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) FinancierOrErr() (*Financier, error) {
	if e.loadedTypes[0] {
		if e.Financier == nil {
			// The edge financier was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: financier.Label}
		}
		return e.Financier, nil
	}
	return nil, &NotLoadedError{edge: "financier"}
}

// HistorytakingOrErr returns the Historytaking value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) HistorytakingOrErr() (*Nurse, error) {
	if e.loadedTypes[1] {
		if e.Historytaking == nil {
			// The edge historytaking was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: nurse.Label}
		}
		return e.Historytaking, nil
	}
	return nil, &NotLoadedError{edge: "historytaking"}
}

// UserPatientrightsOrErr returns the UserPatientrights value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) UserPatientrightsOrErr() (*Patientrights, error) {
	if e.loadedTypes[2] {
		if e.UserPatientrights == nil {
			// The edge UserPatientrights was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: patientrights.Label}
		}
		return e.UserPatientrights, nil
	}
	return nil, &NotLoadedError{edge: "UserPatientrights"}
}

// MedicalrecordstaffOrErr returns the Medicalrecordstaff value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) MedicalrecordstaffOrErr() (*Medicalrecordstaff, error) {
	if e.loadedTypes[3] {
		if e.Medicalrecordstaff == nil {
			// The edge medicalrecordstaff was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: medicalrecordstaff.Label}
		}
		return e.Medicalrecordstaff, nil
	}
	return nil, &NotLoadedError{edge: "medicalrecordstaff"}
}

// User2registrarOrErr returns the User2registrar value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) User2registrarOrErr() (*Registrar, error) {
	if e.loadedTypes[4] {
		if e.User2registrar == nil {
			// The edge user2registrar was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: registrar.Label}
		}
		return e.User2registrar, nil
	}
	return nil, &NotLoadedError{edge: "user2registrar"}
}

// UserstatusOrErr returns the Userstatus value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) UserstatusOrErr() (*Userstatus, error) {
	if e.loadedTypes[5] {
		if e.Userstatus == nil {
			// The edge userstatus was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: userstatus.Label}
		}
		return e.Userstatus, nil
	}
	return nil, &NotLoadedError{edge: "userstatus"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // email
		&sql.NullString{}, // password
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*User) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // user_id
		&sql.NullInt64{}, // userstatus_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[0])
	} else if value.Valid {
		u.Email = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field password", values[1])
	} else if value.Valid {
		u.Password = value.String
	}
	values = values[2:]
	if len(values) == len(user.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_id", value)
		} else if value.Valid {
			u.user_id = new(int)
			*u.user_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field userstatus_id", value)
		} else if value.Valid {
			u.userstatus_id = new(int)
			*u.userstatus_id = int(value.Int64)
		}
	}
	return nil
}

// QueryFinancier queries the financier edge of the User.
func (u *User) QueryFinancier() *FinancierQuery {
	return (&UserClient{config: u.config}).QueryFinancier(u)
}

// QueryHistorytaking queries the historytaking edge of the User.
func (u *User) QueryHistorytaking() *NurseQuery {
	return (&UserClient{config: u.config}).QueryHistorytaking(u)
}

// QueryUserPatientrights queries the UserPatientrights edge of the User.
func (u *User) QueryUserPatientrights() *PatientrightsQuery {
	return (&UserClient{config: u.config}).QueryUserPatientrights(u)
}

// QueryMedicalrecordstaff queries the medicalrecordstaff edge of the User.
func (u *User) QueryMedicalrecordstaff() *MedicalrecordstaffQuery {
	return (&UserClient{config: u.config}).QueryMedicalrecordstaff(u)
}

// QueryUser2registrar queries the user2registrar edge of the User.
func (u *User) QueryUser2registrar() *RegistrarQuery {
	return (&UserClient{config: u.config}).QueryUser2registrar(u)
}

// QueryUserstatus queries the userstatus edge of the User.
func (u *User) QueryUserstatus() *UserstatusQuery {
	return (&UserClient{config: u.config}).QueryUserstatus(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteString(", password=")
	builder.WriteString(u.Password)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
