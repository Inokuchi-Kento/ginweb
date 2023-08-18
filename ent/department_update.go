// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"ginweb/ent/department"
	"ginweb/ent/employee"
	"ginweb/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DepartmentUpdate is the builder for updating Department entities.
type DepartmentUpdate struct {
	config
	hooks    []Hook
	mutation *DepartmentMutation
}

// Where appends a list predicates to the DepartmentUpdate builder.
func (du *DepartmentUpdate) Where(ps ...predicate.Department) *DepartmentUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetName sets the "name" field.
func (du *DepartmentUpdate) SetName(s string) *DepartmentUpdate {
	du.mutation.SetName(s)
	return du
}

// AddEmployeeIDs adds the "employees" edge to the Employee entity by IDs.
func (du *DepartmentUpdate) AddEmployeeIDs(ids ...int) *DepartmentUpdate {
	du.mutation.AddEmployeeIDs(ids...)
	return du
}

// AddEmployees adds the "employees" edges to the Employee entity.
func (du *DepartmentUpdate) AddEmployees(e ...*Employee) *DepartmentUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return du.AddEmployeeIDs(ids...)
}

// Mutation returns the DepartmentMutation object of the builder.
func (du *DepartmentUpdate) Mutation() *DepartmentMutation {
	return du.mutation
}

// ClearEmployees clears all "employees" edges to the Employee entity.
func (du *DepartmentUpdate) ClearEmployees() *DepartmentUpdate {
	du.mutation.ClearEmployees()
	return du
}

// RemoveEmployeeIDs removes the "employees" edge to Employee entities by IDs.
func (du *DepartmentUpdate) RemoveEmployeeIDs(ids ...int) *DepartmentUpdate {
	du.mutation.RemoveEmployeeIDs(ids...)
	return du
}

// RemoveEmployees removes "employees" edges to Employee entities.
func (du *DepartmentUpdate) RemoveEmployees(e ...*Employee) *DepartmentUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return du.RemoveEmployeeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DepartmentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DepartmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			if du.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DepartmentUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DepartmentUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DepartmentUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DepartmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   department.Table,
			Columns: department.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: department.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: department.FieldName,
		})
	}
	if du.mutation.EmployeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.EmployeesTable,
			Columns: []string{department.EmployeesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedEmployeesIDs(); len(nodes) > 0 && !du.mutation.EmployeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.EmployeesTable,
			Columns: []string{department.EmployeesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.EmployeesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.EmployeesTable,
			Columns: []string{department.EmployeesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{department.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DepartmentUpdateOne is the builder for updating a single Department entity.
type DepartmentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DepartmentMutation
}

// SetName sets the "name" field.
func (duo *DepartmentUpdateOne) SetName(s string) *DepartmentUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// AddEmployeeIDs adds the "employees" edge to the Employee entity by IDs.
func (duo *DepartmentUpdateOne) AddEmployeeIDs(ids ...int) *DepartmentUpdateOne {
	duo.mutation.AddEmployeeIDs(ids...)
	return duo
}

// AddEmployees adds the "employees" edges to the Employee entity.
func (duo *DepartmentUpdateOne) AddEmployees(e ...*Employee) *DepartmentUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return duo.AddEmployeeIDs(ids...)
}

// Mutation returns the DepartmentMutation object of the builder.
func (duo *DepartmentUpdateOne) Mutation() *DepartmentMutation {
	return duo.mutation
}

// ClearEmployees clears all "employees" edges to the Employee entity.
func (duo *DepartmentUpdateOne) ClearEmployees() *DepartmentUpdateOne {
	duo.mutation.ClearEmployees()
	return duo
}

// RemoveEmployeeIDs removes the "employees" edge to Employee entities by IDs.
func (duo *DepartmentUpdateOne) RemoveEmployeeIDs(ids ...int) *DepartmentUpdateOne {
	duo.mutation.RemoveEmployeeIDs(ids...)
	return duo
}

// RemoveEmployees removes "employees" edges to Employee entities.
func (duo *DepartmentUpdateOne) RemoveEmployees(e ...*Employee) *DepartmentUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return duo.RemoveEmployeeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DepartmentUpdateOne) Select(field string, fields ...string) *DepartmentUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Department entity.
func (duo *DepartmentUpdateOne) Save(ctx context.Context) (*Department, error) {
	var (
		err  error
		node *Department
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DepartmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			if duo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DepartmentUpdateOne) SaveX(ctx context.Context) *Department {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DepartmentUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DepartmentUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DepartmentUpdateOne) sqlSave(ctx context.Context) (_node *Department, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   department.Table,
			Columns: department.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: department.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Department.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, department.FieldID)
		for _, f := range fields {
			if !department.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != department.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: department.FieldName,
		})
	}
	if duo.mutation.EmployeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.EmployeesTable,
			Columns: []string{department.EmployeesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedEmployeesIDs(); len(nodes) > 0 && !duo.mutation.EmployeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.EmployeesTable,
			Columns: []string{department.EmployeesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.EmployeesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.EmployeesTable,
			Columns: []string{department.EmployeesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Department{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{department.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
