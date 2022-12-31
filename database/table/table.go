package table

import (
	"errors"
	"fmt"
)

var (
	ErrColumnNotFound = errors.New("column not found")
)

type Table struct {
	name       string
	columnsMap map[string]*column
	rows       []*row

	rowsCounter int64
}

func CreateTable(name string) *Table {
	return &Table{
		name:        name,
		columnsMap:  make(map[string]*column),
		rowsCounter: 0,
	}
}

func (t *Table) AddColumn(name string, dataType DataType, constraints []Constraint) *Table {
	t.columnsMap[name] = NewColumn(name, dataType, constraints)

	return t
}

func (t *Table) SelectAll() map[string][]interface{} {
	allColumnsData := make(map[string][]interface{})

	for _, r := range t.rows {
		dataMap := r.getDataMap()
		for k, v := range dataMap {
			allColumnsData[k] = append(allColumnsData[k], v)
		}
	}

	return allColumnsData
}

func (t *Table) InsertRow(dataMap map[string]interface{}) error {
	newRow := newEmptyRow()

	//iterating all the column in a table
	for colName, col := range t.columnsMap {
		data, ok := dataMap[colName]
		if ok {
			if err := t.validateConstraints(col, data); err != nil {
				return err
			}

			if err := col.validateDataType(data); err != nil {
				return err
			}

			newRow.setValue(colName, data)

			delete(dataMap, colName)
		} else {
			if !col.isNullable() {
				return errors.New(fmt.Sprintf("column %s is not nullable!", col.getColumnName()))
			}
		}

	}

	for colName := range dataMap {
		return errors.New(fmt.Sprintf("column %s not found!", colName))
	}

	//at this point the insertion was accepted
	t.rowsCounter++
	newRow.build(t.rowsCounter)
	t.rows = append(t.rows, newRow)

	return nil
}

func (t *Table) validateConstraints(col *column, data interface{}) error {
	constraints := col.getConstraints()

	for _, constraint := range constraints {
		switch constraint {
		case NotNullConstraintType:
			if data == nil {
				return errors.New(fmt.Sprintf("%s voilated!", constraint.String()))
			}

		case UniqueConstraintType:
			for _, r := range t.rows {
				colData := r.getDataByColumn(col.getColumnName())
				if colData == data {
					return errors.New(fmt.Sprintf("%s voilated!", constraint.String()))
				}
			}
		}
	}

	return nil
}
