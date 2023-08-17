// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"ginweb/ent/migrate"

	"ginweb/ent/department"
	"ginweb/ent/employee"
	"ginweb/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Department is the client for interacting with the Department builders.
	Department *DepartmentClient
	// Employee is the client for interacting with the Employee builders.
	Employee *EmployeeClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Department = NewDepartmentClient(c.config)
	c.Employee = NewEmployeeClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Department: NewDepartmentClient(cfg),
		Employee:   NewEmployeeClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:     cfg,
		Department: NewDepartmentClient(cfg),
		Employee:   NewEmployeeClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Department.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Department.Use(hooks...)
	c.Employee.Use(hooks...)
	c.User.Use(hooks...)
}

// DepartmentClient is a client for the Department schema.
type DepartmentClient struct {
	config
}

// NewDepartmentClient returns a client for the Department from the given config.
func NewDepartmentClient(c config) *DepartmentClient {
	return &DepartmentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `department.Hooks(f(g(h())))`.
func (c *DepartmentClient) Use(hooks ...Hook) {
	c.hooks.Department = append(c.hooks.Department, hooks...)
}

// Create returns a create builder for Department.
func (c *DepartmentClient) Create() *DepartmentCreate {
	mutation := newDepartmentMutation(c.config, OpCreate)
	return &DepartmentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Department entities.
func (c *DepartmentClient) CreateBulk(builders ...*DepartmentCreate) *DepartmentCreateBulk {
	return &DepartmentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Department.
func (c *DepartmentClient) Update() *DepartmentUpdate {
	mutation := newDepartmentMutation(c.config, OpUpdate)
	return &DepartmentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DepartmentClient) UpdateOne(d *Department) *DepartmentUpdateOne {
	mutation := newDepartmentMutation(c.config, OpUpdateOne, withDepartment(d))
	return &DepartmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DepartmentClient) UpdateOneID(id int) *DepartmentUpdateOne {
	mutation := newDepartmentMutation(c.config, OpUpdateOne, withDepartmentID(id))
	return &DepartmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Department.
func (c *DepartmentClient) Delete() *DepartmentDelete {
	mutation := newDepartmentMutation(c.config, OpDelete)
	return &DepartmentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DepartmentClient) DeleteOne(d *Department) *DepartmentDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DepartmentClient) DeleteOneID(id int) *DepartmentDeleteOne {
	builder := c.Delete().Where(department.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DepartmentDeleteOne{builder}
}

// Query returns a query builder for Department.
func (c *DepartmentClient) Query() *DepartmentQuery {
	return &DepartmentQuery{
		config: c.config,
	}
}

// Get returns a Department entity by its id.
func (c *DepartmentClient) Get(ctx context.Context, id int) (*Department, error) {
	return c.Query().Where(department.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DepartmentClient) GetX(ctx context.Context, id int) *Department {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryEmployees queries the employees edge of a Department.
func (c *DepartmentClient) QueryEmployees(d *Department) *EmployeeQuery {
	query := &EmployeeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(department.Table, department.FieldID, id),
			sqlgraph.To(employee.Table, employee.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, department.EmployeesTable, department.EmployeesColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DepartmentClient) Hooks() []Hook {
	return c.hooks.Department
}

// EmployeeClient is a client for the Employee schema.
type EmployeeClient struct {
	config
}

// NewEmployeeClient returns a client for the Employee from the given config.
func NewEmployeeClient(c config) *EmployeeClient {
	return &EmployeeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `employee.Hooks(f(g(h())))`.
func (c *EmployeeClient) Use(hooks ...Hook) {
	c.hooks.Employee = append(c.hooks.Employee, hooks...)
}

// Create returns a create builder for Employee.
func (c *EmployeeClient) Create() *EmployeeCreate {
	mutation := newEmployeeMutation(c.config, OpCreate)
	return &EmployeeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Employee entities.
func (c *EmployeeClient) CreateBulk(builders ...*EmployeeCreate) *EmployeeCreateBulk {
	return &EmployeeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Employee.
func (c *EmployeeClient) Update() *EmployeeUpdate {
	mutation := newEmployeeMutation(c.config, OpUpdate)
	return &EmployeeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EmployeeClient) UpdateOne(e *Employee) *EmployeeUpdateOne {
	mutation := newEmployeeMutation(c.config, OpUpdateOne, withEmployee(e))
	return &EmployeeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EmployeeClient) UpdateOneID(id int) *EmployeeUpdateOne {
	mutation := newEmployeeMutation(c.config, OpUpdateOne, withEmployeeID(id))
	return &EmployeeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Employee.
func (c *EmployeeClient) Delete() *EmployeeDelete {
	mutation := newEmployeeMutation(c.config, OpDelete)
	return &EmployeeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *EmployeeClient) DeleteOne(e *Employee) *EmployeeDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *EmployeeClient) DeleteOneID(id int) *EmployeeDeleteOne {
	builder := c.Delete().Where(employee.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EmployeeDeleteOne{builder}
}

// Query returns a query builder for Employee.
func (c *EmployeeClient) Query() *EmployeeQuery {
	return &EmployeeQuery{
		config: c.config,
	}
}

// Get returns a Employee entity by its id.
func (c *EmployeeClient) Get(ctx context.Context, id int) (*Employee, error) {
	return c.Query().Where(employee.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EmployeeClient) GetX(ctx context.Context, id int) *Employee {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDepartment queries the department edge of a Employee.
func (c *EmployeeClient) QueryDepartment(e *Employee) *DepartmentQuery {
	query := &DepartmentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(employee.Table, employee.FieldID, id),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, employee.DepartmentTable, employee.DepartmentColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EmployeeClient) Hooks() []Hook {
	return c.hooks.Employee
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
