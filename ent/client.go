// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"recomCore/ent/migrate"

	"recomCore/ent/attribute"
	"recomCore/ent/product"
	"recomCore/ent/rate"
	"recomCore/ent/user"
	"recomCore/ent/vector"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Attribute is the client for interacting with the Attribute builders.
	Attribute *AttributeClient
	// Product is the client for interacting with the Product builders.
	Product *ProductClient
	// Rate is the client for interacting with the Rate builders.
	Rate *RateClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// Vector is the client for interacting with the Vector builders.
	Vector *VectorClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Attribute = NewAttributeClient(c.config)
	c.Product = NewProductClient(c.config)
	c.Rate = NewRateClient(c.config)
	c.User = NewUserClient(c.config)
	c.Vector = NewVectorClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Attribute: NewAttributeClient(cfg),
		Product:   NewProductClient(cfg),
		Rate:      NewRateClient(cfg),
		User:      NewUserClient(cfg),
		Vector:    NewVectorClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		ctx:       ctx,
		config:    cfg,
		Attribute: NewAttributeClient(cfg),
		Product:   NewProductClient(cfg),
		Rate:      NewRateClient(cfg),
		User:      NewUserClient(cfg),
		Vector:    NewVectorClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Attribute.
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
	c.Attribute.Use(hooks...)
	c.Product.Use(hooks...)
	c.Rate.Use(hooks...)
	c.User.Use(hooks...)
	c.Vector.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Attribute.Intercept(interceptors...)
	c.Product.Intercept(interceptors...)
	c.Rate.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
	c.Vector.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AttributeMutation:
		return c.Attribute.mutate(ctx, m)
	case *ProductMutation:
		return c.Product.mutate(ctx, m)
	case *RateMutation:
		return c.Rate.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	case *VectorMutation:
		return c.Vector.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// AttributeClient is a client for the Attribute schema.
type AttributeClient struct {
	config
}

