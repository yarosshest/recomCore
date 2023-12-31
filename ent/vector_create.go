// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recomCore/ent/product"
	"recomCore/ent/vector"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VectorCreate is the builder for creating a Vector entity.
type VectorCreate struct {
	config
	mutation *VectorMutation
	hooks    []Hook
}

// SetVector sets the "vector" field.
func (vc *VectorCreate) SetVector(b []byte) *VectorCreate {
	vc.mutation.SetVector(b)
	return vc
}

// SetOwnerID sets the "owner" edge to the Product entity by ID.
func (vc *VectorCreate) SetOwnerID(id int) *VectorCreate {
	vc.mutation.SetOwnerID(id)
	return vc
}

// SetNillableOwnerID sets the "owner" edge to the Product entity by ID if the given value is not nil.
func (vc *VectorCreate) SetNillableOwnerID(id *int) *VectorCreate {
	if id != nil {
		vc = vc.SetOwnerID(*id)
	}
	return vc
}

// SetOwner sets the "owner" edge to the Product entity.
func (vc *VectorCreate) SetOwner(p *Product) *VectorCreate {
	return vc.SetOwnerID(p.ID)
}

// Mutation returns the VectorMutation object of the builder.
func (vc *VectorCreate) Mutation() *VectorMutation {
	return vc.mutation
}

// Save creates the Vector in the database.
func (vc *VectorCreate) Save(ctx context.Context) (*Vector, error) {
	return withHooks(ctx, vc.sqlSave, vc.mutation, vc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VectorCreate) SaveX(ctx context.Context) *Vector {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VectorCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VectorCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VectorCreate) check() error {
	if _, ok := vc.mutation.Vector(); !ok {
		return &ValidationError{Name: "vector", err: errors.New(`ent: missing required field "Vector.vector"`)}
	}
	return nil
}

func (vc *VectorCreate) sqlSave(ctx context.Context) (*Vector, error) {
	if err := vc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	vc.mutation.id = &_node.ID
	vc.mutation.done = true
	return _node, nil
}

func (vc *VectorCreate) createSpec() (*Vector, *sqlgraph.CreateSpec) {
	var (
		_node = &Vector{config: vc.config}
		_spec = sqlgraph.NewCreateSpec(vector.Table, sqlgraph.NewFieldSpec(vector.FieldID, field.TypeInt))
	)
	if value, ok := vc.mutation.Vector(); ok {
		_spec.SetField(vector.FieldVector, field.TypeBytes, value)
		_node.Vector = value
	}
	if nodes := vc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vector.OwnerTable,
			Columns: []string{vector.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.product_vectors = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VectorCreateBulk is the builder for creating many Vector entities in bulk.
type VectorCreateBulk struct {
	config
	builders []*VectorCreate
}

// Save creates the Vector entities in the database.
func (vcb *VectorCreateBulk) Save(ctx context.Context) ([]*Vector, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Vector, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VectorMutation)
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
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VectorCreateBulk) SaveX(ctx context.Context) []*Vector {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VectorCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VectorCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
