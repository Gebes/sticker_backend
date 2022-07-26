// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"gebes.io/sticker_backend/pkg/ent/sticker"
	"gebes.io/sticker_backend/pkg/ent/user"
)

// Sticker is the model entity for the Sticker schema.
type Sticker struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// LocationDescription holds the value of the "location_description" field.
	LocationDescription string `json:"location_description"`
	// Latitude holds the value of the "latitude" field.
	Latitude float64 `json:"latitude"`
	// Longitude holds the value of the "longitude" field.
	Longitude float64 `json:"longitude"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StickerQuery when eager-loading is set.
	Edges         StickerEdges `json:"edges"`
	sticker_owner *string
}

// StickerEdges holds the relations/edges for other nodes in the graph.
type StickerEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StickerEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Sticker) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case sticker.FieldLatitude, sticker.FieldLongitude:
			values[i] = new(sql.NullFloat64)
		case sticker.FieldID:
			values[i] = new(sql.NullInt64)
		case sticker.FieldLocationDescription:
			values[i] = new(sql.NullString)
		case sticker.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case sticker.ForeignKeys[0]: // sticker_owner
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Sticker", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Sticker fields.
func (s *Sticker) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sticker.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case sticker.FieldLocationDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location_description", values[i])
			} else if value.Valid {
				s.LocationDescription = value.String
			}
		case sticker.FieldLatitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field latitude", values[i])
			} else if value.Valid {
				s.Latitude = value.Float64
			}
		case sticker.FieldLongitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field longitude", values[i])
			} else if value.Valid {
				s.Longitude = value.Float64
			}
		case sticker.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case sticker.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sticker_owner", values[i])
			} else if value.Valid {
				s.sticker_owner = new(string)
				*s.sticker_owner = value.String
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Sticker entity.
func (s *Sticker) QueryOwner() *UserQuery {
	return (&StickerClient{config: s.config}).QueryOwner(s)
}

// Update returns a builder for updating this Sticker.
// Note that you need to call Sticker.Unwrap() before calling this method if this Sticker
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Sticker) Update() *StickerUpdateOne {
	return (&StickerClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Sticker entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Sticker) Unwrap() *Sticker {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Sticker is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Sticker) String() string {
	var builder strings.Builder
	builder.WriteString("Sticker(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", location_description=")
	builder.WriteString(s.LocationDescription)
	builder.WriteString(", latitude=")
	builder.WriteString(fmt.Sprintf("%v", s.Latitude))
	builder.WriteString(", longitude=")
	builder.WriteString(fmt.Sprintf("%v", s.Longitude))
	builder.WriteString(", created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Stickers is a parsable slice of Sticker.
type Stickers []*Sticker

func (s Stickers) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
