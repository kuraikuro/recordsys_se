// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/abilitypatientrights"
	"github.com/team10/app/ent/patientrights"
	"github.com/team10/app/ent/patientrightstype"
	"github.com/team10/app/ent/predicate"
)

// PatientrightstypeQuery is the builder for querying Patientrightstype entities.
type PatientrightstypeQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Patientrightstype
	// eager-loading edges.
	withPatientrightstypePatientrights        *PatientrightsQuery
	withPatientrightstypeAbilitypatientrights *AbilitypatientrightsQuery
	withFKs                                   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (pq *PatientrightstypeQuery) Where(ps ...predicate.Patientrightstype) *PatientrightstypeQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *PatientrightstypeQuery) Limit(limit int) *PatientrightstypeQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *PatientrightstypeQuery) Offset(offset int) *PatientrightstypeQuery {
	pq.offset = &offset
	return pq
}

// Order adds an order step to the query.
func (pq *PatientrightstypeQuery) Order(o ...OrderFunc) *PatientrightstypeQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryPatientrightstypePatientrights chains the current query on the PatientrightstypePatientrights edge.
func (pq *PatientrightstypeQuery) QueryPatientrightstypePatientrights() *PatientrightsQuery {
	query := &PatientrightsQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(patientrightstype.Table, patientrightstype.FieldID, pq.sqlQuery()),
			sqlgraph.To(patientrights.Table, patientrights.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, patientrightstype.PatientrightstypePatientrightsTable, patientrightstype.PatientrightstypePatientrightsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPatientrightstypeAbilitypatientrights chains the current query on the PatientrightstypeAbilitypatientrights edge.
func (pq *PatientrightstypeQuery) QueryPatientrightstypeAbilitypatientrights() *AbilitypatientrightsQuery {
	query := &AbilitypatientrightsQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(patientrightstype.Table, patientrightstype.FieldID, pq.sqlQuery()),
			sqlgraph.To(abilitypatientrights.Table, abilitypatientrights.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, patientrightstype.PatientrightstypeAbilitypatientrightsTable, patientrightstype.PatientrightstypeAbilitypatientrightsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Patientrightstype entity in the query. Returns *NotFoundError when no patientrightstype was found.
func (pq *PatientrightstypeQuery) First(ctx context.Context) (*Patientrightstype, error) {
	pas, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(pas) == 0 {
		return nil, &NotFoundError{patientrightstype.Label}
	}
	return pas[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PatientrightstypeQuery) FirstX(ctx context.Context) *Patientrightstype {
	pa, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return pa
}

// FirstID returns the first Patientrightstype id in the query. Returns *NotFoundError when no id was found.
func (pq *PatientrightstypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{patientrightstype.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (pq *PatientrightstypeQuery) FirstXID(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Patientrightstype entity in the query, returns an error if not exactly one entity was returned.
func (pq *PatientrightstypeQuery) Only(ctx context.Context) (*Patientrightstype, error) {
	pas, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(pas) {
	case 1:
		return pas[0], nil
	case 0:
		return nil, &NotFoundError{patientrightstype.Label}
	default:
		return nil, &NotSingularError{patientrightstype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PatientrightstypeQuery) OnlyX(ctx context.Context) *Patientrightstype {
	pa, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return pa
}

// OnlyID returns the only Patientrightstype id in the query, returns an error if not exactly one id was returned.
func (pq *PatientrightstypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{patientrightstype.Label}
	default:
		err = &NotSingularError{patientrightstype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PatientrightstypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Patientrightstypes.
func (pq *PatientrightstypeQuery) All(ctx context.Context) ([]*Patientrightstype, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *PatientrightstypeQuery) AllX(ctx context.Context) []*Patientrightstype {
	pas, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return pas
}

// IDs executes the query and returns a list of Patientrightstype ids.
func (pq *PatientrightstypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pq.Select(patientrightstype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PatientrightstypeQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PatientrightstypeQuery) Count(ctx context.Context) (int, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PatientrightstypeQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PatientrightstypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PatientrightstypeQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PatientrightstypeQuery) Clone() *PatientrightstypeQuery {
	return &PatientrightstypeQuery{
		config:     pq.config,
		limit:      pq.limit,
		offset:     pq.offset,
		order:      append([]OrderFunc{}, pq.order...),
		unique:     append([]string{}, pq.unique...),
		predicates: append([]predicate.Patientrightstype{}, pq.predicates...),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

//  WithPatientrightstypePatientrights tells the query-builder to eager-loads the nodes that are connected to
// the "PatientrightstypePatientrights" edge. The optional arguments used to configure the query builder of the edge.
func (pq *PatientrightstypeQuery) WithPatientrightstypePatientrights(opts ...func(*PatientrightsQuery)) *PatientrightstypeQuery {
	query := &PatientrightsQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withPatientrightstypePatientrights = query
	return pq
}

//  WithPatientrightstypeAbilitypatientrights tells the query-builder to eager-loads the nodes that are connected to
// the "PatientrightstypeAbilitypatientrights" edge. The optional arguments used to configure the query builder of the edge.
func (pq *PatientrightstypeQuery) WithPatientrightstypeAbilitypatientrights(opts ...func(*AbilitypatientrightsQuery)) *PatientrightstypeQuery {
	query := &AbilitypatientrightsQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withPatientrightstypeAbilitypatientrights = query
	return pq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Permission string `json:"Permission,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Patientrightstype.Query().
//		GroupBy(patientrightstype.FieldPermission).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pq *PatientrightstypeQuery) GroupBy(field string, fields ...string) *PatientrightstypeGroupBy {
	group := &PatientrightstypeGroupBy{config: pq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Permission string `json:"Permission,omitempty"`
//	}
//
//	client.Patientrightstype.Query().
//		Select(patientrightstype.FieldPermission).
//		Scan(ctx, &v)
//
func (pq *PatientrightstypeQuery) Select(field string, fields ...string) *PatientrightstypeSelect {
	selector := &PatientrightstypeSelect{config: pq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(), nil
	}
	return selector
}

func (pq *PatientrightstypeQuery) prepareQuery(ctx context.Context) error {
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PatientrightstypeQuery) sqlAll(ctx context.Context) ([]*Patientrightstype, error) {
	var (
		nodes       = []*Patientrightstype{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [2]bool{
			pq.withPatientrightstypePatientrights != nil,
			pq.withPatientrightstypeAbilitypatientrights != nil,
		}
	)
	if pq.withPatientrightstypeAbilitypatientrights != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, patientrightstype.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Patientrightstype{config: pq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pq.withPatientrightstypePatientrights; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Patientrightstype)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Patientrights(func(s *sql.Selector) {
			s.Where(sql.InValues(patientrightstype.PatientrightstypePatientrightsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.Patientrightstype_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "Patientrightstype_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "Patientrightstype_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.PatientrightstypePatientrights = append(node.Edges.PatientrightstypePatientrights, n)
		}
	}

	if query := pq.withPatientrightstypeAbilitypatientrights; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Patientrightstype)
		for i := range nodes {
			if fk := nodes[i].Abilitypatientrights_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(abilitypatientrights.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "Abilitypatientrights_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.PatientrightstypeAbilitypatientrights = n
			}
		}
	}

	return nodes, nil
}

func (pq *PatientrightstypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PatientrightstypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (pq *PatientrightstypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patientrightstype.Table,
			Columns: patientrightstype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patientrightstype.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PatientrightstypeQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(patientrightstype.Table)
	selector := builder.Select(t1.Columns(patientrightstype.Columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(patientrightstype.Columns...)...)
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PatientrightstypeGroupBy is the builder for group-by Patientrightstype entities.
type PatientrightstypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PatientrightstypeGroupBy) Aggregate(fns ...AggregateFunc) *PatientrightstypeGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scan the result into the given value.
func (pgb *PatientrightstypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pgb.path(ctx)
	if err != nil {
		return err
	}
	pgb.sql = query
	return pgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pgb *PatientrightstypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrightstypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PatientrightstypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pgb *PatientrightstypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := pgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrightstypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrightstype.Label}
	default:
		err = fmt.Errorf("ent: PatientrightstypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pgb *PatientrightstypeGroupBy) StringX(ctx context.Context) string {
	v, err := pgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrightstypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PatientrightstypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pgb *PatientrightstypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := pgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrightstypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrightstype.Label}
	default:
		err = fmt.Errorf("ent: PatientrightstypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pgb *PatientrightstypeGroupBy) IntX(ctx context.Context) int {
	v, err := pgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrightstypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PatientrightstypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pgb *PatientrightstypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrightstypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrightstype.Label}
	default:
		err = fmt.Errorf("ent: PatientrightstypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pgb *PatientrightstypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrightstypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PatientrightstypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pgb *PatientrightstypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrightstypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrightstype.Label}
	default:
		err = fmt.Errorf("ent: PatientrightstypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pgb *PatientrightstypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := pgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pgb *PatientrightstypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pgb.sqlQuery().Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *PatientrightstypeGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql
	columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
	columns = append(columns, pgb.fields...)
	for _, fn := range pgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(pgb.fields...)
}

// PatientrightstypeSelect is the builder for select fields of Patientrightstype entities.
type PatientrightstypeSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ps *PatientrightstypeSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ps.path(ctx)
	if err != nil {
		return err
	}
	ps.sql = query
	return ps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ps *PatientrightstypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ps *PatientrightstypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PatientrightstypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ps *PatientrightstypeSelect) StringsX(ctx context.Context) []string {
	v, err := ps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ps *PatientrightstypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrightstype.Label}
	default:
		err = fmt.Errorf("ent: PatientrightstypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ps *PatientrightstypeSelect) StringX(ctx context.Context) string {
	v, err := ps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ps *PatientrightstypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PatientrightstypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ps *PatientrightstypeSelect) IntsX(ctx context.Context) []int {
	v, err := ps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ps *PatientrightstypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrightstype.Label}
	default:
		err = fmt.Errorf("ent: PatientrightstypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ps *PatientrightstypeSelect) IntX(ctx context.Context) int {
	v, err := ps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ps *PatientrightstypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PatientrightstypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ps *PatientrightstypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ps *PatientrightstypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrightstype.Label}
	default:
		err = fmt.Errorf("ent: PatientrightstypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ps *PatientrightstypeSelect) Float64X(ctx context.Context) float64 {
	v, err := ps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ps *PatientrightstypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PatientrightstypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ps *PatientrightstypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := ps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ps *PatientrightstypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrightstype.Label}
	default:
		err = fmt.Errorf("ent: PatientrightstypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ps *PatientrightstypeSelect) BoolX(ctx context.Context) bool {
	v, err := ps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ps *PatientrightstypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ps.sqlQuery().Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ps *PatientrightstypeSelect) sqlQuery() sql.Querier {
	selector := ps.sql
	selector.Select(selector.Columns(ps.fields...)...)
	return selector
}
