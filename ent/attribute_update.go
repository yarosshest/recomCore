// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recomCore/ent/attribute"
	"recomCore/ent/predicate"
	"recomCore/ent/product"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AttributeUpdate is the builder for updating Attribute entities.
type AttributeUpdate struct {
	config
	hooks    []Hook
	mutation *AttributeMutation
}

// Where appends a list predicates to the AttributeUpdate builder.
func (au *AttributeUpdate) Where(ps ...predicate.Attribute) *AttributeUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *AttributeUpdate) SetName(s string) *AttributeUpdate {
	au.mutation.SetName(s)
	return au
}

// SetValueType sets the "value_type" field.
func (au *AttributeUpdate) SetValueType(s string) *AttributeUpdate {
	au.mutation.SetValueType(s)
	return au
}

// SetNillableValueType sets the "value_type" field if the given value is not nil.
func (au *AttributeUpdate) SetNillableValueType(s *string) *AttributeUpdate {
	if s != nil {
		au.SetValueType(*s)
	}
	return au
}

// ClearValueType clears the value of the "value_type" field.
func (au *AttributeUpdate) ClearValueType() *AttributeUpdate {
	au.mutation.ClearValueType()
	return au
}

// SetValue sets the "value" field.
func (au *AttributeUpdate) SetValue(s string) *AttributeUpdate {
	au.mutation.SetValue(s)
	return au
}

// SetValueDescription sets the "value_description" field.
func (au *AttributeUpdate) SetValueDescription(s string) *AttributeUpdate {
	au.mutation.SetValueDescription(s)
	return au
}

// SetNillableValueDescription sets the "value_description" field if the given value is not nil.
func (au *AttributeUpdate) SetNillableValueDescription(s *string) *AttributeUpdate {
	if s != nil {
		au.SetValueDescription(*s)
	}
	return au
}

// ClearValueDescription clears the value of the "value_description" field.
func (au *AttributeUpdate) ClearValueDescription() *AttributeUpdate {
	au.mutation.ClearValueDescription()
	return au
}

// SetOwnerID sets the "owner" edge to the Product entity by ID.
func (au *AttributeUpdate) SetOwnerID(id int) *AttributeUpdate {
	au.mutation.SetOwnerID(id)
	return au
}

// SetNillableOwnerID sets the "owner" edge to the Product entity by ID if the given value is not nil.
func (au *AttributeUpdate) SetNillableOwnerID(id *int) *AttributeUpdate {
	if id != nil {
		au = au.SetOwnerID(*id)
	}
	return au
}

// SetOwner sets the "owner" edge to the Product entity.
func (au *AttributeUpdate) SetOwner(p *Product) *AttributeUpdate {
	return au.SetOwnerID(p.ID)
}

// Mutation returns the AttributeMutation object of the builder.
func (au *AttributeUpdate) Mutation() *AttributeMutation {
	return au.mutation
}

