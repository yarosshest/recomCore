// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recomCore/ent/predicate"
	"recomCore/ent/product"
	"recomCore/ent/vector"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VectorUpdate is the builder for updating Vector entities.
type VectorUpdate struct {
	config
	hooks    []Hook
	mutation *VectorMutation
}

// Where appends a list predicates to the VectorUpdate builder.
func (vu *VectorUpdate) Where(ps ...predicate.Vector) *VectorUpdate {
	vu.mutation.Where(ps...)
	return vu
}

// SetVector sets the "vector" field.
func (vu *VectorUpdate) SetVector(b []byte) *VectorUpdate {
	vu.mutation.SetVector(b)
	return vu
}

// SetOwnerID sets the "owner" edge to the Product entity by ID.
func (vu *VectorUpdate) SetOwnerID(id int) *VectorUpdate {
	vu.mutation.SetOwnerID(id)
	return vu
}

// SetNillableOwnerID sets the "owner" edge to the Product entity by ID if the given value is not nil.
func (vu *VectorUpdate) SetNillableOwnerID(id *int) *VectorUpdate {
	if id != nil {
		vu = vu.SetOwnerID(*id)
	}
	return vu
}

// SetOwner sets the "owner" edge to the Product entity.
func (vu *VectorUpdate) SetOwner(p *Product) *VectorUpdate {
	return vu.SetOwnerID(p.ID)
}

// Mutation returns the VectorMutation object of the builder.
func (vu *VectorUpdate) Mutation() *VectorMutation {
	return vu.mutation
}

// ClearOwner clears the "owner" edge to the Product entity.
func (vu *VectorUpdate) ClearOwner() *VectorUpdate {
	vu.mutation.ClearOwner()
	return vu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *VectorUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, vu.sqlSave, vu.mutation, vu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vu *VectorUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VectorUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VectorUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vu *VectorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(vector.Table, vector.Columns, sqlgraph.NewFieldSpec(vector.FieldID, field.TypeInt))
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.Vector(); ok {
		_spec.SetField(vector.FieldVector, field.TypeBytes, value)
	}
	if vu.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vector.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vu.mutation.done = true
	return n, nil
}

// VectorUpdateOne is the builder for updating a single Vector entity.
type VectorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VectorMutation
}

// SetVector sets the "vector" field.
func (vuo *VectorUpdateOne) SetVector(b []byte) *VectorUpdateOne {
	vuo.mutation.SetVector(b)
	return vuo
}

// SetOwnerID sets the "owner" edge to the Product entity by ID.
func (vuo *VectorUpdateOne) SetOwnerID(id int) *VectorUpdateOne {
	vuo.mutation.SetOwnerID(id)
	return vuo
}

// SetNillableOwnerID sets the "owner" edge to the Product entity by ID if the given value is not nil.
func (vuo *VectorUpdateOne) SetNillableOwnerID(id *int) *VectorUpdateOne {
	if id != nil {
		vuo = vuo.SetOwnerID(*id)
	}
	return vuo
}

// SetOwner sets the "owner" edge to the Product entity.
func (vuo *VectorUpdateOne) SetOwner(p *Product) *VectorUpdateOne {
	return vuo.SetOwnerID(p.ID)
}

// Mutation returns the VectorMutation object of the builder.
func (vuo *VectorUpdateOne) Mutation() *VectorMutation {
	return vuo.mutation
}

// ClearOwner clears the "owner" edge to the Product entity.
func (vuo *VectorUpdateOne) ClearOwner() *VectorUpdateOne {
	vuo.mutation.ClearOwner()
	return vuo
}

// Where appends a list predicates to the VectorUpdate builder.
func (vuo *VectorUpdateOne) Where(ps ...predicate.Vector) *VectorUpdateOne {
	vuo.mutation.Where(ps...)
	return vuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vuo *VectorUpdateOne) Select(field string, fields ...string) *VectorUpdateOne {
	vuo.fields = append([]string{field}, fields...)
	return vuo
}

// Save executes the query and returns the updated Vector entity.
func (vuo *VectorUpdateOne) Save(ctx context.Context) (*Vector, error) {
	return withHooks(ctx, vuo.sqlSave, vuo.mutation, vuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *VectorUpdateOne) SaveX(ctx context.Context) *Vector {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VectorUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VectorUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vuo *VectorUpdateOne) sqlSave(ctx context.Context) (_node *Vector, err error) {
	_spec := sqlgraph.NewUpdateSpec(vector.Table, vector.Columns, sqlgraph.NewFieldSpec(vector.FieldID, field.TypeInt))
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Vector.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vector.FieldID)
		for _, f := range fields {
			if !vector.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != vector.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vuo.mutation.Vector(); ok {
		_spec.SetField(vector.FieldVector, field.TypeBytes, value)
	}
	if vuo.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Vector{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vector.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vuo.mutation.done = true
	return _node, nil
}
