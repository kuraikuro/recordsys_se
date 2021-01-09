// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/abilitypatientrights"
	"github.com/team10/app/ent/patientrightstype"
	"github.com/team10/app/ent/predicate"
)

// AbilitypatientrightsUpdate is the builder for updating Abilitypatientrights entities.
type AbilitypatientrightsUpdate struct {
	config
	hooks      []Hook
	mutation   *AbilitypatientrightsMutation
	predicates []predicate.Abilitypatientrights
}

// Where adds a new predicate for the builder.
func (au *AbilitypatientrightsUpdate) Where(ps ...predicate.Abilitypatientrights) *AbilitypatientrightsUpdate {
	au.predicates = append(au.predicates, ps...)
	return au
}

// SetOperative sets the Operative field.
func (au *AbilitypatientrightsUpdate) SetOperative(i int) *AbilitypatientrightsUpdate {
	au.mutation.ResetOperative()
	au.mutation.SetOperative(i)
	return au
}

// AddOperative adds i to Operative.
func (au *AbilitypatientrightsUpdate) AddOperative(i int) *AbilitypatientrightsUpdate {
	au.mutation.AddOperative(i)
	return au
}

// SetMedicalSupplies sets the MedicalSupplies field.
func (au *AbilitypatientrightsUpdate) SetMedicalSupplies(i int) *AbilitypatientrightsUpdate {
	au.mutation.ResetMedicalSupplies()
	au.mutation.SetMedicalSupplies(i)
	return au
}

// AddMedicalSupplies adds i to MedicalSupplies.
func (au *AbilitypatientrightsUpdate) AddMedicalSupplies(i int) *AbilitypatientrightsUpdate {
	au.mutation.AddMedicalSupplies(i)
	return au
}

// SetExamine sets the Examine field.
func (au *AbilitypatientrightsUpdate) SetExamine(i int) *AbilitypatientrightsUpdate {
	au.mutation.ResetExamine()
	au.mutation.SetExamine(i)
	return au
}

// AddExamine adds i to Examine.
func (au *AbilitypatientrightsUpdate) AddExamine(i int) *AbilitypatientrightsUpdate {
	au.mutation.AddExamine(i)
	return au
}

// AddAbilitypatientrightsPatientrightstypeIDs adds the AbilitypatientrightsPatientrightstype edge to Patientrightstype by ids.
func (au *AbilitypatientrightsUpdate) AddAbilitypatientrightsPatientrightstypeIDs(ids ...int) *AbilitypatientrightsUpdate {
	au.mutation.AddAbilitypatientrightsPatientrightstypeIDs(ids...)
	return au
}

// AddAbilitypatientrightsPatientrightstype adds the AbilitypatientrightsPatientrightstype edges to Patientrightstype.
func (au *AbilitypatientrightsUpdate) AddAbilitypatientrightsPatientrightstype(p ...*Patientrightstype) *AbilitypatientrightsUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return au.AddAbilitypatientrightsPatientrightstypeIDs(ids...)
}

// Mutation returns the AbilitypatientrightsMutation object of the builder.
func (au *AbilitypatientrightsUpdate) Mutation() *AbilitypatientrightsMutation {
	return au.mutation
}

// RemoveAbilitypatientrightsPatientrightstypeIDs removes the AbilitypatientrightsPatientrightstype edge to Patientrightstype by ids.
func (au *AbilitypatientrightsUpdate) RemoveAbilitypatientrightsPatientrightstypeIDs(ids ...int) *AbilitypatientrightsUpdate {
	au.mutation.RemoveAbilitypatientrightsPatientrightstypeIDs(ids...)
	return au
}

