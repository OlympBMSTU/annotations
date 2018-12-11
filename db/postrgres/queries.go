package postgres 

const (
	GET_ANNOTATIONS = "SELECT * FROM annotations "
	GET_BY_AUTHOR_ID = "SELECT * FROM annotations WHERE id = $1" 
)