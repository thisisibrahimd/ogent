// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"ariga.io/ogent/internal/integration/ogent/ent/alltypes"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// AllTypes is the model entity for the AllTypes schema.
type AllTypes struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// Int holds the value of the "int" field.
	Int int `json:"int,omitempty"`
	// Int8 holds the value of the "int8" field.
	Int8 int8 `json:"int8,omitempty"`
	// Int16 holds the value of the "int16" field.
	Int16 int16 `json:"int16,omitempty"`
	// Int32 holds the value of the "int32" field.
	Int32 int32 `json:"int32,omitempty"`
	// Int64 holds the value of the "int64" field.
	Int64 int64 `json:"int64,omitempty"`
	// Uint holds the value of the "uint" field.
	Uint uint `json:"uint,omitempty"`
	// Uint8 holds the value of the "uint8" field.
	Uint8 uint8 `json:"uint8,omitempty"`
	// Uint16 holds the value of the "uint16" field.
	Uint16 uint16 `json:"uint16,omitempty"`
	// Uint32 holds the value of the "uint32" field.
	Uint32 uint32 `json:"uint32,omitempty"`
	// Uint64 holds the value of the "uint64" field.
	Uint64 uint64 `json:"uint64,omitempty"`
	// Float32 holds the value of the "float32" field.
	Float32 float32 `json:"float32,omitempty"`
	// Float64 holds the value of the "float64" field.
	Float64 float64 `json:"float64,omitempty"`
	// StringType holds the value of the "string_type" field.
	StringType string `json:"string_type,omitempty"`
	// Bool holds the value of the "bool" field.
	Bool bool `json:"bool,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID uuid.UUID `json:"uuid,omitempty"`
	// Time holds the value of the "time" field.
	Time time.Time `json:"time,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// State holds the value of the "state" field.
	State alltypes.State `json:"state,omitempty"`
	// Bytes holds the value of the "bytes" field.
	Bytes []byte `json:"bytes,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AllTypes) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case alltypes.FieldBytes:
			values[i] = new([]byte)
		case alltypes.FieldBool:
			values[i] = new(sql.NullBool)
		case alltypes.FieldFloat32, alltypes.FieldFloat64:
			values[i] = new(sql.NullFloat64)
		case alltypes.FieldID, alltypes.FieldInt, alltypes.FieldInt8, alltypes.FieldInt16, alltypes.FieldInt32, alltypes.FieldInt64, alltypes.FieldUint, alltypes.FieldUint8, alltypes.FieldUint16, alltypes.FieldUint32, alltypes.FieldUint64:
			values[i] = new(sql.NullInt64)
		case alltypes.FieldStringType, alltypes.FieldText, alltypes.FieldState:
			values[i] = new(sql.NullString)
		case alltypes.FieldTime:
			values[i] = new(sql.NullTime)
		case alltypes.FieldUUID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AllTypes", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AllTypes fields.
func (at *AllTypes) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case alltypes.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			at.ID = uint32(value.Int64)
		case alltypes.FieldInt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field int", values[i])
			} else if value.Valid {
				at.Int = int(value.Int64)
			}
		case alltypes.FieldInt8:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field int8", values[i])
			} else if value.Valid {
				at.Int8 = int8(value.Int64)
			}
		case alltypes.FieldInt16:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field int16", values[i])
			} else if value.Valid {
				at.Int16 = int16(value.Int64)
			}
		case alltypes.FieldInt32:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field int32", values[i])
			} else if value.Valid {
				at.Int32 = int32(value.Int64)
			}
		case alltypes.FieldInt64:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field int64", values[i])
			} else if value.Valid {
				at.Int64 = value.Int64
			}
		case alltypes.FieldUint:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field uint", values[i])
			} else if value.Valid {
				at.Uint = uint(value.Int64)
			}
		case alltypes.FieldUint8:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field uint8", values[i])
			} else if value.Valid {
				at.Uint8 = uint8(value.Int64)
			}
		case alltypes.FieldUint16:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field uint16", values[i])
			} else if value.Valid {
				at.Uint16 = uint16(value.Int64)
			}
		case alltypes.FieldUint32:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field uint32", values[i])
			} else if value.Valid {
				at.Uint32 = uint32(value.Int64)
			}
		case alltypes.FieldUint64:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field uint64", values[i])
			} else if value.Valid {
				at.Uint64 = uint64(value.Int64)
			}
		case alltypes.FieldFloat32:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field float32", values[i])
			} else if value.Valid {
				at.Float32 = float32(value.Float64)
			}
		case alltypes.FieldFloat64:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field float64", values[i])
			} else if value.Valid {
				at.Float64 = value.Float64
			}
		case alltypes.FieldStringType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_type", values[i])
			} else if value.Valid {
				at.StringType = value.String
			}
		case alltypes.FieldBool:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field bool", values[i])
			} else if value.Valid {
				at.Bool = value.Bool
			}
		case alltypes.FieldUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value != nil {
				at.UUID = *value
			}
		case alltypes.FieldTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field time", values[i])
			} else if value.Valid {
				at.Time = value.Time
			}
		case alltypes.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				at.Text = value.String
			}
		case alltypes.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				at.State = alltypes.State(value.String)
			}
		case alltypes.FieldBytes:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field bytes", values[i])
			} else if value != nil {
				at.Bytes = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AllTypes.
// Note that you need to call AllTypes.Unwrap() before calling this method if this AllTypes
// was returned from a transaction, and the transaction was committed or rolled back.
func (at *AllTypes) Update() *AllTypesUpdateOne {
	return (&AllTypesClient{config: at.config}).UpdateOne(at)
}

// Unwrap unwraps the AllTypes entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (at *AllTypes) Unwrap() *AllTypes {
	tx, ok := at.config.driver.(*txDriver)
	if !ok {
		panic("ent: AllTypes is not a transactional entity")
	}
	at.config.driver = tx.drv
	return at
}

// String implements the fmt.Stringer.
func (at *AllTypes) String() string {
	var builder strings.Builder
	builder.WriteString("AllTypes(")
	builder.WriteString(fmt.Sprintf("id=%v", at.ID))
	builder.WriteString(", int=")
	builder.WriteString(fmt.Sprintf("%v", at.Int))
	builder.WriteString(", int8=")
	builder.WriteString(fmt.Sprintf("%v", at.Int8))
	builder.WriteString(", int16=")
	builder.WriteString(fmt.Sprintf("%v", at.Int16))
	builder.WriteString(", int32=")
	builder.WriteString(fmt.Sprintf("%v", at.Int32))
	builder.WriteString(", int64=")
	builder.WriteString(fmt.Sprintf("%v", at.Int64))
	builder.WriteString(", uint=")
	builder.WriteString(fmt.Sprintf("%v", at.Uint))
	builder.WriteString(", uint8=")
	builder.WriteString(fmt.Sprintf("%v", at.Uint8))
	builder.WriteString(", uint16=")
	builder.WriteString(fmt.Sprintf("%v", at.Uint16))
	builder.WriteString(", uint32=")
	builder.WriteString(fmt.Sprintf("%v", at.Uint32))
	builder.WriteString(", uint64=")
	builder.WriteString(fmt.Sprintf("%v", at.Uint64))
	builder.WriteString(", float32=")
	builder.WriteString(fmt.Sprintf("%v", at.Float32))
	builder.WriteString(", float64=")
	builder.WriteString(fmt.Sprintf("%v", at.Float64))
	builder.WriteString(", string_type=")
	builder.WriteString(at.StringType)
	builder.WriteString(", bool=")
	builder.WriteString(fmt.Sprintf("%v", at.Bool))
	builder.WriteString(", uuid=")
	builder.WriteString(fmt.Sprintf("%v", at.UUID))
	builder.WriteString(", time=")
	builder.WriteString(at.Time.Format(time.ANSIC))
	builder.WriteString(", text=")
	builder.WriteString(at.Text)
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", at.State))
	builder.WriteString(", bytes=")
	builder.WriteString(fmt.Sprintf("%v", at.Bytes))
	builder.WriteByte(')')
	return builder.String()
}

// AllTypesSlice is a parsable slice of AllTypes.
type AllTypesSlice []*AllTypes

func (at AllTypesSlice) config(cfg config) {
	for _i := range at {
		at[_i].config = cfg
	}
}
