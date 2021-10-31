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
)

var ops = [...]string{
	OpEQ:  " = ",
	OpNEQ: " <> ",
	OpGT:  " > ",
	OpGTE: " >= ",
	OpLT:  " < ",
	OpLTE: " <= ",
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
