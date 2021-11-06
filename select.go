package sqlbuilder

import "strings"

type SelectBuilder struct {
	*Builder
	table   string
	columns []string
	where   BuildFunc
}

func Select(columns ...string) *SelectBuilder {
	sb := &SelectBuilder{
		Builder: NewBuilder(),
		columns: columns,
	}
	return sb
}

func (sb *SelectBuilder) From(table string) *SelectBuilder {
	sb.table = table
	return sb
}

func (sb *SelectBuilder) Where(pred BuildFunc) *SelectBuilder {
	sb.where = pred
	return sb
}

func (sb *SelectBuilder) Query() (string, []interface{}) {
	sb.WriteString("SELECT ")
	sb.WriteString(strings.Join(sb.columns, ", "))
	sb.WriteString(" FROM ")
	sb.WriteString(sb.table)
	sb.WriteString(" WHERE ")
	sb.where(sb.Builder)

	return sb.String(), sb.args
}
