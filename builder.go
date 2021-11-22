package sqlbuilder

import (
	"fmt"
	"strings"
)

// Dialect names
const (
	Postgres = "postgres"
	Sqlite   = "sqlite"
	Mysql    = "mysql"
)

// the BuildFunc type is used to add sql expressions to the provided builder
type BuildFunc func(*Builder)

// Builder is where we put our built sql query
type Builder struct {
	*strings.Builder
	args        []interface{}
	dialect     string
	paramsCount int
}

// NewBuilder creates a builder
func NewBuilder() *Builder {
	b := &Builder{
		Builder:     &strings.Builder{},
		args:        []interface{}{},
		dialect:     Postgres,
		paramsCount: 0,
	}
	return b
}

// Expr adds a raw sql expression to the query builder
func Expr(expr string, args ...interface{}) BuildFunc {
	bf := func(b *Builder) {
		b.WriteString(expr)
		for _, arg := range args {
			b.WriteArg(arg)
		}
	}
	return bf
}

// WriteArg adds an input argument to builder
func (b *Builder) WriteArg(arg interface{}) {
	b.paramsCount++
	param := "?"
	if b.dialect == Postgres {
		param = fmt.Sprintf("$%d", b.paramsCount)
	}
	b.WriteString(param)
	b.args = append(b.args, arg)
}

// DialectBuilder is for building queries for a custom sql dialect
type DialectBuilder struct {
	dialect string
}

// Dialect creates a custom builder for the given sql dialect
func Dialect(dialect string) DialectBuilder {
	db := DialectBuilder{dialect: dialect}
	return db
}

// Select creates a SelectBuilder for the configured dialect
func (db DialectBuilder) Select(columns ...string) *SelectBuilder {
	sb := Select(columns...)
	sb.dialect = db.dialect
	return sb
}

// Insert creates an InsertBuilder for the configured dialect
func (db DialectBuilder) Insert(table string) *InsertBuilder {
	ib := Insert(table)
	ib.dialect = db.dialect
	return ib
}

// Update creates an UpdateBuilder for the configured dialect
func (db DialectBuilder) Update(table string) *UpdateBuilder {
	ub := Update(table)
	ub.dialect = db.dialect
	return ub
}

// DeleteFrom creates a DeleteBuilder for the configured dialect
func (db DialectBuilder) DeleteFrom(table string) *DeleteBuilder {
	dlb := DeleteFrom(table)
	dlb.dialect = db.dialect
	return dlb
}
