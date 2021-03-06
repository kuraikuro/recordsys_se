// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/financier"
	"github.com/team10/app/ent/medicalrecordstaff"
	"github.com/team10/app/ent/nurse"
	"github.com/team10/app/ent/patientrights"
	"github.com/team10/app/ent/predicate"
	"github.com/team10/app/ent/registrar"
	"github.com/team10/app/ent/user"
	"github.com/team10/app/ent/userstatus"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks      []Hook
	mutation   *UserMutation
	predicates []predicate.User
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetEmail sets the email field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetPassword sets the password field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetFinancierID sets the financier edge to Financier by id.
func (uu *UserUpdate) SetFinancierID(id int) *UserUpdate {
	uu.mutation.SetFinancierID(id)
	return uu
}

// SetNillableFinancierID sets the financier edge to Financier by id if the given value is not nil.
func (uu *UserUpdate) SetNillableFinancierID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetFinancierID(*id)
	}
	return uu
}

// SetFinancier sets the financier edge to Financier.
func (uu *UserUpdate) SetFinancier(f *Financier) *UserUpdate {
	return uu.SetFinancierID(f.ID)
}

// SetHistorytakingID sets the historytaking edge to Nurse by id.
func (uu *UserUpdate) SetHistorytakingID(id int) *UserUpdate {
	uu.mutation.SetHistorytakingID(id)
	return uu
}

// SetNillableHistorytakingID sets the historytaking edge to Nurse by id if the given value is not nil.
func (uu *UserUpdate) SetNillableHistorytakingID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetHistorytakingID(*id)
	}
	return uu
}

// SetHistorytaking sets the historytaking edge to Nurse.
func (uu *UserUpdate) SetHistorytaking(n *Nurse) *UserUpdate {
	return uu.SetHistorytakingID(n.ID)
}

// SetUserPatientrightsID sets the UserPatientrights edge to Patientrights by id.
func (uu *UserUpdate) SetUserPatientrightsID(id int) *UserUpdate {
	uu.mutation.SetUserPatientrightsID(id)
	return uu
}

// SetNillableUserPatientrightsID sets the UserPatientrights edge to Patientrights by id if the given value is not nil.
func (uu *UserUpdate) SetNillableUserPatientrightsID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetUserPatientrightsID(*id)
	}
	return uu
}

// SetUserPatientrights sets the UserPatientrights edge to Patientrights.
func (uu *UserUpdate) SetUserPatientrights(p *Patientrights) *UserUpdate {
	return uu.SetUserPatientrightsID(p.ID)
}

// SetMedicalrecordstaffID sets the medicalrecordstaff edge to Medicalrecordstaff by id.
func (uu *UserUpdate) SetMedicalrecordstaffID(id int) *UserUpdate {
	uu.mutation.SetMedicalrecordstaffID(id)
	return uu
}

// SetNillableMedicalrecordstaffID sets the medicalrecordstaff edge to Medicalrecordstaff by id if the given value is not nil.
func (uu *UserUpdate) SetNillableMedicalrecordstaffID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetMedicalrecordstaffID(*id)
	}
	return uu
}

// SetMedicalrecordstaff sets the medicalrecordstaff edge to Medicalrecordstaff.
func (uu *UserUpdate) SetMedicalrecordstaff(m *Medicalrecordstaff) *UserUpdate {
	return uu.SetMedicalrecordstaffID(m.ID)
}

// SetUser2registrarID sets the user2registrar edge to Registrar by id.
func (uu *UserUpdate) SetUser2registrarID(id int) *UserUpdate {
	uu.mutation.SetUser2registrarID(id)
	return uu
}

// SetNillableUser2registrarID sets the user2registrar edge to Registrar by id if the given value is not nil.
func (uu *UserUpdate) SetNillableUser2registrarID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetUser2registrarID(*id)
	}
	return uu
}

// SetUser2registrar sets the user2registrar edge to Registrar.
func (uu *UserUpdate) SetUser2registrar(r *Registrar) *UserUpdate {
	return uu.SetUser2registrarID(r.ID)
}

// SetUserstatusID sets the userstatus edge to Userstatus by id.
func (uu *UserUpdate) SetUserstatusID(id int) *UserUpdate {
	uu.mutation.SetUserstatusID(id)
	return uu
}

// SetNillableUserstatusID sets the userstatus edge to Userstatus by id if the given value is not nil.
func (uu *UserUpdate) SetNillableUserstatusID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetUserstatusID(*id)
	}
	return uu
}

