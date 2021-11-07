package sqlbuilder

import (
	"fmt"
	"strings"
)

type BuildFunc func(*Builder)

// Builder is where we put our built sql query
type Builder struct {
	*strings.Builder
	args        []interface{}
	paramsCount int
}

// NewBuilder creates a builder
func NewBuilder() *Builder {
	b := &Builder{
		Builder:     &strings.Builder{},
		args:        []interface{}{},
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
	param := fmt.Sprintf("$%d", b.paramsCount)
	b.WriteString(param)
	b.args = append(b.args, arg)
}
