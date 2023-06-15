// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recomCore/ent/product"
	"recomCore/ent/rate"
	"recomCore/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RateCreate is the builder for creating a Rate entity.
type RateCreate struct {
	config
	mutation *RateMutation
	hooks    []Hook
}

// SetRate sets the "rate" field.
func (rc *RateCreate) SetRate(b bool) *RateCreate {
	rc.mutation.SetRate(b)
	return rc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (rc *RateCreate) SetOwnerID(id int) *RateCreate {
	rc.mutation.SetOwnerID(id)
	return rc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (rc *RateCreate) SetNillableOwnerID(id *int) *RateCreate {
	if id != nil {
		rc = rc.SetOwnerID(*id)
	}
	return rc
}

// SetOwner sets the "owner" edge to the User entity.
func (rc *RateCreate) SetOwner(u *User) *RateCreate {
	return rc.SetOwnerID(u.ID)
}

// SetSubjectID sets the "subject" edge to the Product entity by ID.
func (rc *RateCreate) SetSubjectID(id int) *RateCreate {
	rc.mutation.SetSubjectID(id)
	return rc
}

// SetNillableSubjectID sets the "subject" edge to the Product entity by ID if the given value is not nil.
func (rc *RateCreate) SetNillableSubjectID(id *int) *RateCreate {
	if id != nil {
		rc = rc.SetSubjectID(*id)
	}
	return rc
}

// SetSubject sets the "subject" edge to the Product entity.
func (rc *RateCreate) SetSubject(p *Product) *RateCreate {
	return rc.SetSubjectID(p.ID)
}

// Mutation returns the RateMutation object of the builder.
func (rc *RateCreate) Mutation() *RateMutation {
	return rc.mutation
}

// Save creates the Rate in the database.
func (rc *RateCreate) Save(ctx context.Context) (*Rate, error) {
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RateCreate) SaveX(ctx context.Context) *Rate {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RateCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RateCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RateCreate) check() error {
	if _, ok := rc.mutation.Rate(); !ok {
		return &ValidationError{Name: "rate", err: errors.New(`ent: missing required field "Rate.rate"`)}
	}
	return nil
}

func (rc *RateCreate) sqlSave(ctx context.Context) (*Rate, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RateCreate) createSpec() (*Rate, *sqlgraph.CreateSpec) {
	var (
		_node = &Rate{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(rate.Table, sqlgraph.NewFieldSpec(rate.FieldID, field.TypeInt))
	)
	if value, ok := rc.mutation.Rate(); ok {
		_spec.SetField(rate.FieldRate, field.TypeBool, value)
		_node.Rate = value
	}
	if nodes := rc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rate.OwnerTable,
			Columns: []string{rate.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_rates = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.SubjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rate.SubjectTable,
			Columns: []string{rate.SubjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.product_rates = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RateCreateBulk is the builder for creating many Rate entities in bulk.
type RateCreateBulk struct {
	config
	builders []*RateCreate
}

// Save creates the Rate entities in the database.
func (rcb *RateCreateBulk) Save(ctx context.Context) ([]*Rate, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Rate, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RateMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RateCreateBulk) SaveX(ctx context.Context) []*Rate {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RateCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RateCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
