package golang

type OpFlags uint64

const (
	OpEq OpFlags = 1 << iota
	OpNeq
	OpGt
	OpGte
	OpLt
	OpLte

	OpLike
	OpILike
	OpNotLike
	OpNotILike
	OpSimilar
	OpNotSimilar

	OpIn
	OpNotIn

	OpIsNull
	OpIsNotNull

	OpBetween
	OpNotBetween

	OpExists
	OpNotExists

	OpArrayContains
	OpArrayContainedBy
	OpArrayOverlap

	OpJsonContains
	OpJsonContainedBy
	OpJsonHasKey
	OpJsonHasAnyKeys
	OpJsonHasAllKeys
)

const NullMask = OpIsNull | OpIsNotNull

const BoolOps = OpEq | OpNeq
const NullBoolOps = BoolOps | NullMask
const FloatOps = OpEq | OpNeq | OpGt | OpGte | OpLt | OpLte | OpBetween | OpNotBetween
const NullFloatOps = FloatOps | NullMask
const IntOps = OpEq | OpNeq | OpGt | OpGte | OpLt | OpLte
const NullIntOps = IntOps | NullMask
const StringOps = OpEq | OpNeq | OpLike | OpILike | OpNotLike | OpNotILike | OpSimilar | OpNotSimilar
const NullStringOps = StringOps | NullMask
const ArrayOps = OpEq | OpNeq | OpIn | OpNotIn | OpArrayContains | OpArrayContainedBy | OpArrayOverlap
const NullArrayOps = ArrayOps | NullMask
const JsonOps = OpEq | OpNeq | OpIn | OpNotIn | OpJsonContains | OpJsonContainedBy | OpJsonHasKey | OpJsonHasAnyKeys | OpJsonHasAllKeys
const NullJsonOps = JsonOps | NullMask
const TimeOps = OpEq | OpNeq | OpGt | OpGte | OpLt | OpLte | OpBetween | OpNotBetween
const NullTimeOps = TimeOps | NullMask

const ExistsOps = OpExists | OpNotExists

func GetFlags(flag OpFlags) []OpFlags {
	var ops []OpFlags
	for i := 0; i < 64; i++ {
		if flag&(1<<i) != 0 {
			ops = append(ops, 1<<i)
		}
	}
	return ops
}

func (op OpFlags) Name() string {
	switch op {
	case OpEq:
		return "Eq"
	case OpNeq:
		return "Neq"
	case OpGt:
		return "Gt"
	case OpGte:
		return "Gte"
	case OpLt:
		return "Lt"
	case OpLte:
		return "Lte"
	case OpLike:
		return "Like"
	case OpILike:
		return "ILike"
	case OpNotLike:
		return "NotLike"
	case OpNotILike:
		return "NotILike"
	case OpSimilar:
		return "Similar"
	case OpNotSimilar:
		return "NotSimilar"
	case OpIn:
		return "In"
	case OpNotIn:
		return "NotIn"
	case OpIsNull:
		return "IsNull"
	case OpIsNotNull:
		return "IsNotNull"
	case OpBetween:
		return "Between"
	case OpNotBetween:
		return "NotBetween"
	case OpExists:
		return "Exists"
	case OpNotExists:
		return "NotExists"
	case OpArrayContains:
		return "ArrayContains"
	case OpArrayContainedBy:
		return "ArrayContainedBy"
	case OpArrayOverlap:
		return "ArrayOverlap"
	case OpJsonContains:
		return "JsonContains"
	case OpJsonContainedBy:
		return "JsonContainedBy"
	case OpJsonHasKey:
		return "JsonHasKey"
	case OpJsonHasAnyKeys:
		return "JsonHasAnyKeys"
	case OpJsonHasAllKeys:
		return "JsonHasAllKeys"
	default:
		return ""
	}
}

func (op OpFlags) String() string {
	switch op {
	case OpEq:
		return "="
	case OpNeq:
		return "!="
	case OpGt:
		return ">"
	case OpGte:
		return ">="
	case OpLt:
		return "<"
	case OpLte:
		return "<="
	case OpLike:
		return "LIKE"
	case OpILike:
		return "ILIKE"
	case OpNotLike:
		return "NOT LIKE"
	case OpNotILike:
		return "NOT ILIKE"
	case OpSimilar:
		return "SIMILAR TO"
	case OpNotSimilar:
		return "NOT SIMILAR TO"
	case OpIn:
		return "IN"
	case OpNotIn:
		return "NOT IN"
	case OpIsNull:
		return "IS NULL"
	case OpIsNotNull:
		return "IS NOT NULL"
	case OpBetween:
		return "BETWEEN"
	case OpNotBetween:
		return "NOT BETWEEN"
	case OpExists:
		return "EXISTS"
	case OpNotExists:
		return "NOT EXISTS"
	case OpArrayContains:
		return "@>"
	case OpArrayContainedBy:
		return "<@"
	case OpArrayOverlap:
		return "&&"
	case OpJsonContains:
		return "@>"
	case OpJsonContainedBy:
		return "<@"
	case OpJsonHasKey:
		return "?"
	case OpJsonHasAnyKeys:
		return "?|"
	case OpJsonHasAllKeys:
		return "?&"
	default:
		return ""
	}
}