// NewAttributeClient returns a client for the Attribute from the given config.
func NewAttributeClient(c config) *AttributeClient {
	return &AttributeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `attribute.Hooks(f(g(h())))`.
func (c *AttributeClient) Use(hooks ...Hook) {
	c.hooks.Attribute = append(c.hooks.Attribute, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `attribute.Intercept(f(g(h())))`.
func (c *AttributeClient) Intercept(interceptors ...Interceptor) {
	c.inters.Attribute = append(c.inters.Attribute, interceptors...)
}

// Create returns a builder for creating a Attribute entity.
func (c *AttributeClient) Create() *AttributeCreate {
	mutation := newAttributeMutation(c.config, OpCreate)
	return &AttributeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Attribute entities.
func (c *AttributeClient) CreateBulk(builders ...*AttributeCreate) *AttributeCreateBulk {
	return &AttributeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Attribute.
func (c *AttributeClient) Update() *AttributeUpdate {
	mutation := newAttributeMutation(c.config, OpUpdate)
	return &AttributeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AttributeClient) UpdateOne(a *Attribute) *AttributeUpdateOne {
	mutation := newAttributeMutation(c.config, OpUpdateOne, withAttribute(a))
	return &AttributeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AttributeClient) UpdateOneID(id int) *AttributeUpdateOne {
	mutation := newAttributeMutation(c.config, OpUpdateOne, withAttributeID(id))
	return &AttributeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Attribute.
func (c *AttributeClient) Delete() *AttributeDelete {
	mutation := newAttributeMutation(c.config, OpDelete)
	return &AttributeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AttributeClient) DeleteOne(a *Attribute) *AttributeDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AttributeClient) DeleteOneID(id int) *AttributeDeleteOne {
	builder := c.Delete().Where(attribute.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AttributeDeleteOne{builder}
}

// Query returns a query builder for Attribute.
func (c *AttributeClient) Query() *AttributeQuery {
	return &AttributeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAttribute},
		inters: c.Interceptors(),
	}
}

// Get returns a Attribute entity by its id.
func (c *AttributeClient) Get(ctx context.Context, id int) (*Attribute, error) {
	return c.Query().Where(attribute.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AttributeClient) GetX(ctx context.Context, id int) *Attribute {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Attribute.
func (c *AttributeClient) QueryOwner(a *Attribute) *ProductQuery {
	query := (&ProductClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(attribute.Table, attribute.FieldID, id),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attribute.OwnerTable, attribute.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AttributeClient) Hooks() []Hook {
	return c.hooks.Attribute
}

// Interceptors returns the client interceptors.
func (c *AttributeClient) Interceptors() []Interceptor {
	return c.inters.Attribute
}

func (c *AttributeClient) mutate(ctx context.Context, m *AttributeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AttributeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AttributeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AttributeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AttributeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Attribute mutation op: %q", m.Op())
	}
}

// ProductClient is a client for the Product schema.
type ProductClient struct {
	config
}

// NewProductClient returns a client for the Product from the given config.
func NewProductClient(c config) *ProductClient {
	return &ProductClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `product.Hooks(f(g(h())))`.
func (c *ProductClient) Use(hooks ...Hook) {
	c.hooks.Product = append(c.hooks.Product, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `product.Intercept(f(g(h())))`.
func (c *ProductClient) Intercept(interceptors ...Interceptor) {
	c.inters.Product = append(c.inters.Product, interceptors...)
}

// Create returns a builder for creating a Product entity.
func (c *ProductClient) Create() *ProductCreate {
	mutation := newProductMutation(c.config, OpCreate)
	return &ProductCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Product entities.
func (c *ProductClient) CreateBulk(builders ...*ProductCreate) *ProductCreateBulk {
	return &ProductCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Product.
func (c *ProductClient) Update() *ProductUpdate {
	mutation := newProductMutation(c.config, OpUpdate)
	return &ProductUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProductClient) UpdateOne(pr *Product) *ProductUpdateOne {
	mutation := newProductMutation(c.config, OpUpdateOne, withProduct(pr))
	return &ProductUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProductClient) UpdateOneID(id int) *ProductUpdateOne {
	mutation := newProductMutation(c.config, OpUpdateOne, withProductID(id))
	return &ProductUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Product.
func (c *ProductClient) Delete() *ProductDelete {
	mutation := newProductMutation(c.config, OpDelete)
	return &ProductDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ProductClient) DeleteOne(pr *Product) *ProductDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ProductClient) DeleteOneID(id int) *ProductDeleteOne {
	builder := c.Delete().Where(product.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProductDeleteOne{builder}
}

// Query returns a query builder for Product.
func (c *ProductClient) Query() *ProductQuery {
	return &ProductQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeProduct},
		inters: c.Interceptors(),
	}
}

// Get returns a Product entity by its id.
func (c *ProductClient) Get(ctx context.Context, id int) (*Product, error) {
	return c.Query().Where(product.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProductClient) GetX(ctx context.Context, id int) *Product {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAttributes queries the attributes edge of a Product.
func (c *ProductClient) QueryAttributes(pr *Product) *AttributeQuery {
	query := (&AttributeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(product.Table, product.FieldID, id),
			sqlgraph.To(attribute.Table, attribute.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, product.AttributesTable, product.AttributesColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRates queries the rates edge of a Product.
func (c *ProductClient) QueryRates(pr *Product) *RateQuery {
	query := (&RateClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(product.Table, product.FieldID, id),
			sqlgraph.To(rate.Table, rate.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, product.RatesTable, product.RatesColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryVectors queries the vectors edge of a Product.
func (c *ProductClient) QueryVectors(pr *Product) *VectorQuery {
	query := (&VectorClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(product.Table, product.FieldID, id),
			sqlgraph.To(vector.Table, vector.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, product.VectorsTable, product.VectorsColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProductClient) Hooks() []Hook {
	return c.hooks.Product
}

// Interceptors returns the client interceptors.
func (c *ProductClient) Interceptors() []Interceptor {
	return c.inters.Product
}

func (c *ProductClient) mutate(ctx context.Context, m *ProductMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ProductCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ProductUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ProductUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ProductDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Product mutation op: %q", m.Op())
	}
}

// RateClient is a client for the Rate schema.
type RateClient struct {
	config
}

// NewRateClient returns a client for the Rate from the given config.
func NewRateClient(c config) *RateClient {
	return &RateClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `rate.Hooks(f(g(h())))`.
func (c *RateClient) Use(hooks ...Hook) {
	c.hooks.Rate = append(c.hooks.Rate, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `rate.Intercept(f(g(h())))`.
func (c *RateClient) Intercept(interceptors ...Interceptor) {
	c.inters.Rate = append(c.inters.Rate, interceptors...)
}

// Create returns a builder for creating a Rate entity.
func (c *RateClient) Create() *RateCreate {
	mutation := newRateMutation(c.config, OpCreate)
	return &RateCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Rate entities.
func (c *RateClient) CreateBulk(builders ...*RateCreate) *RateCreateBulk {
	return &RateCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Rate.
func (c *RateClient) Update() *RateUpdate {
	mutation := newRateMutation(c.config, OpUpdate)
	return &RateUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RateClient) UpdateOne(r *Rate) *RateUpdateOne {
	mutation := newRateMutation(c.config, OpUpdateOne, withRate(r))
	return &RateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RateClient) UpdateOneID(id int) *RateUpdateOne {
	mutation := newRateMutation(c.config, OpUpdateOne, withRateID(id))
	return &RateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Rate.
func (c *RateClient) Delete() *RateDelete {
	mutation := newRateMutation(c.config, OpDelete)
	return &RateDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RateClient) DeleteOne(r *Rate) *RateDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RateClient) DeleteOneID(id int) *RateDeleteOne {
	builder := c.Delete().Where(rate.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RateDeleteOne{builder}
}

// Query returns a query builder for Rate.
func (c *RateClient) Query() *RateQuery {
	return &RateQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRate},
		inters: c.Interceptors(),
	}
}

// Get returns a Rate entity by its id.
func (c *RateClient) Get(ctx context.Context, id int) (*Rate, error) {
	return c.Query().Where(rate.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RateClient) GetX(ctx context.Context, id int) *Rate {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Rate.
func (c *RateClient) QueryOwner(r *Rate) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(rate.Table, rate.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, rate.OwnerTable, rate.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySubject queries the subject edge of a Rate.
func (c *RateClient) QuerySubject(r *Rate) *ProductQuery {
	query := (&ProductClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(rate.Table, rate.FieldID, id),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, rate.SubjectTable, rate.SubjectColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RateClient) Hooks() []Hook {
	return c.hooks.Rate
}

// Interceptors returns the client interceptors.
func (c *RateClient) Interceptors() []Interceptor {
	return c.inters.Rate
}

func (c *RateClient) mutate(ctx context.Context, m *RateMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RateCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RateUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RateDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Rate mutation op: %q", m.Op())
	}
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

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
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

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
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
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
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

// QueryRates queries the rates edge of a User.
func (c *UserClient) QueryRates(u *User) *RateQuery {
	query := (&RateClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(rate.Table, rate.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.RatesTable, user.RatesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// VectorClient is a client for the Vector schema.
type VectorClient struct {
	config
}

// NewVectorClient returns a client for the Vector from the given config.
func NewVectorClient(c config) *VectorClient {
	return &VectorClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `vector.Hooks(f(g(h())))`.
func (c *VectorClient) Use(hooks ...Hook) {
	c.hooks.Vector = append(c.hooks.Vector, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `vector.Intercept(f(g(h())))`.
func (c *VectorClient) Intercept(interceptors ...Interceptor) {
	c.inters.Vector = append(c.inters.Vector, interceptors...)
}

// Create returns a builder for creating a Vector entity.
func (c *VectorClient) Create() *VectorCreate {
	mutation := newVectorMutation(c.config, OpCreate)
	return &VectorCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Vector entities.
func (c *VectorClient) CreateBulk(builders ...*VectorCreate) *VectorCreateBulk {
	return &VectorCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Vector.
func (c *VectorClient) Update() *VectorUpdate {
	mutation := newVectorMutation(c.config, OpUpdate)
	return &VectorUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VectorClient) UpdateOne(v *Vector) *VectorUpdateOne {
	mutation := newVectorMutation(c.config, OpUpdateOne, withVector(v))
	return &VectorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VectorClient) UpdateOneID(id int) *VectorUpdateOne {
	mutation := newVectorMutation(c.config, OpUpdateOne, withVectorID(id))
	return &VectorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Vector.
func (c *VectorClient) Delete() *VectorDelete {
	mutation := newVectorMutation(c.config, OpDelete)
	return &VectorDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *VectorClient) DeleteOne(v *Vector) *VectorDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *VectorClient) DeleteOneID(id int) *VectorDeleteOne {
	builder := c.Delete().Where(vector.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VectorDeleteOne{builder}
}

// Query returns a query builder for Vector.
func (c *VectorClient) Query() *VectorQuery {
	return &VectorQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeVector},
		inters: c.Interceptors(),
	}
}

// Get returns a Vector entity by its id.
func (c *VectorClient) Get(ctx context.Context, id int) (*Vector, error) {
	return c.Query().Where(vector.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VectorClient) GetX(ctx context.Context, id int) *Vector {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Vector.
func (c *VectorClient) QueryOwner(v *Vector) *ProductQuery {
	query := (&ProductClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(vector.Table, vector.FieldID, id),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, vector.OwnerTable, vector.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *VectorClient) Hooks() []Hook {
	return c.hooks.Vector
}

// Interceptors returns the client interceptors.
func (c *VectorClient) Interceptors() []Interceptor {
	return c.inters.Vector
}

func (c *VectorClient) mutate(ctx context.Context, m *VectorMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&VectorCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&VectorUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&VectorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&VectorDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Vector mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Attribute, Product, Rate, User, Vector []ent.Hook
	}
	inters struct {
		Attribute, Product, Rate, User, Vector []ent.Interceptor
	}
)