// ClearOwner clears the "owner" edge to the Product entity.
func (au *AttributeUpdate) ClearOwner() *AttributeUpdate {
	au.mutation.ClearOwner()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AttributeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AttributeUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AttributeUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AttributeUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AttributeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(attribute.Table, attribute.Columns, sqlgraph.NewFieldSpec(attribute.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(attribute.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.ValueType(); ok {
		_spec.SetField(attribute.FieldValueType, field.TypeString, value)
	}
	if au.mutation.ValueTypeCleared() {
		_spec.ClearField(attribute.FieldValueType, field.TypeString)
	}
	if value, ok := au.mutation.Value(); ok {
		_spec.SetField(attribute.FieldValue, field.TypeString, value)
	}
	if value, ok := au.mutation.ValueDescription(); ok {
		_spec.SetField(attribute.FieldValueDescription, field.TypeString, value)
	}
	if au.mutation.ValueDescriptionCleared() {
		_spec.ClearField(attribute.FieldValueDescription, field.TypeString)
	}
	if au.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attribute.OwnerTable,
			Columns: []string{attribute.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attribute.OwnerTable,
			Columns: []string{attribute.OwnerColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attribute.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AttributeUpdateOne is the builder for updating a single Attribute entity.
type AttributeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AttributeMutation
}

// SetName sets the "name" field.
func (auo *AttributeUpdateOne) SetName(s string) *AttributeUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetValueType sets the "value_type" field.
func (auo *AttributeUpdateOne) SetValueType(s string) *AttributeUpdateOne {
	auo.mutation.SetValueType(s)
	return auo
}

// SetNillableValueType sets the "value_type" field if the given value is not nil.
func (auo *AttributeUpdateOne) SetNillableValueType(s *string) *AttributeUpdateOne {
	if s != nil {
		auo.SetValueType(*s)
	}
	return auo
}

// ClearValueType clears the value of the "value_type" field.
func (auo *AttributeUpdateOne) ClearValueType() *AttributeUpdateOne {
	auo.mutation.ClearValueType()
	return auo
}

// SetValue sets the "value" field.
func (auo *AttributeUpdateOne) SetValue(s string) *AttributeUpdateOne {
	auo.mutation.SetValue(s)
	return auo
}

// SetValueDescription sets the "value_description" field.
func (auo *AttributeUpdateOne) SetValueDescription(s string) *AttributeUpdateOne {
	auo.mutation.SetValueDescription(s)
	return auo
}

// SetNillableValueDescription sets the "value_description" field if the given value is not nil.
func (auo *AttributeUpdateOne) SetNillableValueDescription(s *string) *AttributeUpdateOne {
	if s != nil {
		auo.SetValueDescription(*s)
	}
	return auo
}

// ClearValueDescription clears the value of the "value_description" field.
func (auo *AttributeUpdateOne) ClearValueDescription() *AttributeUpdateOne {
	auo.mutation.ClearValueDescription()
	return auo
}

// SetOwnerID sets the "owner" edge to the Product entity by ID.
func (auo *AttributeUpdateOne) SetOwnerID(id int) *AttributeUpdateOne {
	auo.mutation.SetOwnerID(id)
	return auo
}

// SetNillableOwnerID sets the "owner" edge to the Product entity by ID if the given value is not nil.
func (auo *AttributeUpdateOne) SetNillableOwnerID(id *int) *AttributeUpdateOne {
	if id != nil {
		auo = auo.SetOwnerID(*id)
	}
	return auo
}

// SetOwner sets the "owner" edge to the Product entity.
func (auo *AttributeUpdateOne) SetOwner(p *Product) *AttributeUpdateOne {
	return auo.SetOwnerID(p.ID)
}

// Mutation returns the AttributeMutation object of the builder.
func (auo *AttributeUpdateOne) Mutation() *AttributeMutation {
	return auo.mutation
}

// ClearOwner clears the "owner" edge to the Product entity.
func (auo *AttributeUpdateOne) ClearOwner() *AttributeUpdateOne {
	auo.mutation.ClearOwner()
	return auo
}

// Where appends a list predicates to the AttributeUpdate builder.
func (auo *AttributeUpdateOne) Where(ps ...predicate.Attribute) *AttributeUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AttributeUpdateOne) Select(field string, fields ...string) *AttributeUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Attribute entity.
func (auo *AttributeUpdateOne) Save(ctx context.Context) (*Attribute, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AttributeUpdateOne) SaveX(ctx context.Context) *Attribute {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AttributeUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AttributeUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AttributeUpdateOne) sqlSave(ctx context.Context) (_node *Attribute, err error) {
	_spec := sqlgraph.NewUpdateSpec(attribute.Table, attribute.Columns, sqlgraph.NewFieldSpec(attribute.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Attribute.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attribute.FieldID)
		for _, f := range fields {
			if !attribute.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != attribute.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(attribute.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.ValueType(); ok {
		_spec.SetField(attribute.FieldValueType, field.TypeString, value)
	}
	if auo.mutation.ValueTypeCleared() {
		_spec.ClearField(attribute.FieldValueType, field.TypeString)
	}
	if value, ok := auo.mutation.Value(); ok {
		_spec.SetField(attribute.FieldValue, field.TypeString, value)
	}
	if value, ok := auo.mutation.ValueDescription(); ok {
		_spec.SetField(attribute.FieldValueDescription, field.TypeString, value)
	}
	if auo.mutation.ValueDescriptionCleared() {
		_spec.ClearField(attribute.FieldValueDescription, field.TypeString)
	}
	if auo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attribute.OwnerTable,
			Columns: []string{attribute.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attribute.OwnerTable,
			Columns: []string{attribute.OwnerColumn},
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
	_node = &Attribute{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attribute.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
