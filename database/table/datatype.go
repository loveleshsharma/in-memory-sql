package table

type DataType interface {
	isValidType(interface{}) bool
	string() string
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

func (n NumberType) string() string {
	return "number type"
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

func (v VarcharType) string() string {
	return "varchar type"
}

var NumberDataType NumberType
var VarcharDataType VarcharType
