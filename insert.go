package sqlbuilder

type InsertBuilder struct {
	*Builder
	table   string
	columns []string
	values  [][]interface{}
}

func Insert(table string) *InsertBuilder {
	b := &InsertBuilder{
		Builder: NewBuilder(),
		table:   table,
		columns: []string{},
		values:  [][]interface{}{},
	}
	return b
}

func (b *InsertBuilder) Columns(columns ...string) *InsertBuilder {
	b.columns = columns
	return b
}

func (b *InsertBuilder) Values(values ...interface{}) *InsertBuilder {
	b.values = append(b.values, values)
	return b
}

func (b *InsertBuilder) Query() (string, []interface{}) {
	b.WriteString("INSERT INTO ")
	b.WriteString(b.table)
	b.WriteString(" (")

	for i, c := range b.columns {
		if i > 0 {
			b.WriteString(",")
		}
		b.WriteString(c)
	}

	b.WriteString(") ")
	b.WriteString("VALUES ")

	for i, row := range b.values {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString("(")
		for j, col := range row {
			if j > 0 {
				b.WriteString(", ")
			}
			b.WriteArg(col)
		}
		b.WriteString(")")
	}

	return b.String(), b.args
}
