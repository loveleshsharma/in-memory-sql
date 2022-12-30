package table

import (
	"errors"
	"fmt"
)

type column struct {
	name        string
	dataType    DataType
	constraints []Constraint

	dataMap map[int64]interface{}
}

func NewColumn(name string, dataType DataType, constraints []Constraint) *column {
	return &column{
		name:        name,
		dataType:    dataType,
		constraints: constraints,
		dataMap:     make(map[int64]interface{}),
	}
}

func (c *column) insertData(data interface{}, index int64) error {
	fmt.Println("data ", data)
	if err := c.satisfiesConstraints(data); err != nil {
		return err
	}

	if !c.dataType.isValidType(data) {
		return errors.New("invalid data type")
	}

	c.dataMap[index] = data

	return nil
}

func (c *column) getAllValues() []interface{} {
	values := make([]interface{}, len(c.dataMap))

	for _, v := range c.dataMap {
		values = append(values, v)
	}

	return values
}

func (c *column) satisfiesConstraints(data interface{}) error {

	for _, constraint := range c.constraints {
		if !c.isConstraintSatisfies(constraint, data) {
			return errors.New(constraint.String() + "does not satisfy!")
		}
	}

	return nil
}

func (c *column) isConstraintSatisfies(constraint Constraint, data interface{}) bool {

	switch constraint {
	case NotNullConstraintType:
		if data == nil {
			return false
		}

	case UniqueConstraintType:
		for _, v := range c.dataMap {
			if v == data {
				return false
			}
		}
	}

	return true
}

func (c *column) isNullable() bool {
	for _, c := range c.constraints {
		if c == NotNullConstraintType {
			return true
		}
	}

	return false
}

func (c *column) getColumnName() string {
	return c.name
}
