package sqlbuilder

// UpdateBuilder is a builder for update queries
type UpdateBuilder struct {
	*Builder
	table   string
	columns map[string]interface{}
	where   BuildFunc
}

// Update creates a builder for update queries
func Update(table string) *UpdateBuilder {
	b := &UpdateBuilder{
		table:   table,
		columns: map[string]interface{}{},
		Builder: NewBuilder(),
	}
	return b
}

// Set sets the column to the provided value
func (b *UpdateBuilder) Set(column string, value interface{}) *UpdateBuilder {
	b.columns[column] = value
	return b
}

// Where adds a where condition to the update query
func (b *UpdateBuilder) Where(cond BuildFunc) *UpdateBuilder {
	b.where = cond
	return b
}

// Query builds the sql query and returns it along with its arguments
func (b *UpdateBuilder) Query() (string, []interface{}) {
	b.WriteString("UPDATE ")
	b.WriteString(b.table)
	b.WriteString(" SET ")

	for c, v := range b.columns {
		if b.paramsCount > 0 {
			b.WriteString(", ")
		}
		b.WriteString(c)
		b.WriteString(" = ")
		b.WriteArg(v)
	}

	if b.where != nil {
		b.WriteString(" WHERE ")
		b.where(b.Builder)
	}

	return b.String(), b.args
}
