package sqlbuilder

type Op int

const (
	// Comparison Operators
	OpEQ  Op = iota // =
	OpNEQ           // <>
	OpGT            // >
	OpGTE           // >=
	OpLT            // <
	OpLTE           // <=

	// Logical operation
	OpAND // AND
	OpOR  // OR
	OpNot // NOT
)

var ops = [...]string{
	OpEQ:  " = ",
	OpNEQ: " <> ",
	OpGT:  " > ",
	OpGTE: " >= ",
	OpLT:  " < ",
	OpLTE: " <= ",
	OpAND: " AND ",
	OpOR:  " OR ",
	OpNot: " NOT ",
}

// writeComp writes a comparaison operator
func writeComp(op Op, column string, value interface{}) BuildFunc {
	bf := func(b *Builder) {
		b.WriteString(column)
		b.WriteString(ops[op])
		b.WriteArg(value)
	}
	return bf
}

// Eq builds a `=` operator
func Eq(column string, value interface{}) BuildFunc {
	return writeComp(OpEQ, column, value)
}

// Neq builds a `<>` operator
func Neq(column string, value interface{}) BuildFunc {
	return writeComp(OpNEQ, column, value)
}

// Gt builds a `>` operator
func Gt(column string, value interface{}) BuildFunc {
	return writeComp(OpGT, column, value)
}

// Gte builds a `>=` operator
func Gte(column string, value interface{}) BuildFunc {
	return writeComp(OpGTE, column, value)
}

// Lt builds a `<` operator
func Lt(column string, value interface{}) BuildFunc {
	return writeComp(OpLT, column, value)
}

// Lte builds a `<=` operator
func Lte(column string, value interface{}) BuildFunc {
	return writeComp(OpLTE, column, value)
}

// buildLog build a logical operator
func buildLog(op Op, fns ...BuildFunc) BuildFunc {
	bf := func(b *Builder) {
		for i, fn := range fns {
			if i > 0 || op == OpNot {
				b.WriteString(ops[op])
			}

			b.WriteString("(")
			fn(b)
			b.WriteString(")")

			if op == OpNot {
				break
			}
		}
	}
	return bf
}

// And builds an `AND` operator
func And(fns ...BuildFunc) BuildFunc {
	return buildLog(OpAND, fns...)
}

// Or builds an `OR` operator
func Or(fns ...BuildFunc) BuildFunc {
	return buildLog(OpOR, fns...)
}

// Not builds a `NOT` operator
func Not(fn BuildFunc) BuildFunc {
	return buildLog(OpNot, fn)
}
