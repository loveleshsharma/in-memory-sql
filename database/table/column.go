package table

import (
	"errors"
	"fmt"
)

type column struct {
	name        string
	dataType    DataType
	constraints []Constraint
}

func NewColumn(name string, dataType DataType, constraints []Constraint) *column {
	return &column{
		name:        name,
		dataType:    dataType,
		constraints: constraints,
	}
}

func (c *column) validateDataType(data interface{}) error {
	if data != nil { //value can be null if the column is nullable
		if !c.dataType.isValidType(data) {
			return errors.New(fmt.Sprintf("invalid value for %s: %v", c.dataType.string(), data))
		}
	}

	return nil
}
func (c *column) isNullable() bool {
	for _, c := range c.constraints {
		if c == NotNullConstraintType {
			return false
		}
	}

	return true
}

func (c *column) getConstraints() []Constraint {
	return c.constraints
}

func (c *column) getColumnName() string {
	return c.name
}
