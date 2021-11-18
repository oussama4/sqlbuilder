package sqlbuilder

type DeleteBuilder struct {
	*Builder
	table string
	where BuildFunc
}

func DeleteFrom(table string) *DeleteBuilder {
	db := &DeleteBuilder{
		Builder: NewBuilder(),
		table:   table,
	}
	return db
}

func (db *DeleteBuilder) Where(cond BuildFunc) *DeleteBuilder {
	db.where = cond
	return db
}

func (db *DeleteBuilder) Query() (string, []interface{}) {
	db.WriteString("DELETE FROM ")
	db.WriteString(db.table)

	if db.where != nil {
		db.WriteString(" WHERE ")
		db.where(db.Builder)
	}

	return db.String(), db.args
}
