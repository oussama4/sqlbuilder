package sqlbuilder

// DeleteBuilder is a builder for delete queries
type DeleteBuilder struct {
	*Builder
	table string
	where BuildFunc
}

// DeleteFrom create a builder for delete queries
func DeleteFrom(table string) *DeleteBuilder {
	db := &DeleteBuilder{
		Builder: NewBuilder(),
		table:   table,
	}
	return db
}

// Where adds a where condition to the delete query
func (db *DeleteBuilder) Where(cond BuildFunc) *DeleteBuilder {
	db.where = cond
	return db
}

// Query builds the sql query and returns it along with its arguments
func (db *DeleteBuilder) Query() (string, []interface{}) {
	db.WriteString("DELETE FROM ")
	db.WriteString(db.table)

	if db.where != nil {
		db.WriteString(" WHERE ")
		db.where(db.Builder)
	}

	return db.String(), db.args
}
