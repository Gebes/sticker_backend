// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gebes.io/sticker_backend/pkg/ent/predicate"
	"gebes.io/sticker_backend/pkg/ent/sticker"
	"gebes.io/sticker_backend/pkg/ent/user"
)

// StickerUpdate is the builder for updating Sticker entities.
type StickerUpdate struct {
	config
	hooks    []Hook
	mutation *StickerMutation
}

// Where appends a list predicates to the StickerUpdate builder.
func (su *StickerUpdate) Where(ps ...predicate.Sticker) *StickerUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetLocationDescription sets the "location_description" field.
func (su *StickerUpdate) SetLocationDescription(s string) *StickerUpdate {
	su.mutation.SetLocationDescription(s)
	return su
}

// SetLatitude sets the "latitude" field.
func (su *StickerUpdate) SetLatitude(f float64) *StickerUpdate {
	su.mutation.ResetLatitude()
	su.mutation.SetLatitude(f)
	return su
}

// AddLatitude adds f to the "latitude" field.
func (su *StickerUpdate) AddLatitude(f float64) *StickerUpdate {
	su.mutation.AddLatitude(f)
	return su
}

// SetLongitude sets the "longitude" field.
func (su *StickerUpdate) SetLongitude(f float64) *StickerUpdate {
	su.mutation.ResetLongitude()
	su.mutation.SetLongitude(f)
	return su
}

// AddLongitude adds f to the "longitude" field.
func (su *StickerUpdate) AddLongitude(f float64) *StickerUpdate {
	su.mutation.AddLongitude(f)
	return su
}