// SetUserstatus sets the userstatus edge to Userstatus.
func (uu *UserUpdate) SetUserstatus(u *Userstatus) *UserUpdate {
	return uu.SetUserstatusID(u.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearFinancier clears the financier edge to Financier.
func (uu *UserUpdate) ClearFinancier() *UserUpdate {
	uu.mutation.ClearFinancier()
	return uu
}

// ClearHistorytaking clears the historytaking edge to Nurse.
func (uu *UserUpdate) ClearHistorytaking() *UserUpdate {
	uu.mutation.ClearHistorytaking()
	return uu
}

// ClearUserPatientrights clears the UserPatientrights edge to Patientrights.
func (uu *UserUpdate) ClearUserPatientrights() *UserUpdate {
	uu.mutation.ClearUserPatientrights()
	return uu
}

// ClearMedicalrecordstaff clears the medicalrecordstaff edge to Medicalrecordstaff.
func (uu *UserUpdate) ClearMedicalrecordstaff() *UserUpdate {
	uu.mutation.ClearMedicalrecordstaff()
	return uu
}

// ClearUser2registrar clears the user2registrar edge to Registrar.
func (uu *UserUpdate) ClearUser2registrar() *UserUpdate {
	uu.mutation.ClearUser2registrar()
	return uu
}

// ClearUserstatus clears the userstatus edge to Userstatus.
func (uu *UserUpdate) ClearUserstatus() *UserUpdate {
	uu.mutation.ClearUserstatus()
	return uu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return 0, &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return 0, &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
	}
	if uu.mutation.FinancierCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.FinancierTable,
			Columns: []string{user.FinancierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: financier.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.FinancierIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.FinancierTable,
			Columns: []string{user.FinancierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: financier.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.HistorytakingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.HistorytakingTable,
			Columns: []string{user.HistorytakingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.HistorytakingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.HistorytakingTable,
			Columns: []string{user.HistorytakingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.UserPatientrightsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.UserPatientrightsTable,
			Columns: []string{user.UserPatientrightsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrights.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UserPatientrightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.UserPatientrightsTable,
			Columns: []string{user.UserPatientrightsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrights.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.MedicalrecordstaffCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.MedicalrecordstaffTable,
			Columns: []string{user.MedicalrecordstaffColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalrecordstaff.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.MedicalrecordstaffIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.MedicalrecordstaffTable,
			Columns: []string{user.MedicalrecordstaffColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalrecordstaff.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.User2registrarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.User2registrarTable,
			Columns: []string{user.User2registrarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: registrar.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.User2registrarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.User2registrarTable,
			Columns: []string{user.User2registrarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: registrar.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.UserstatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.UserstatusTable,
			Columns: []string{user.UserstatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userstatus.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UserstatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.UserstatusTable,
			Columns: []string{user.UserstatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userstatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// SetEmail sets the email field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetPassword sets the password field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetFinancierID sets the financier edge to Financier by id.
func (uuo *UserUpdateOne) SetFinancierID(id int) *UserUpdateOne {
	uuo.mutation.SetFinancierID(id)
	return uuo
}

// SetNillableFinancierID sets the financier edge to Financier by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableFinancierID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetFinancierID(*id)
	}
	return uuo
}

// SetFinancier sets the financier edge to Financier.
func (uuo *UserUpdateOne) SetFinancier(f *Financier) *UserUpdateOne {
	return uuo.SetFinancierID(f.ID)
}

// SetHistorytakingID sets the historytaking edge to Nurse by id.
func (uuo *UserUpdateOne) SetHistorytakingID(id int) *UserUpdateOne {
	uuo.mutation.SetHistorytakingID(id)
	return uuo
}

// SetNillableHistorytakingID sets the historytaking edge to Nurse by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableHistorytakingID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetHistorytakingID(*id)
	}
	return uuo
}

// SetHistorytaking sets the historytaking edge to Nurse.
func (uuo *UserUpdateOne) SetHistorytaking(n *Nurse) *UserUpdateOne {
	return uuo.SetHistorytakingID(n.ID)
}

// SetUserPatientrightsID sets the UserPatientrights edge to Patientrights by id.
func (uuo *UserUpdateOne) SetUserPatientrightsID(id int) *UserUpdateOne {
	uuo.mutation.SetUserPatientrightsID(id)
	return uuo
}

// SetNillableUserPatientrightsID sets the UserPatientrights edge to Patientrights by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUserPatientrightsID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetUserPatientrightsID(*id)
	}
	return uuo
}

// SetUserPatientrights sets the UserPatientrights edge to Patientrights.
func (uuo *UserUpdateOne) SetUserPatientrights(p *Patientrights) *UserUpdateOne {
	return uuo.SetUserPatientrightsID(p.ID)
}

// SetMedicalrecordstaffID sets the medicalrecordstaff edge to Medicalrecordstaff by id.
func (uuo *UserUpdateOne) SetMedicalrecordstaffID(id int) *UserUpdateOne {
	uuo.mutation.SetMedicalrecordstaffID(id)
	return uuo
}

// SetNillableMedicalrecordstaffID sets the medicalrecordstaff edge to Medicalrecordstaff by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableMedicalrecordstaffID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetMedicalrecordstaffID(*id)
	}
	return uuo
}

// SetMedicalrecordstaff sets the medicalrecordstaff edge to Medicalrecordstaff.
func (uuo *UserUpdateOne) SetMedicalrecordstaff(m *Medicalrecordstaff) *UserUpdateOne {
	return uuo.SetMedicalrecordstaffID(m.ID)
}

// SetUser2registrarID sets the user2registrar edge to Registrar by id.
func (uuo *UserUpdateOne) SetUser2registrarID(id int) *UserUpdateOne {
	uuo.mutation.SetUser2registrarID(id)
	return uuo
}

// SetNillableUser2registrarID sets the user2registrar edge to Registrar by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUser2registrarID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetUser2registrarID(*id)
	}
	return uuo
}

// SetUser2registrar sets the user2registrar edge to Registrar.
func (uuo *UserUpdateOne) SetUser2registrar(r *Registrar) *UserUpdateOne {
	return uuo.SetUser2registrarID(r.ID)
}

// SetUserstatusID sets the userstatus edge to Userstatus by id.
func (uuo *UserUpdateOne) SetUserstatusID(id int) *UserUpdateOne {
	uuo.mutation.SetUserstatusID(id)
	return uuo
}

// SetNillableUserstatusID sets the userstatus edge to Userstatus by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUserstatusID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetUserstatusID(*id)
	}
	return uuo
}

// SetUserstatus sets the userstatus edge to Userstatus.
func (uuo *UserUpdateOne) SetUserstatus(u *Userstatus) *UserUpdateOne {
	return uuo.SetUserstatusID(u.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearFinancier clears the financier edge to Financier.
func (uuo *UserUpdateOne) ClearFinancier() *UserUpdateOne {
	uuo.mutation.ClearFinancier()
	return uuo
}

// ClearHistorytaking clears the historytaking edge to Nurse.
func (uuo *UserUpdateOne) ClearHistorytaking() *UserUpdateOne {
	uuo.mutation.ClearHistorytaking()
	return uuo
}

// ClearUserPatientrights clears the UserPatientrights edge to Patientrights.
func (uuo *UserUpdateOne) ClearUserPatientrights() *UserUpdateOne {
	uuo.mutation.ClearUserPatientrights()
	return uuo
}

// ClearMedicalrecordstaff clears the medicalrecordstaff edge to Medicalrecordstaff.
func (uuo *UserUpdateOne) ClearMedicalrecordstaff() *UserUpdateOne {
	uuo.mutation.ClearMedicalrecordstaff()
	return uuo
}

// ClearUser2registrar clears the user2registrar edge to Registrar.
func (uuo *UserUpdateOne) ClearUser2registrar() *UserUpdateOne {
	uuo.mutation.ClearUser2registrar()
	return uuo
}

// ClearUserstatus clears the userstatus edge to Userstatus.
func (uuo *UserUpdateOne) ClearUserstatus() *UserUpdateOne {
	uuo.mutation.ClearUserstatus()
	return uuo
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return nil, &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return nil, &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}

	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	u, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (u *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing User.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := uuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
	}
	if uuo.mutation.FinancierCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.FinancierTable,
			Columns: []string{user.FinancierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: financier.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.FinancierIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.FinancierTable,
			Columns: []string{user.FinancierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: financier.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.HistorytakingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.HistorytakingTable,
			Columns: []string{user.HistorytakingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.HistorytakingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.HistorytakingTable,
			Columns: []string{user.HistorytakingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.UserPatientrightsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.UserPatientrightsTable,
			Columns: []string{user.UserPatientrightsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrights.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UserPatientrightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.UserPatientrightsTable,
			Columns: []string{user.UserPatientrightsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrights.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.MedicalrecordstaffCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.MedicalrecordstaffTable,
			Columns: []string{user.MedicalrecordstaffColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalrecordstaff.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.MedicalrecordstaffIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.MedicalrecordstaffTable,
			Columns: []string{user.MedicalrecordstaffColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalrecordstaff.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.User2registrarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.User2registrarTable,
			Columns: []string{user.User2registrarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: registrar.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.User2registrarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.User2registrarTable,
			Columns: []string{user.User2registrarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: registrar.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.UserstatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.UserstatusTable,
			Columns: []string{user.UserstatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userstatus.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UserstatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.UserstatusTable,
			Columns: []string{user.UserstatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userstatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	u = &User{config: uuo.config}
	_spec.Assign = u.assignValues
	_spec.ScanValues = u.scanValues()
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return u, nil
}
