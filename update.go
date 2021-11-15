package sqlbuilder

type UpdateBuilder struct {
	*Builder
	table   string
	columns map[string]interface{}
	where   BuildFunc
}

func Update(table string) *UpdateBuilder {
	b := &UpdateBuilder{
		table:   table,
		columns: map[string]interface{}{},
		Builder: NewBuilder(),
	}
	return b
}

func (b *UpdateBuilder) Set(column string, value interface{}) *UpdateBuilder {
	b.columns[column] = value
	return b
}

func (b *UpdateBuilder) Where(pred BuildFunc) *UpdateBuilder {
	b.where = pred
	return b
}

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
