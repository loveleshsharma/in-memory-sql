package table

type row struct {
	rowId   int64
	dataMap map[string]interface{}
}

func newEmptyRow() *row {
	return &row{
		dataMap: make(map[string]interface{}),
	}
}

func (r *row) getDataMap() map[string]interface{} {
	return r.dataMap
}

func (r *row) getDataByColumn(colName string) (interface{}, error) {
	data, ok := r.dataMap[colName]
	if !ok {
		return nil, ErrColumnNotFound
	}

	return data, nil
}

func (r *row) setValue(column string, data interface{}) {
	r.dataMap[column] = data
}

func (r *row) build(rowId int64) {
	r.rowId = rowId
}
