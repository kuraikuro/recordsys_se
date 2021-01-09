// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/doctorinfo"
	"github.com/team10/app/ent/officeroom"
)

// OfficeroomCreate is the builder for creating a Officeroom entity.
type OfficeroomCreate struct {
	config
	mutation *OfficeroomMutation
	hooks    []Hook
}

// SetRoomnumber sets the roomnumber field.
func (oc *OfficeroomCreate) SetRoomnumber(s string) *OfficeroomCreate {
	oc.mutation.SetRoomnumber(s)
	return oc
}

// AddOfficeroom2doctorinfoIDs adds the officeroom2doctorinfo edge to Doctorinfo by ids.
func (oc *OfficeroomCreate) AddOfficeroom2doctorinfoIDs(ids ...int) *OfficeroomCreate {
	oc.mutation.AddOfficeroom2doctorinfoIDs(ids...)
	return oc
}

// AddOfficeroom2doctorinfo adds the officeroom2doctorinfo edges to Doctorinfo.
func (oc *OfficeroomCreate) AddOfficeroom2doctorinfo(d ...*Doctorinfo) *OfficeroomCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return oc.AddOfficeroom2doctorinfoIDs(ids...)
}

// Mutation returns the OfficeroomMutation object of the builder.
func (oc *OfficeroomCreate) Mutation() *OfficeroomMutation {
	return oc.mutation
}

// Save creates the Officeroom in the database.
func (oc *OfficeroomCreate) Save(ctx context.Context) (*Officeroom, error) {
	if _, ok := oc.mutation.Roomnumber(); !ok {
		return nil, &ValidationError{Name: "roomnumber", err: errors.New("ent: missing required field \"roomnumber\"")}
	}
	if v, ok := oc.mutation.Roomnumber(); ok {
		if err := officeroom.RoomnumberValidator(v); err != nil {
			return nil, &ValidationError{Name: "roomnumber", err: fmt.Errorf("ent: validator failed for field \"roomnumber\": %w", err)}
		}
	}
	var (
		err  error
		node *Officeroom
	)
	if len(oc.hooks) == 0 {
		node, err = oc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OfficeroomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oc.mutation = mutation
			node, err = oc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oc.hooks) - 1; i >= 0; i-- {
			mut = oc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OfficeroomCreate) SaveX(ctx context.Context) *Officeroom {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (oc *OfficeroomCreate) sqlSave(ctx context.Context) (*Officeroom, error) {
	o, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	o.ID = int(id)
	return o, nil
}

func (oc *OfficeroomCreate) createSpec() (*Officeroom, *sqlgraph.CreateSpec) {
	var (
		o     = &Officeroom{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: officeroom.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: officeroom.FieldID,
			},
		}
	)
	if value, ok := oc.mutation.Roomnumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: officeroom.FieldRoomnumber,
		})
		o.Roomnumber = value
	}
	if nodes := oc.mutation.Officeroom2doctorinfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   officeroom.Officeroom2doctorinfoTable,
			Columns: []string{officeroom.Officeroom2doctorinfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctorinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return o, _spec
}