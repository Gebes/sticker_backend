// Code generated by entc, DO NOT EDIT.

package sticker

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gebes.io/sticker_backend/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// LocationDescription applies equality check predicate on the "location_description" field. It's identical to LocationDescriptionEQ.
func LocationDescription(v string) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocationDescription), v))
	})
}

// Latitude applies equality check predicate on the "latitude" field. It's identical to LatitudeEQ.
func Latitude(v float64) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLatitude), v))
	})
}

// Longitude applies equality check predicate on the "longitude" field. It's identical to LongitudeEQ.
func Longitude(v float64) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLongitude), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// LocationDescriptionEQ applies the EQ predicate on the "location_description" field.
func LocationDescriptionEQ(v string) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocationDescription), v))
	})
}

// LocationDescriptionNEQ applies the NEQ predicate on the "location_description" field.
func LocationDescriptionNEQ(v string) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLocationDescription), v))
	})
}

// LocationDescriptionIn applies the In predicate on the "location_description" field.
func LocationDescriptionIn(vs ...string) predicate.Sticker {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sticker(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLocationDescription), v...))
	})
}

// LocationDescriptionNotIn applies the NotIn predicate on the "location_description" field.
func LocationDescriptionNotIn(vs ...string) predicate.Sticker {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sticker(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLocationDescription), v...))
	})
}

// LocationDescriptionGT applies the GT predicate on the "location_description" field.
func LocationDescriptionGT(v string) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLocationDescription), v))
	})
}

// LocationDescriptionGTE applies the GTE predicate on the "location_description" field.
func LocationDescriptionGTE(v string) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLocationDescription), v))
	})
}

// LocationDescriptionLT applies the LT predicate on the "location_description" field.
func LocationDescriptionLT(v string) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLocationDescription), v))
	})
}

// LocationDescriptionLTE applies the LTE predicate on the "location_description" field.
func LocationDescriptionLTE(v string) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLocationDescription), v))
	})
}

// LocationDescriptionContains applies the Contains predicate on the "location_description" field.
func LocationDescriptionContains(v string) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLocationDescription), v))
	})
}

// LocationDescriptionHasPrefix applies the HasPrefix predicate on the "location_description" field.
func LocationDescriptionHasPrefix(v string) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLocationDescription), v))
	})
}

// LocationDescriptionHasSuffix applies the HasSuffix predicate on the "location_description" field.
func LocationDescriptionHasSuffix(v string) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLocationDescription), v))
	})
}

// LocationDescriptionEqualFold applies the EqualFold predicate on the "location_description" field.
func LocationDescriptionEqualFold(v string) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLocationDescription), v))
	})
}

// LocationDescriptionContainsFold applies the ContainsFold predicate on the "location_description" field.
func LocationDescriptionContainsFold(v string) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLocationDescription), v))
	})
}

// LatitudeEQ applies the EQ predicate on the "latitude" field.
func LatitudeEQ(v float64) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLatitude), v))
	})
}

// LatitudeNEQ applies the NEQ predicate on the "latitude" field.
func LatitudeNEQ(v float64) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLatitude), v))
	})
}

// LatitudeIn applies the In predicate on the "latitude" field.
func LatitudeIn(vs ...float64) predicate.Sticker {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sticker(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLatitude), v...))
	})
}

// LatitudeNotIn applies the NotIn predicate on the "latitude" field.
func LatitudeNotIn(vs ...float64) predicate.Sticker {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sticker(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLatitude), v...))
	})
}

// LatitudeGT applies the GT predicate on the "latitude" field.
func LatitudeGT(v float64) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLatitude), v))
	})
}

// LatitudeGTE applies the GTE predicate on the "latitude" field.
func LatitudeGTE(v float64) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLatitude), v))
	})
}

// LatitudeLT applies the LT predicate on the "latitude" field.
func LatitudeLT(v float64) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLatitude), v))
	})
}

// LatitudeLTE applies the LTE predicate on the "latitude" field.
func LatitudeLTE(v float64) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLatitude), v))
	})
}

// LongitudeEQ applies the EQ predicate on the "longitude" field.
func LongitudeEQ(v float64) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLongitude), v))
	})
}

// LongitudeNEQ applies the NEQ predicate on the "longitude" field.
func LongitudeNEQ(v float64) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLongitude), v))
	})
}

// LongitudeIn applies the In predicate on the "longitude" field.
func LongitudeIn(vs ...float64) predicate.Sticker {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sticker(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLongitude), v...))
	})
}

// LongitudeNotIn applies the NotIn predicate on the "longitude" field.
func LongitudeNotIn(vs ...float64) predicate.Sticker {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sticker(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLongitude), v...))
	})
}

// LongitudeGT applies the GT predicate on the "longitude" field.
func LongitudeGT(v float64) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLongitude), v))
	})
}

// LongitudeGTE applies the GTE predicate on the "longitude" field.
func LongitudeGTE(v float64) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLongitude), v))
	})
}

// LongitudeLT applies the LT predicate on the "longitude" field.
func LongitudeLT(v float64) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLongitude), v))
	})
}

// LongitudeLTE applies the LTE predicate on the "longitude" field.
func LongitudeLTE(v float64) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLongitude), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Sticker {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sticker(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Sticker {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Sticker(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Sticker) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Sticker) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Sticker) predicate.Sticker {
	return predicate.Sticker(func(s *sql.Selector) {
		p(s.Not())
	})
}