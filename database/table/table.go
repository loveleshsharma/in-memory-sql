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

	rowsCounter int64
}

func NewTable(name string) Table {
	return Table{
		name:        name,
		columnsMap:  nil,
		rowsCounter: 0,
	}
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
	for k, v := range t.columnsMap {
		allColumnsData[k] = v.getAllValues()
	}

	return allColumnsData
}

func (t *Table) Insert(data map[string]interface{}) error {
	newRow := t.createNewRow()
	t.rowsCounter++

	for k, v := range data {
		currColumn, err := t.getColumnByName(k)
		if err != nil {
			//column not found
			t.rowsCounter--
			return ErrColumnNotFound
		}

		if err = currColumn.insertData(v, t.rowsCounter); err != nil {
			return err
		}

		delete(newRow, k)
	}

	//iterating remaining items in the row that were not inserted to check if they were not nullable
	for _, v := range newRow {
		if !v.isNullable() {
			return errors.New("column " + v.getColumnName() + " is not nullable")
		}
	}

	fmt.Println("insert was successful")
	return nil
}

func (t *Table) createNewRow() map[string]column {
	row := make(map[string]column, len(t.columnsMap))

	i := 0
	for k, v := range t.columnsMap {
		row[k] = *v
		i++
	}

	return row
}

func (t *Table) getColumnByName(name string) (*column, error) {
	val, ok := t.columnsMap[name]
	if !ok {
		return nil, errors.New("column not found")
	}

	return val, nil
}

/*
	create table employee(
		empName varchar2(1000) not null unique
	);

	insert into employee(empName, salary) values('Lovelesh', 100);
*/
