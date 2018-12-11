package result

type DbData struct {
	data interface{}
}

func EmptyData() DbData {
	return NewData(nil)
}

func NewData(data interface{}) DbData {
	return DbData{data}
}

func CreateDbData(data interface{}) DbData {
	return DbData{
		data,
	}
}

func (dat DbData) GetData() interface{} {
	return dat.data
}
