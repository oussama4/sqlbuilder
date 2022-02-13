package sqlbuilder

// InsertBuilder is a builder for insert queries
type InsertBuilder struct {
	*Builder
	table     string
	columns   []string
	values    [][]interface{}
	returning []string
}

// Insert creates a builder for insert queries
func Insert(table string) *InsertBuilder {
	b := &InsertBuilder{
		Builder:   NewBuilder(),
		table:     table,
		columns:   []string{},
		values:    [][]interface{}{},
		returning: []string{},
	}
	return b
}

// Columns adds a list of columns to the insert query
func (b *InsertBuilder) Columns(columns ...string) *InsertBuilder {
	b.columns = columns
	return b
}

// Values adds column values to the insert query
func (b *InsertBuilder) Values(values ...interface{}) *InsertBuilder {
	b.values = append(b.values, values)
	return b
}

func (b *InsertBuilder) Returning(columns ...string) *InsertBuilder {
	b.returning = columns
	return b
}

// Query builds the sql query and returns it along with its arguments
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

	if len(b.returning) > 0 {
		b.WriteString(" RETURNING ")
		for i, r := range b.returning {
			if i > 0 {
				b.WriteString(", ")
			}
			b.WriteString(r)
		}
	}

	return b.String(), b.args
}