// SetEdition sets the "edition" field.
func (su *StickerUpdate) SetEdition(s sticker.Edition) *StickerUpdate {
	su.mutation.SetEdition(s)
	return su
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (su *StickerUpdate) SetOwnerID(id string) *StickerUpdate {
	su.mutation.SetOwnerID(id)
	return su
}

// SetOwner sets the "owner" edge to the User entity.
func (su *StickerUpdate) SetOwner(u *User) *StickerUpdate {
	return su.SetOwnerID(u.ID)
}

// Mutation returns the StickerMutation object of the builder.
func (su *StickerUpdate) Mutation() *StickerMutation {
	return su.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (su *StickerUpdate) ClearOwner() *StickerUpdate {
	su.mutation.ClearOwner()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StickerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StickerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *StickerUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StickerUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StickerUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *StickerUpdate) check() error {
	if v, ok := su.mutation.LocationDescription(); ok {
		if err := sticker.LocationDescriptionValidator(v); err != nil {
			return &ValidationError{Name: "location_description", err: fmt.Errorf(`ent: validator failed for field "Sticker.location_description": %w`, err)}
		}
	}
	if v, ok := su.mutation.Edition(); ok {
		if err := sticker.EditionValidator(v); err != nil {
			return &ValidationError{Name: "edition", err: fmt.Errorf(`ent: validator failed for field "Sticker.edition": %w`, err)}
		}
	}
	if _, ok := su.mutation.OwnerID(); su.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Sticker.owner"`)
	}
	return nil
}

func (su *StickerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sticker.Table,
			Columns: sticker.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sticker.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.LocationDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sticker.FieldLocationDescription,
		})
	}
	if value, ok := su.mutation.Latitude(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sticker.FieldLatitude,
		})
	}
	if value, ok := su.mutation.AddedLatitude(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sticker.FieldLatitude,
		})
	}
	if value, ok := su.mutation.Longitude(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sticker.FieldLongitude,
		})
	}
	if value, ok := su.mutation.AddedLongitude(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sticker.FieldLongitude,
		})
	}
	if value, ok := su.mutation.Edition(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: sticker.FieldEdition,
		})
	}
	if su.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sticker.OwnerTable,
			Columns: []string{sticker.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sticker.OwnerTable,
			Columns: []string{sticker.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sticker.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// StickerUpdateOne is the builder for updating a single Sticker entity.
type StickerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StickerMutation
}

// SetLocationDescription sets the "location_description" field.
func (suo *StickerUpdateOne) SetLocationDescription(s string) *StickerUpdateOne {
	suo.mutation.SetLocationDescription(s)
	return suo
}

// SetLatitude sets the "latitude" field.
func (suo *StickerUpdateOne) SetLatitude(f float64) *StickerUpdateOne {
	suo.mutation.ResetLatitude()
	suo.mutation.SetLatitude(f)
	return suo
}

// AddLatitude adds f to the "latitude" field.
func (suo *StickerUpdateOne) AddLatitude(f float64) *StickerUpdateOne {
	suo.mutation.AddLatitude(f)
	return suo
}

// SetLongitude sets the "longitude" field.
func (suo *StickerUpdateOne) SetLongitude(f float64) *StickerUpdateOne {
	suo.mutation.ResetLongitude()
	suo.mutation.SetLongitude(f)
	return suo
}

// AddLongitude adds f to the "longitude" field.
func (suo *StickerUpdateOne) AddLongitude(f float64) *StickerUpdateOne {
	suo.mutation.AddLongitude(f)
	return suo
}

// SetEdition sets the "edition" field.
func (suo *StickerUpdateOne) SetEdition(s sticker.Edition) *StickerUpdateOne {
	suo.mutation.SetEdition(s)
	return suo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (suo *StickerUpdateOne) SetOwnerID(id string) *StickerUpdateOne {
	suo.mutation.SetOwnerID(id)
	return suo
}

// SetOwner sets the "owner" edge to the User entity.
func (suo *StickerUpdateOne) SetOwner(u *User) *StickerUpdateOne {
	return suo.SetOwnerID(u.ID)
}

// Mutation returns the StickerMutation object of the builder.
func (suo *StickerUpdateOne) Mutation() *StickerMutation {
	return suo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (suo *StickerUpdateOne) ClearOwner() *StickerUpdateOne {
	suo.mutation.ClearOwner()
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StickerUpdateOne) Select(field string, fields ...string) *StickerUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Sticker entity.
func (suo *StickerUpdateOne) Save(ctx context.Context) (*Sticker, error) {
	var (
		err  error
		node *Sticker
	)
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StickerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StickerUpdateOne) SaveX(ctx context.Context) *Sticker {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StickerUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StickerUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *StickerUpdateOne) check() error {
	if v, ok := suo.mutation.LocationDescription(); ok {
		if err := sticker.LocationDescriptionValidator(v); err != nil {
			return &ValidationError{Name: "location_description", err: fmt.Errorf(`ent: validator failed for field "Sticker.location_description": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Edition(); ok {
		if err := sticker.EditionValidator(v); err != nil {
			return &ValidationError{Name: "edition", err: fmt.Errorf(`ent: validator failed for field "Sticker.edition": %w`, err)}
		}
	}
	if _, ok := suo.mutation.OwnerID(); suo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Sticker.owner"`)
	}
	return nil
}

func (suo *StickerUpdateOne) sqlSave(ctx context.Context) (_node *Sticker, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sticker.Table,
			Columns: sticker.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sticker.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Sticker.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sticker.FieldID)
		for _, f := range fields {
			if !sticker.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sticker.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.LocationDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sticker.FieldLocationDescription,
		})
	}
	if value, ok := suo.mutation.Latitude(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sticker.FieldLatitude,
		})
	}
	if value, ok := suo.mutation.AddedLatitude(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sticker.FieldLatitude,
		})
	}
	if value, ok := suo.mutation.Longitude(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sticker.FieldLongitude,
		})
	}
	if value, ok := suo.mutation.AddedLongitude(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sticker.FieldLongitude,
		})
	}
	if value, ok := suo.mutation.Edition(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: sticker.FieldEdition,
		})
	}
	if suo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sticker.OwnerTable,
			Columns: []string{sticker.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sticker.OwnerTable,
			Columns: []string{sticker.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Sticker{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sticker.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
