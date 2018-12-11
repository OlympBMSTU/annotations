package postgres 

import (
	result "github.com/OlympBMSTU/annotations/db/result"
	"github.com/jackc/pgx"

)

const PG_UNIQUE_CONSTRAINT_CODE = "23505"
const FUNDAMENTAL_QUERY_EMPTY_ERROR = "no rows in result set"


func parseError(err error) result.DbStatus {
	var code int
	var descr string
	if err == nil {
		code = result.NO_ERROR
		descr = ""
	} else {
		switch typedError := err.(type) {
		case pgx.PgError:
			if typedError.Code == PG_UNIQUE_CONSTRAINT_CODE {
				code = result.CONSTRAINT_ERROR
			} else {
				// uncategorized error
				code = result.DB_CONN_ERROR
			}
			descr = typedError.Message
	
		// other errors, also fundamental from pkg/errors
		default:
			// its fucking crutch, but pgx returns fundamental error
			// which is prizvate when row.Scan fails
			// fundamental - is a private package type of package pkg/error
			// also fundamental returns, when there is not enough sended
			// parameters to query
			if typedError.Error() == FUNDAMENTAL_QUERY_EMPTY_ERROR {
				code = result.QUERY_ERROR
			} else {
				code = result.DB_CONN_ERROR
			}
			descr = typedError.Error()

		}
	}

	return result.NewStatus(code, descr)
	// status.SetCode(code)
	// status.SetDescription(descr)

	// return status
}


func CreateResult(data interface{}, err error, params ...interface{}) result.DbResult {
	
	if err != nil {
		return ErrorResult(err)
	}
	if len(params) == 1 {
		return result.OkResult(data, params[0])
	}
	return result.OkResult(data)
}

func ErrorResult(params ...interface{}) result.DbResult {
	if len(params) == 1 {
		return result.NewDbResult(result.DbData{nil}, result.NewStatus{result.NO_ERROR})
	} else {
		return result.DbResult{
			DbData{nil},
			DbStatus{
				params[0].(int),
				params[1].(string),
			},
		}
	}
}
