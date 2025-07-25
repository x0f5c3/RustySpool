// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"spoolman-go/ent/spool"
	"spoolman-go/ent/spoolfield"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// SpoolField is the model entity for the SpoolField schema.
type SpoolField struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SpoolID holds the value of the "spool_id" field.
	SpoolID int `json:"spool_id,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SpoolFieldQuery when eager-loading is set.
	Edges        SpoolFieldEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SpoolFieldEdges holds the relations/edges for other nodes in the graph.
type SpoolFieldEdges struct {
	// Spool holds the value of the spool edge.
	Spool *Spool `json:"spool,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SpoolOrErr returns the Spool value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SpoolFieldEdges) SpoolOrErr() (*Spool, error) {
	if e.Spool != nil {
		return e.Spool, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: spool.Label}
	}
	return nil, &NotLoadedError{edge: "spool"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SpoolField) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case spoolfield.FieldID, spoolfield.FieldSpoolID:
			values[i] = new(sql.NullInt64)
		case spoolfield.FieldKey, spoolfield.FieldValue:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SpoolField fields.
func (sf *SpoolField) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case spoolfield.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sf.ID = int(value.Int64)
		case spoolfield.FieldSpoolID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field spool_id", values[i])
			} else if value.Valid {
				sf.SpoolID = int(value.Int64)
			}
		case spoolfield.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				sf.Key = value.String
			}
		case spoolfield.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				sf.Value = value.String
			}
		default:
			sf.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the SpoolField.
// This includes values selected through modifiers, order, etc.
func (sf *SpoolField) GetValue(name string) (ent.Value, error) {
	return sf.selectValues.Get(name)
}

// QuerySpool queries the "spool" edge of the SpoolField entity.
func (sf *SpoolField) QuerySpool() *SpoolQuery {
	return NewSpoolFieldClient(sf.config).QuerySpool(sf)
}

// Update returns a builder for updating this SpoolField.
// Note that you need to call SpoolField.Unwrap() before calling this method if this SpoolField
// was returned from a transaction, and the transaction was committed or rolled back.
func (sf *SpoolField) Update() *SpoolFieldUpdateOne {
	return NewSpoolFieldClient(sf.config).UpdateOne(sf)
}

// Unwrap unwraps the SpoolField entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sf *SpoolField) Unwrap() *SpoolField {
	_tx, ok := sf.config.driver.(*txDriver)
	if !ok {
		panic("ent: SpoolField is not a transactional entity")
	}
	sf.config.driver = _tx.drv
	return sf
}

// String implements the fmt.Stringer.
func (sf *SpoolField) String() string {
	var builder strings.Builder
	builder.WriteString("SpoolField(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sf.ID))
	builder.WriteString("spool_id=")
	builder.WriteString(fmt.Sprintf("%v", sf.SpoolID))
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(sf.Key)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(sf.Value)
	builder.WriteByte(')')
	return builder.String()
}

// SpoolFields is a parsable slice of SpoolField.
type SpoolFields []*SpoolField
