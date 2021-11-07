package sqlbuilder

import "strings"

type joinType int

const (
	inner joinType = iota
	left
	right
)

type join struct {
	joinType joinType
	table    string
	on       BuildFunc
}

func buildJoin(j join, b *Builder) {
	switch j.joinType {
	case inner:
		b.WriteString(" INNER")
	case left:
		b.WriteString(" LEFT")
	case right:
		b.WriteString(" RIGHT")
	}
	b.WriteString(" JOIN ")
	b.WriteString(j.table)
	j.on(b)
}

type SelectBuilder struct {
	*Builder
	table   string
	columns []string
	joins   []join
	where   BuildFunc
	group   []string
	having  BuildFunc
	order   []string
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

func (sb *SelectBuilder) Join(table string, on BuildFunc) *SelectBuilder {
	sb.joins = append(sb.joins, join{inner, table, on})
	return sb
}

func (sb *SelectBuilder) LeftJoin(table string, on BuildFunc) *SelectBuilder {
	sb.joins = append(sb.joins, join{left, table, on})
	return sb
}

func (sb *SelectBuilder) RightJoin(table string, on BuildFunc) *SelectBuilder {
	sb.joins = append(sb.joins, join{right, table, on})
	return sb
}

func (sb *SelectBuilder) GroupBy(columns ...string) *SelectBuilder {
	sb.group = append(sb.group, columns...)
	return sb
}

func (sb *SelectBuilder) Having(cond BuildFunc) *SelectBuilder {
	sb.having = cond
	return sb
}

func (sb *SelectBuilder) OrderBy(columns ...string) *SelectBuilder {
	sb.order = append(sb.order, columns...)
	return sb
}

func (sb *SelectBuilder) Query() (string, []interface{}) {
	sb.WriteString("SELECT ")
	sb.WriteString(strings.Join(sb.columns, ", "))
	sb.WriteString(" FROM ")
	sb.WriteString(sb.table)

	for _, j := range sb.joins {
		buildJoin(j, sb.Builder)
	}

	sb.WriteString(" WHERE ")
	sb.where(sb.Builder)

	sb.WriteString(strings.Join(sb.group, ", "))
	sb.WriteString(" HAVING ")
	sb.having(sb.Builder)

	sb.WriteString(strings.Join(sb.order, ", "))

	return sb.String(), sb.args
}
