package table

import (
	"testing"
)

func TestInsertDataShouldReturnErrorIfNotNullConstraintDoesNotSatisfy(t *testing.T) {
	testColumn := column{
		name:     "col1",
		dataType: NumberDataType,
		constraints: []Constraint{
			NotNullConstraintType,
		},
		dataMap: nil,
	}

	actualError := testColumn.insertData(nil, 1)

	if actualError == nil {
		t.Errorf("expected error but it is nil")
	}
}

func TestInsertDataShouldReturnErrorIfUniqueConstraintDoesNotSatisfy(t *testing.T) {
	testColumn := column{
		name:     "col1",
		dataType: NumberDataType,
		constraints: []Constraint{
			NotNullConstraintType,
			UniqueConstraintType,
		},
		dataMap: map[int64]interface{}{
			1: int64(10),
			2: int64(20),
		},
	}

	actualError := testColumn.insertData(int64(20), 1)

	if actualError == nil {
		t.Errorf("expected error but it is nil")
	}
}

func TestInsertDataShouldReturnErrorIfDataTypeDoesNotMatch(t *testing.T) {
	testColumn := column{
		name:        "col1",
		dataType:    NumberDataType,
		constraints: nil,
		dataMap:     nil,
	}

	actualError := testColumn.insertData("abc", 1)

	if actualError == nil {
		t.Errorf("expected error but it is nil")
	}
}

func TestInsertDataShouldReturnNilIfDataTypeMatches(t *testing.T) {
	testColumn := column{
		name:        "col1",
		dataType:    NumberDataType,
		constraints: nil,
		dataMap:     nil,
	}

	actualError := testColumn.insertData(int64(10), 1)

	if actualError != nil {
		t.Errorf("expected nil but got error: %v", actualError)
	}
}
