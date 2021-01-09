// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/bill"
	"github.com/team10/app/ent/paytype"
	"github.com/team10/app/ent/predicate"
)

// PaytypeUpdate is the builder for updating Paytype entities.
type PaytypeUpdate struct {
	config
	hooks      []Hook
	mutation   *PaytypeMutation
	predicates []predicate.Paytype
}

// Where adds a new predicate for the builder.
func (pu *PaytypeUpdate) Where(ps ...predicate.Paytype) *PaytypeUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetPaytype sets the paytype field.
func (pu *PaytypeUpdate) SetPaytype(s string) *PaytypeUpdate {
	pu.mutation.SetPaytype(s)
	return pu
}

// AddBillIDs adds the bills edge to Bill by ids.
func (pu *PaytypeUpdate) AddBillIDs(ids ...int) *PaytypeUpdate {
	pu.mutation.AddBillIDs(ids...)
	return pu
}

// AddBills adds the bills edges to Bill.
func (pu *PaytypeUpdate) AddBills(b ...*Bill) *PaytypeUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pu.AddBillIDs(ids...)
}

// Mutation returns the PaytypeMutation object of the builder.
func (pu *PaytypeUpdate) Mutation() *PaytypeMutation {
	return pu.mutation
}

// RemoveBillIDs removes the bills edge to Bill by ids.
func (pu *PaytypeUpdate) RemoveBillIDs(ids ...int) *PaytypeUpdate {
	pu.mutation.RemoveBillIDs(ids...)
	return pu
}

// RemoveBills removes bills edges to Bill.
func (pu *PaytypeUpdate) RemoveBills(b ...*Bill) *PaytypeUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pu.RemoveBillIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PaytypeUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := pu.mutation.Paytype(); ok {
		if err := paytype.PaytypeValidator(v); err != nil {
			return 0, &ValidationError{Name: "paytype", err: fmt.Errorf("ent: validator failed for field \"paytype\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaytypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PaytypeUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PaytypeUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PaytypeUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PaytypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   paytype.Table,
			Columns: paytype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: paytype.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Paytype(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: paytype.FieldPaytype,
		})
	}
	if nodes := pu.mutation.RemovedBillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   paytype.BillsTable,
			Columns: []string{paytype.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.BillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   paytype.BillsTable,
			Columns: []string{paytype.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{paytype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PaytypeUpdateOne is the builder for updating a single Paytype entity.
type PaytypeUpdateOne struct {
	config
	hooks    []Hook
	mutation *PaytypeMutation
}

// SetPaytype sets the paytype field.
func (puo *PaytypeUpdateOne) SetPaytype(s string) *PaytypeUpdateOne {
	puo.mutation.SetPaytype(s)
	return puo
}

// AddBillIDs adds the bills edge to Bill by ids.
func (puo *PaytypeUpdateOne) AddBillIDs(ids ...int) *PaytypeUpdateOne {
	puo.mutation.AddBillIDs(ids...)
	return puo
}

// AddBills adds the bills edges to Bill.
func (puo *PaytypeUpdateOne) AddBills(b ...*Bill) *PaytypeUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return puo.AddBillIDs(ids...)
}

// Mutation returns the PaytypeMutation object of the builder.
func (puo *PaytypeUpdateOne) Mutation() *PaytypeMutation {
	return puo.mutation
}

// RemoveBillIDs removes the bills edge to Bill by ids.
func (puo *PaytypeUpdateOne) RemoveBillIDs(ids ...int) *PaytypeUpdateOne {
	puo.mutation.RemoveBillIDs(ids...)
	return puo
}

// RemoveBills removes bills edges to Bill.
func (puo *PaytypeUpdateOne) RemoveBills(b ...*Bill) *PaytypeUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return puo.RemoveBillIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *PaytypeUpdateOne) Save(ctx context.Context) (*Paytype, error) {
	if v, ok := puo.mutation.Paytype(); ok {
		if err := paytype.PaytypeValidator(v); err != nil {
			return nil, &ValidationError{Name: "paytype", err: fmt.Errorf("ent: validator failed for field \"paytype\": %w", err)}
		}
	}

	var (
		err  error
		node *Paytype
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaytypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PaytypeUpdateOne) SaveX(ctx context.Context) *Paytype {
	pa, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pa
}

// Exec executes the query on the entity.
func (puo *PaytypeUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PaytypeUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PaytypeUpdateOne) sqlSave(ctx context.Context) (pa *Paytype, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   paytype.Table,
			Columns: paytype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: paytype.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Paytype.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Paytype(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: paytype.FieldPaytype,
		})
	}
	if nodes := puo.mutation.RemovedBillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   paytype.BillsTable,
			Columns: []string{paytype.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.BillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   paytype.BillsTable,
			Columns: []string{paytype.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pa = &Paytype{config: puo.config}
	_spec.Assign = pa.assignValues
	_spec.ScanValues = pa.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{paytype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pa, nil
}