// RemoveAbilitypatientrightsPatientrightstype removes AbilitypatientrightsPatientrightstype edges to Patientrightstype.
func (au *AbilitypatientrightsUpdate) RemoveAbilitypatientrightsPatientrightstype(p ...*Patientrightstype) *AbilitypatientrightsUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return au.RemoveAbilitypatientrightsPatientrightstypeIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (au *AbilitypatientrightsUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AbilitypatientrightsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AbilitypatientrightsUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AbilitypatientrightsUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AbilitypatientrightsUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AbilitypatientrightsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   abilitypatientrights.Table,
			Columns: abilitypatientrights.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: abilitypatientrights.FieldID,
			},
		},
	}
	if ps := au.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Operative(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: abilitypatientrights.FieldOperative,
		})
	}
	if value, ok := au.mutation.AddedOperative(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: abilitypatientrights.FieldOperative,
		})
	}
	if value, ok := au.mutation.MedicalSupplies(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: abilitypatientrights.FieldMedicalSupplies,
		})
	}
	if value, ok := au.mutation.AddedMedicalSupplies(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: abilitypatientrights.FieldMedicalSupplies,
		})
	}
	if value, ok := au.mutation.Examine(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: abilitypatientrights.FieldExamine,
		})
	}
	if value, ok := au.mutation.AddedExamine(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: abilitypatientrights.FieldExamine,
		})
	}
	if nodes := au.mutation.RemovedAbilitypatientrightsPatientrightstypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   abilitypatientrights.AbilitypatientrightsPatientrightstypeTable,
			Columns: []string{abilitypatientrights.AbilitypatientrightsPatientrightstypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrightstype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AbilitypatientrightsPatientrightstypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   abilitypatientrights.AbilitypatientrightsPatientrightstypeTable,
			Columns: []string{abilitypatientrights.AbilitypatientrightsPatientrightstypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrightstype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{abilitypatientrights.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AbilitypatientrightsUpdateOne is the builder for updating a single Abilitypatientrights entity.
type AbilitypatientrightsUpdateOne struct {
	config
	hooks    []Hook
	mutation *AbilitypatientrightsMutation
}

// SetOperative sets the Operative field.
func (auo *AbilitypatientrightsUpdateOne) SetOperative(i int) *AbilitypatientrightsUpdateOne {
	auo.mutation.ResetOperative()
	auo.mutation.SetOperative(i)
	return auo
}

// AddOperative adds i to Operative.
func (auo *AbilitypatientrightsUpdateOne) AddOperative(i int) *AbilitypatientrightsUpdateOne {
	auo.mutation.AddOperative(i)
	return auo
}

// SetMedicalSupplies sets the MedicalSupplies field.
func (auo *AbilitypatientrightsUpdateOne) SetMedicalSupplies(i int) *AbilitypatientrightsUpdateOne {
	auo.mutation.ResetMedicalSupplies()
	auo.mutation.SetMedicalSupplies(i)
	return auo
}

// AddMedicalSupplies adds i to MedicalSupplies.
func (auo *AbilitypatientrightsUpdateOne) AddMedicalSupplies(i int) *AbilitypatientrightsUpdateOne {
	auo.mutation.AddMedicalSupplies(i)
	return auo
}

// SetExamine sets the Examine field.
func (auo *AbilitypatientrightsUpdateOne) SetExamine(i int) *AbilitypatientrightsUpdateOne {
	auo.mutation.ResetExamine()
	auo.mutation.SetExamine(i)
	return auo
}

// AddExamine adds i to Examine.
func (auo *AbilitypatientrightsUpdateOne) AddExamine(i int) *AbilitypatientrightsUpdateOne {
	auo.mutation.AddExamine(i)
	return auo
}

// AddAbilitypatientrightsPatientrightstypeIDs adds the AbilitypatientrightsPatientrightstype edge to Patientrightstype by ids.
func (auo *AbilitypatientrightsUpdateOne) AddAbilitypatientrightsPatientrightstypeIDs(ids ...int) *AbilitypatientrightsUpdateOne {
	auo.mutation.AddAbilitypatientrightsPatientrightstypeIDs(ids...)
	return auo
}

// AddAbilitypatientrightsPatientrightstype adds the AbilitypatientrightsPatientrightstype edges to Patientrightstype.
func (auo *AbilitypatientrightsUpdateOne) AddAbilitypatientrightsPatientrightstype(p ...*Patientrightstype) *AbilitypatientrightsUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return auo.AddAbilitypatientrightsPatientrightstypeIDs(ids...)
}

// Mutation returns the AbilitypatientrightsMutation object of the builder.
func (auo *AbilitypatientrightsUpdateOne) Mutation() *AbilitypatientrightsMutation {
	return auo.mutation
}

// RemoveAbilitypatientrightsPatientrightstypeIDs removes the AbilitypatientrightsPatientrightstype edge to Patientrightstype by ids.
func (auo *AbilitypatientrightsUpdateOne) RemoveAbilitypatientrightsPatientrightstypeIDs(ids ...int) *AbilitypatientrightsUpdateOne {
	auo.mutation.RemoveAbilitypatientrightsPatientrightstypeIDs(ids...)
	return auo
}

// RemoveAbilitypatientrightsPatientrightstype removes AbilitypatientrightsPatientrightstype edges to Patientrightstype.
func (auo *AbilitypatientrightsUpdateOne) RemoveAbilitypatientrightsPatientrightstype(p ...*Patientrightstype) *AbilitypatientrightsUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return auo.RemoveAbilitypatientrightsPatientrightstypeIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (auo *AbilitypatientrightsUpdateOne) Save(ctx context.Context) (*Abilitypatientrights, error) {

	var (
		err  error
		node *Abilitypatientrights
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AbilitypatientrightsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AbilitypatientrightsUpdateOne) SaveX(ctx context.Context) *Abilitypatientrights {
	a, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return a
}

// Exec executes the query on the entity.
func (auo *AbilitypatientrightsUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AbilitypatientrightsUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AbilitypatientrightsUpdateOne) sqlSave(ctx context.Context) (a *Abilitypatientrights, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   abilitypatientrights.Table,
			Columns: abilitypatientrights.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: abilitypatientrights.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Abilitypatientrights.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := auo.mutation.Operative(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: abilitypatientrights.FieldOperative,
		})
	}
	if value, ok := auo.mutation.AddedOperative(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: abilitypatientrights.FieldOperative,
		})
	}
	if value, ok := auo.mutation.MedicalSupplies(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: abilitypatientrights.FieldMedicalSupplies,
		})
	}
	if value, ok := auo.mutation.AddedMedicalSupplies(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: abilitypatientrights.FieldMedicalSupplies,
		})
	}
	if value, ok := auo.mutation.Examine(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: abilitypatientrights.FieldExamine,
		})
	}
	if value, ok := auo.mutation.AddedExamine(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: abilitypatientrights.FieldExamine,
		})
	}
	if nodes := auo.mutation.RemovedAbilitypatientrightsPatientrightstypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   abilitypatientrights.AbilitypatientrightsPatientrightstypeTable,
			Columns: []string{abilitypatientrights.AbilitypatientrightsPatientrightstypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrightstype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AbilitypatientrightsPatientrightstypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   abilitypatientrights.AbilitypatientrightsPatientrightstypeTable,
			Columns: []string{abilitypatientrights.AbilitypatientrightsPatientrightstypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrightstype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	a = &Abilitypatientrights{config: auo.config}
	_spec.Assign = a.assignValues
	_spec.ScanValues = a.scanValues()
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{abilitypatientrights.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return a, nil
}