package result

const (
	NO_ERROR         = 0 // Ok
	CREATED          = 1 // Ok
	PARSE_ERROR      = 2 // Same that empty result, maybe need to delete, also or db error,
	QUERY_ERROR      = 3 // Its unused also dberror
	DB_CONN_ERROR    = 4 // No db, table, connection dropped
	EMPTY_RESULT     = 5 // No data for query
	CONSTRAINT_ERROR = 6 // Duplicates unique keys --- not used
	NO_SUBJECT_ERROR = 7 // While inserting new excerciese we have to choose exisiting subject
)

type DbStatus struct {
	code  int
	descr string
}

func NewStatus(code int, descr string) DbStatus {
	return DbStatus{
		code: code,
		descr: descr,
	}
}

func DefaultStatus() DbStatus {
	return NewStatus(NO_ERROR, "")
}

func (status *DbStatus) SetCode(code int) {
	status.code = code
}

func (status *DbStatus) SetDescription(descr string) {
	status.descr = descr
}



func (status DbStatus) GetCode() int {
	return status.code
}

func (status DbStatus) GetDescription() string {
	return status.descr
}

func (status DbStatus) IsError() bool {
	return !(status.code == NO_ERROR || status.code == CREATED)
}