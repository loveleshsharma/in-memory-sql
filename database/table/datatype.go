package table

type DataType interface {
	isValidType(interface{}) bool
}

type NumberType struct {
}

func (n NumberType) isValidType(data interface{}) bool {
	_, ok := data.(int64)
	if !ok {
		return false
	}

	return true
}

type VarcharType struct {
	value string
}

func (v VarcharType) isValidType(data interface{}) bool {
	_, ok := data.(string)
	if !ok {
		return false
	}

	return true
}

var NumberDataType NumberType
var VarcharDataType VarcharType
