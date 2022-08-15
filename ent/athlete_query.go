// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"locathlete-server/ent/athlete"
	"locathlete-server/ent/athleteschool"
	"locathlete-server/ent/predicate"
	"locathlete-server/ent/school"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AthleteQuery is the builder for querying Athlete entities.
type AthleteQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Athlete
	// eager-loading edges.
	withSchools        *SchoolQuery
	withAthleteSchools *AthleteSchoolQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AthleteQuery builder.
func (aq *AthleteQuery) Where(ps ...predicate.Athlete) *AthleteQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *AthleteQuery) Limit(limit int) *AthleteQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *AthleteQuery) Offset(offset int) *AthleteQuery {
	aq.offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AthleteQuery) Unique(unique bool) *AthleteQuery {
	aq.unique = &unique
	return aq
}

// Order adds an order step to the query.
func (aq *AthleteQuery) Order(o ...OrderFunc) *AthleteQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QuerySchools chains the current query on the "schools" edge.
func (aq *AthleteQuery) QuerySchools() *SchoolQuery {
	query := &SchoolQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(athlete.Table, athlete.FieldID, selector),
			sqlgraph.To(school.Table, school.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, athlete.SchoolsTable, athlete.SchoolsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAthleteSchools chains the current query on the "athlete_schools" edge.
func (aq *AthleteQuery) QueryAthleteSchools() *AthleteSchoolQuery {
	query := &AthleteSchoolQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(athlete.Table, athlete.FieldID, selector),
			sqlgraph.To(athleteschool.Table, athleteschool.AthleteColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, athlete.AthleteSchoolsTable, athlete.AthleteSchoolsColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Athlete entity from the query.
// Returns a *NotFoundError when no Athlete was found.
func (aq *AthleteQuery) First(ctx context.Context) (*Athlete, error) {
	nodes, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{athlete.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AthleteQuery) FirstX(ctx context.Context) *Athlete {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Athlete ID from the query.
// Returns a *NotFoundError when no Athlete ID was found.
func (aq *AthleteQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{athlete.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AthleteQuery) FirstIDX(ctx context.Context) int {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Athlete entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Athlete entity is found.
// Returns a *NotFoundError when no Athlete entities are found.
func (aq *AthleteQuery) Only(ctx context.Context) (*Athlete, error) {
	nodes, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{athlete.Label}
	default:
		return nil, &NotSingularError{athlete.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AthleteQuery) OnlyX(ctx context.Context) *Athlete {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Athlete ID in the query.
// Returns a *NotSingularError when more than one Athlete ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AthleteQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{athlete.Label}
	default:
		err = &NotSingularError{athlete.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AthleteQuery) OnlyIDX(ctx context.Context) int {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Athletes.
func (aq *AthleteQuery) All(ctx context.Context) ([]*Athlete, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *AthleteQuery) AllX(ctx context.Context) []*Athlete {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Athlete IDs.
func (aq *AthleteQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := aq.Select(athlete.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AthleteQuery) IDsX(ctx context.Context) []int {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AthleteQuery) Count(ctx context.Context) (int, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AthleteQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AthleteQuery) Exist(ctx context.Context) (bool, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AthleteQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AthleteQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AthleteQuery) Clone() *AthleteQuery {
	if aq == nil {
		return nil
	}
	return &AthleteQuery{
		config:             aq.config,
		limit:              aq.limit,
		offset:             aq.offset,
		order:              append([]OrderFunc{}, aq.order...),
		predicates:         append([]predicate.Athlete{}, aq.predicates...),
		withSchools:        aq.withSchools.Clone(),
		withAthleteSchools: aq.withAthleteSchools.Clone(),
		// clone intermediate query.
		sql:    aq.sql.Clone(),
		path:   aq.path,
		unique: aq.unique,
	}
}

// WithSchools tells the query-builder to eager-load the nodes that are connected to
// the "schools" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AthleteQuery) WithSchools(opts ...func(*SchoolQuery)) *AthleteQuery {
	query := &SchoolQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withSchools = query
	return aq
}

// WithAthleteSchools tells the query-builder to eager-load the nodes that are connected to
// the "athlete_schools" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AthleteQuery) WithAthleteSchools(opts ...func(*AthleteSchoolQuery)) *AthleteQuery {
	query := &AthleteSchoolQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withAthleteSchools = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Bio string `json:"bio,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Athlete.Query().
//		GroupBy(athlete.FieldBio).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AthleteQuery) GroupBy(field string, fields ...string) *AthleteGroupBy {
	grbuild := &AthleteGroupBy{config: aq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(ctx), nil
	}
	grbuild.label = athlete.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Bio string `json:"bio,omitempty"`
//	}
//
//	client.Athlete.Query().
//		Select(athlete.FieldBio).
//		Scan(ctx, &v)
func (aq *AthleteQuery) Select(fields ...string) *AthleteSelect {
	aq.fields = append(aq.fields, fields...)
	selbuild := &AthleteSelect{AthleteQuery: aq}
	selbuild.label = athlete.Label
	selbuild.flds, selbuild.scan = &aq.fields, selbuild.Scan
	return selbuild
}

func (aq *AthleteQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aq.fields {
		if !athlete.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AthleteQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Athlete, error) {
	var (
		nodes       = []*Athlete{}
		_spec       = aq.querySpec()
		loadedTypes = [2]bool{
			aq.withSchools != nil,
			aq.withAthleteSchools != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Athlete).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Athlete{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := aq.withSchools; query != nil {
		edgeids := make([]driver.Value, len(nodes))
		byid := make(map[int]*Athlete)
		nids := make(map[int]map[*Athlete]struct{})
		for i, node := range nodes {
			edgeids[i] = node.ID
			byid[node.ID] = node
			node.Edges.Schools = []*School{}
		}
		query.Where(func(s *sql.Selector) {
			joinT := sql.Table(athlete.SchoolsTable)
			s.Join(joinT).On(s.C(school.FieldID), joinT.C(athlete.SchoolsPrimaryKey[1]))
			s.Where(sql.InValues(joinT.C(athlete.SchoolsPrimaryKey[0]), edgeids...))
			columns := s.SelectedColumns()
			s.Select(joinT.C(athlete.SchoolsPrimaryKey[0]))
			s.AppendSelect(columns...)
			s.SetDistinct(false)
		})
		neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]interface{}, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]interface{}{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []interface{}) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Athlete]struct{}{byid[outValue]: struct{}{}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byid[outValue]] = struct{}{}
				return nil
			}
		})
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "schools" node returned %v`, n.ID)
			}
			for kn := range nodes {
				kn.Edges.Schools = append(kn.Edges.Schools, n)
			}
		}
	}

	if query := aq.withAthleteSchools; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Athlete)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.AthleteSchools = []*AthleteSchool{}
		}
		query.Where(predicate.AthleteSchool(func(s *sql.Selector) {
			s.Where(sql.InValues(athlete.AthleteSchoolsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.AthleteID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "athlete_id" returned %v for node %v`, fk, n)
			}
			node.Edges.AthleteSchools = append(node.Edges.AthleteSchools, n)
		}
	}

	return nodes, nil
}

func (aq *AthleteQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.fields
	if len(aq.fields) > 0 {
		_spec.Unique = aq.unique != nil && *aq.unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AthleteQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (aq *AthleteQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   athlete.Table,
			Columns: athlete.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: athlete.FieldID,
			},
		},
		From:   aq.sql,
		Unique: true,
	}
	if unique := aq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, athlete.FieldID)
		for i := range fields {
			if fields[i] != athlete.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AthleteQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(athlete.Table)
	columns := aq.fields
	if len(columns) == 0 {
		columns = athlete.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.unique != nil && *aq.unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AthleteGroupBy is the group-by builder for Athlete entities.
type AthleteGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AthleteGroupBy) Aggregate(fns ...AggregateFunc) *AthleteGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scans the result into the given value.
func (agb *AthleteGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := agb.path(ctx)
	if err != nil {
		return err
	}
	agb.sql = query
	return agb.sqlScan(ctx, v)
}

func (agb *AthleteGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range agb.fields {
		if !athlete.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := agb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agb *AthleteGroupBy) sqlQuery() *sql.Selector {
	selector := agb.sql.Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(agb.fields)+len(agb.fns))
		for _, f := range agb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(agb.fields...)...)
}

// AthleteSelect is the builder for selecting fields of Athlete entities.
type AthleteSelect struct {
	*AthleteQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (as *AthleteSelect) Scan(ctx context.Context, v interface{}) error {
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	as.sql = as.AthleteQuery.sqlQuery(ctx)
	return as.sqlScan(ctx, v)
}

func (as *AthleteSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := as.sql.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
