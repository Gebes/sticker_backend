// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gebes.io/sticker_backend/pkg/ent/sticker"
	"gebes.io/sticker_backend/pkg/ent/user"
)

// StickerCreate is the builder for creating a Sticker entity.
type StickerCreate struct {
	config
	mutation *StickerMutation
	hooks    []Hook
}

// SetLocationDescription sets the "location_description" field.
func (sc *StickerCreate) SetLocationDescription(s string) *StickerCreate {
	sc.mutation.SetLocationDescription(s)
	return sc
}

// SetLatitude sets the "latitude" field.
func (sc *StickerCreate) SetLatitude(f float64) *StickerCreate {
	sc.mutation.SetLatitude(f)
	return sc
}

// SetLongitude sets the "longitude" field.
func (sc *StickerCreate) SetLongitude(f float64) *StickerCreate {
	sc.mutation.SetLongitude(f)
	return sc
}

// SetEdition sets the "edition" field.
func (sc *StickerCreate) SetEdition(s sticker.Edition) *StickerCreate {
	sc.mutation.SetEdition(s)
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *StickerCreate) SetCreatedAt(t time.Time) *StickerCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *StickerCreate) SetNillableCreatedAt(t *time.Time) *StickerCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (sc *StickerCreate) SetOwnerID(id string) *StickerCreate {
	sc.mutation.SetOwnerID(id)
	return sc
}

// SetOwner sets the "owner" edge to the User entity.
func (sc *StickerCreate) SetOwner(u *User) *StickerCreate {
	return sc.SetOwnerID(u.ID)
}

// Mutation returns the StickerMutation object of the builder.
func (sc *StickerCreate) Mutation() *StickerMutation {
	return sc.mutation
}

// Save creates the Sticker in the database.
func (sc *StickerCreate) Save(ctx context.Context) (*Sticker, error) {
	var (
		err  error
		node *Sticker
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StickerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StickerCreate) SaveX(ctx context.Context) *Sticker {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StickerCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StickerCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StickerCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := sticker.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StickerCreate) check() error {
	if _, ok := sc.mutation.LocationDescription(); !ok {
		return &ValidationError{Name: "location_description", err: errors.New(`ent: missing required field "Sticker.location_description"`)}
	}
	if v, ok := sc.mutation.LocationDescription(); ok {
		if err := sticker.LocationDescriptionValidator(v); err != nil {
			return &ValidationError{Name: "location_description", err: fmt.Errorf(`ent: validator failed for field "Sticker.location_description": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Latitude(); !ok {
		return &ValidationError{Name: "latitude", err: errors.New(`ent: missing required field "Sticker.latitude"`)}
	}
	if _, ok := sc.mutation.Longitude(); !ok {
		return &ValidationError{Name: "longitude", err: errors.New(`ent: missing required field "Sticker.longitude"`)}
	}
	if _, ok := sc.mutation.Edition(); !ok {
		return &ValidationError{Name: "edition", err: errors.New(`ent: missing required field "Sticker.edition"`)}
	}
	if v, ok := sc.mutation.Edition(); ok {
		if err := sticker.EditionValidator(v); err != nil {
			return &ValidationError{Name: "edition", err: fmt.Errorf(`ent: validator failed for field "Sticker.edition": %w`, err)}
		}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Sticker.created_at"`)}
	}
	if _, ok := sc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Sticker.owner"`)}
	}
	return nil
}

func (sc *StickerCreate) sqlSave(ctx context.Context) (*Sticker, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *StickerCreate) createSpec() (*Sticker, *sqlgraph.CreateSpec) {
	var (
		_node = &Sticker{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sticker.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sticker.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.LocationDescription(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sticker.FieldLocationDescription,
		})
		_node.LocationDescription = value
	}
	if value, ok := sc.mutation.Latitude(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sticker.FieldLatitude,
		})
		_node.Latitude = value
	}
	if value, ok := sc.mutation.Longitude(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: sticker.FieldLongitude,
		})
		_node.Longitude = value
	}
	if value, ok := sc.mutation.Edition(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: sticker.FieldEdition,
		})
		_node.Edition = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sticker.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := sc.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_node.sticker_owner = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StickerCreateBulk is the builder for creating many Sticker entities in bulk.
type StickerCreateBulk struct {
	config
	builders []*StickerCreate
}

// Save creates the Sticker entities in the database.
func (scb *StickerCreateBulk) Save(ctx context.Context) ([]*Sticker, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Sticker, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StickerMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StickerCreateBulk) SaveX(ctx context.Context) []*Sticker {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StickerCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StickerCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
