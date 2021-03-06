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
	"github.com/team10/app/ent/predicate"
	"github.com/team10/app/ent/user"
	"github.com/team10/app/ent/userstatus"
)

// UserstatusQuery is the builder for querying Userstatus entities.
type UserstatusQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Userstatus
	// eager-loading edges.
	withUser *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (uq *UserstatusQuery) Where(ps ...predicate.Userstatus) *UserstatusQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit adds a limit step to the query.
func (uq *UserstatusQuery) Limit(limit int) *UserstatusQuery {
	uq.limit = &limit
	return uq
}

// Offset adds an offset step to the query.
func (uq *UserstatusQuery) Offset(offset int) *UserstatusQuery {
	uq.offset = &offset
	return uq
}

// Order adds an order step to the query.
func (uq *UserstatusQuery) Order(o ...OrderFunc) *UserstatusQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// QueryUser chains the current query on the user edge.
func (uq *UserstatusQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: uq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userstatus.Table, userstatus.FieldID, uq.sqlQuery()),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, userstatus.UserTable, userstatus.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Userstatus entity in the query. Returns *NotFoundError when no userstatus was found.
func (uq *UserstatusQuery) First(ctx context.Context) (*Userstatus, error) {
	us, err := uq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(us) == 0 {
		return nil, &NotFoundError{userstatus.Label}
	}
	return us[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UserstatusQuery) FirstX(ctx context.Context) *Userstatus {
	u, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return u
}

// FirstID returns the first Userstatus id in the query. Returns *NotFoundError when no id was found.
func (uq *UserstatusQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userstatus.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (uq *UserstatusQuery) FirstXID(ctx context.Context) int {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Userstatus entity in the query, returns an error if not exactly one entity was returned.
func (uq *UserstatusQuery) Only(ctx context.Context) (*Userstatus, error) {
	us, err := uq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(us) {
	case 1:
		return us[0], nil
	case 0:
		return nil, &NotFoundError{userstatus.Label}
	default:
		return nil, &NotSingularError{userstatus.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UserstatusQuery) OnlyX(ctx context.Context) *Userstatus {
	u, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

// OnlyID returns the only Userstatus id in the query, returns an error if not exactly one id was returned.
func (uq *UserstatusQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userstatus.Label}
	default:
		err = &NotSingularError{userstatus.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uq *UserstatusQuery) OnlyIDX(ctx context.Context) int {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Userstatuses.
func (uq *UserstatusQuery) All(ctx context.Context) ([]*Userstatus, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return uq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uq *UserstatusQuery) AllX(ctx context.Context) []*Userstatus {
	us, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return us
}

// IDs executes the query and returns a list of Userstatus ids.
func (uq *UserstatusQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := uq.Select(userstatus.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UserstatusQuery) IDsX(ctx context.Context) []int {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UserstatusQuery) Count(ctx context.Context) (int, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return uq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UserstatusQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UserstatusQuery) Exist(ctx context.Context) (bool, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return uq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UserstatusQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UserstatusQuery) Clone() *UserstatusQuery {
	return &UserstatusQuery{
		config:     uq.config,
		limit:      uq.limit,
		offset:     uq.offset,
		order:      append([]OrderFunc{}, uq.order...),
		unique:     append([]string{}, uq.unique...),
		predicates: append([]predicate.Userstatus{}, uq.predicates...),
		// clone intermediate query.
		sql:  uq.sql.Clone(),
		path: uq.path,
	}
}

//  WithUser tells the query-builder to eager-loads the nodes that are connected to
// the "user" edge. The optional arguments used to configure the query builder of the edge.
func (uq *UserstatusQuery) WithUser(opts ...func(*UserQuery)) *UserstatusQuery {
	query := &UserQuery{config: uq.config}
	for _, opt := range opts {
		opt(query)
	}
	uq.withUser = query
	return uq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Userstatus string `json:"Userstatus,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Userstatus.Query().
//		GroupBy(userstatus.FieldUserstatus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (uq *UserstatusQuery) GroupBy(field string, fields ...string) *UserstatusGroupBy {
	group := &UserstatusGroupBy{config: uq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return uq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Userstatus string `json:"Userstatus,omitempty"`
//	}
//
//	client.Userstatus.Query().
//		Select(userstatus.FieldUserstatus).
//		Scan(ctx, &v)
//
func (uq *UserstatusQuery) Select(field string, fields ...string) *UserstatusSelect {
	selector := &UserstatusSelect{config: uq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return uq.sqlQuery(), nil
	}
	return selector
}

func (uq *UserstatusQuery) prepareQuery(ctx context.Context) error {
	if uq.path != nil {
		prev, err := uq.path(ctx)
		if err != nil {
			return err
		}
		uq.sql = prev
	}
	return nil
}

func (uq *UserstatusQuery) sqlAll(ctx context.Context) ([]*Userstatus, error) {
	var (
		nodes       = []*Userstatus{}
		_spec       = uq.querySpec()
		loadedTypes = [1]bool{
			uq.withUser != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Userstatus{config: uq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
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
	if err := sqlgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := uq.withUser; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Userstatus)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.User(func(s *sql.Selector) {
			s.Where(sql.InValues(userstatus.UserColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.userstatus_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "userstatus_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "userstatus_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.User = append(node.Edges.User, n)
		}
	}

	return nodes, nil
}

func (uq *UserstatusQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	return sqlgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *UserstatusQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := uq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (uq *UserstatusQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userstatus.Table,
			Columns: userstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userstatus.FieldID,
			},
		},
		From:   uq.sql,
		Unique: true,
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uq *UserstatusQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(uq.driver.Dialect())
	t1 := builder.Table(userstatus.Table)
	selector := builder.Select(t1.Columns(userstatus.Columns...)...).From(t1)
	if uq.sql != nil {
		selector = uq.sql
		selector.Select(selector.Columns(userstatus.Columns...)...)
	}
	for _, p := range uq.predicates {
		p(selector)
	}
	for _, p := range uq.order {
		p(selector)
	}
	if offset := uq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserstatusGroupBy is the builder for group-by Userstatus entities.
type UserstatusGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UserstatusGroupBy) Aggregate(fns ...AggregateFunc) *UserstatusGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the group-by query and scan the result into the given value.
func (ugb *UserstatusGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ugb.path(ctx)
	if err != nil {
		return err
	}
	ugb.sql = query
	return ugb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ugb *UserstatusGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ugb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ugb *UserstatusGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UserstatusGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ugb *UserstatusGroupBy) StringsX(ctx context.Context) []string {
	v, err := ugb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (ugb *UserstatusGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ugb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userstatus.Label}
	default:
		err = fmt.Errorf("ent: UserstatusGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ugb *UserstatusGroupBy) StringX(ctx context.Context) string {
	v, err := ugb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ugb *UserstatusGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UserstatusGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ugb *UserstatusGroupBy) IntsX(ctx context.Context) []int {
	v, err := ugb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (ugb *UserstatusGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ugb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userstatus.Label}
	default:
		err = fmt.Errorf("ent: UserstatusGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ugb *UserstatusGroupBy) IntX(ctx context.Context) int {
	v, err := ugb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ugb *UserstatusGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UserstatusGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ugb *UserstatusGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ugb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (ugb *UserstatusGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ugb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userstatus.Label}
	default:
		err = fmt.Errorf("ent: UserstatusGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ugb *UserstatusGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ugb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ugb *UserstatusGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UserstatusGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ugb *UserstatusGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ugb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (ugb *UserstatusGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ugb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userstatus.Label}
	default:
		err = fmt.Errorf("ent: UserstatusGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ugb *UserstatusGroupBy) BoolX(ctx context.Context) bool {
	v, err := ugb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ugb *UserstatusGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ugb.sqlQuery().Query()
	if err := ugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ugb *UserstatusGroupBy) sqlQuery() *sql.Selector {
	selector := ugb.sql
	columns := make([]string, 0, len(ugb.fields)+len(ugb.fns))
	columns = append(columns, ugb.fields...)
	for _, fn := range ugb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(ugb.fields...)
}

// UserstatusSelect is the builder for select fields of Userstatus entities.
type UserstatusSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (us *UserstatusSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := us.path(ctx)
	if err != nil {
		return err
	}
	us.sql = query
	return us.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (us *UserstatusSelect) ScanX(ctx context.Context, v interface{}) {
	if err := us.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (us *UserstatusSelect) Strings(ctx context.Context) ([]string, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UserstatusSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (us *UserstatusSelect) StringsX(ctx context.Context) []string {
	v, err := us.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (us *UserstatusSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = us.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userstatus.Label}
	default:
		err = fmt.Errorf("ent: UserstatusSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (us *UserstatusSelect) StringX(ctx context.Context) string {
	v, err := us.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (us *UserstatusSelect) Ints(ctx context.Context) ([]int, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UserstatusSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (us *UserstatusSelect) IntsX(ctx context.Context) []int {
	v, err := us.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (us *UserstatusSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = us.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userstatus.Label}
	default:
		err = fmt.Errorf("ent: UserstatusSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (us *UserstatusSelect) IntX(ctx context.Context) int {
	v, err := us.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (us *UserstatusSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UserstatusSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (us *UserstatusSelect) Float64sX(ctx context.Context) []float64 {
	v, err := us.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (us *UserstatusSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = us.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userstatus.Label}
	default:
		err = fmt.Errorf("ent: UserstatusSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (us *UserstatusSelect) Float64X(ctx context.Context) float64 {
	v, err := us.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (us *UserstatusSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UserstatusSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (us *UserstatusSelect) BoolsX(ctx context.Context) []bool {
	v, err := us.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (us *UserstatusSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = us.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userstatus.Label}
	default:
		err = fmt.Errorf("ent: UserstatusSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (us *UserstatusSelect) BoolX(ctx context.Context) bool {
	v, err := us.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (us *UserstatusSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := us.sqlQuery().Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (us *UserstatusSelect) sqlQuery() sql.Querier {
	selector := us.sql
	selector.Select(selector.Columns(us.fields...)...)
	return selector
}
