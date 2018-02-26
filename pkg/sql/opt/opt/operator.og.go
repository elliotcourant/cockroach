// Code generated by optgen; DO NOT EDIT.

package opt

const (
	UnknownOp Operator = iota

	// ------------------------------------------------------------
	// Scalar Operators
	// ------------------------------------------------------------

	SubqueryOp

	// VariableOp is the typed scalar value of a column in the query. The private
	// field is a Metadata.ColumnIndex that references the column by index.
	VariableOp

	// ConstOp is a typed scalar constant value. The private field is a tree.Datum
	// value having any datum type that's legal in the expression's context.
	ConstOp

	// TrueOp is the boolean true value that is equivalent to the tree.DBoolTrue datum
	// value. It is a separate operator to make matching and replacement simpler and
	// more efficient, as patterns can contain (True) expressions.
	TrueOp

	// FalseOp is the boolean false value that is equivalent to the tree.DBoolFalse
	// datum value. It is a separate operator to make matching and replacement
	// simpler and more efficient, as patterns can contain (False) expressions.
	FalseOp

	PlaceholderOp

	TupleOp

	// ProjectionsOp is a set of typed scalar expressions that will become output
	// columns for a containing Project operator. The private Cols field contains
	// the list of column indexes returned by the expression, as a *opt.ColList. It
	// is not legal for Cols to be empty.
	ProjectionsOp

	// AggregationsOp is a set of aggregate expressions that will become output
	// columns for a containing GroupBy operator. The private Cols field contains
	// the list of column indexes returned by the expression, as a *ColList. It
	// is legal for Cols to be empty.
	AggregationsOp

	ExistsOp

	// AndOp is the boolean conjunction operator that evalutes to true if all of its
	// conditions evaluate to true. If the conditions list is empty, it evalutes to
	// true.
	AndOp

	// OrOp is the boolean disjunction operator that evalutes to true if any of its
	// conditions evaluate to true. If the conditions list is empty, it evaluates to
	// false.
	OrOp

	// NotOp is the boolean negation operator that evaluates to true if its input
	// evalutes to false.
	NotOp

	EqOp

	LtOp

	GtOp

	LeOp

	GeOp

	NeOp

	InOp

	NotInOp

	LikeOp

	NotLikeOp

	ILikeOp

	NotILikeOp

	SimilarToOp

	NotSimilarToOp

	RegMatchOp

	NotRegMatchOp

	RegIMatchOp

	NotRegIMatchOp

	IsOp

	IsNotOp

	ContainsOp

	BitandOp

	BitorOp

	BitxorOp

	PlusOp

	MinusOp

	MultOp

	DivOp

	FloorDivOp

	ModOp

	PowOp

	ConcatOp

	LShiftOp

	RShiftOp

	FetchValOp

	FetchTextOp

	FetchValPathOp

	FetchTextPathOp

	UnaryPlusOp

	UnaryMinusOp

	UnaryComplementOp

	// FunctionOp invokes a builtin SQL function like CONCAT or NOW, passing the given
	// arguments. The private field is an opt.FuncDef struct that provides the name
	// of the function as well as a pointer to the builtin overload definition.
	FunctionOp

	CoalesceOp

	// UnsupportedExprOp is used for interfacing with the old planner code. It can
	// encapsulate a TypedExpr that is otherwise not supported by the optimizer.
	UnsupportedExprOp

	// ------------------------------------------------------------
	// Relational Operators
	// ------------------------------------------------------------

	// ScanOp returns a result set containing every row in the specified table. Rows
	// and columns are not expected to have any particular ordering. The private
	// Table field is a Metadata.TableIndex that references an optbase.Table
	// definition in the query's metadata.
	ScanOp

	// ValuesOp returns a manufactured result set containing a constant number of rows.
	// specified by the Rows list field. Each row must contain the same set of
	// columns in the same order.
	//
	// The Rows field contains a list of Tuples, one for each row. Each tuple has
	// the same length (same with that of Cols).
	//
	// The Cols field contains the set of column indices returned by each row
	// as a *ColList. It is legal for Cols to be empty.
	ValuesOp

	// SelectOp filters rows from its input result set, based on the boolean filter
	// predicate expression. Rows which do not match the filter are discarded.
	SelectOp

	// ProjectOp modifies the set of columns returned by the input result set. Columns
	// can be removed, reordered, or renamed. In addition, new columns can be
	// synthesized. Projections is a scalar Projections list operator that contains
	// the list of expressions that describe the output columns. The Cols field of
	// the Projections operator provides the indexes of each of the output columns.
	ProjectOp

	// InnerJoinOp creates a result set that combines columns from its left and right
	// inputs, based upon its "on" join predicate. Rows which do not match the
	// predicate are filtered. While expressions in the predicate can refer to
	// columns projected by either the left or right inputs, the inputs are not
	// allowed to refer to the other's projected columns.
	InnerJoinOp

	LeftJoinOp

	RightJoinOp

	FullJoinOp

	SemiJoinOp

	AntiJoinOp

	// InnerJoinApplyOp has the same join semantics as InnerJoin. However, unlike
	// InnerJoin, it allows the right input to refer to columns projected by the
	// left input.
	InnerJoinApplyOp

	LeftJoinApplyOp

	RightJoinApplyOp

	FullJoinApplyOp

	SemiJoinApplyOp

	AntiJoinApplyOp

	// GroupByOp is an operator that is used for performing aggregations (for queries
	// with aggregate functions, HAVING clauses and/or group by expressions). It
	// groups results that are equal on the grouping columns and computes
	// aggregations as described by Aggregations (which is always an Aggregations
	// operator). The arguments of the aggregations are columns from the input.
	GroupByOp

	UnionOp

	IntersectOp

	ExceptOp

	// ------------------------------------------------------------
	// Enforcer Operators
	// ------------------------------------------------------------

	// SortOp enforces the ordering of rows returned by its input expression. Rows can
	// be sorted by one or more of the input columns, each of which can be sorted in
	// either ascending or descending order. See the Ordering field in the
	// PhysicalProps struct.
	// TODO(andyk): Add the Ordering field.
	SortOp

	// NumOperators tracks the total count of operators.
	NumOperators
)

const opNames = "unknownsubqueryvariableconsttruefalseplaceholdertupleprojectionsaggregationsexistsandornoteqltgtlegeneinnot-inlikenot-likei-likenot-i-likesimilar-tonot-similar-toreg-matchnot-reg-matchreg-i-matchnot-reg-i-matchisis-notcontainsbitandbitorbitxorplusminusmultdivfloor-divmodpowconcatl-shiftr-shiftfetch-valfetch-textfetch-val-pathfetch-text-pathunary-plusunary-minusunary-complementfunctioncoalesceunsupported-exprscanvaluesselectprojectinner-joinleft-joinright-joinfull-joinsemi-joinanti-joininner-join-applyleft-join-applyright-join-applyfull-join-applysemi-join-applyanti-join-applygroup-byunionintersectexceptsort"

var opIndexes = [...]uint32{0, 7, 15, 23, 28, 32, 37, 48, 53, 64, 76, 82, 85, 87, 90, 92, 94, 96, 98, 100, 102, 104, 110, 114, 122, 128, 138, 148, 162, 171, 184, 195, 210, 212, 218, 226, 232, 237, 243, 247, 252, 256, 259, 268, 271, 274, 280, 287, 294, 303, 313, 327, 342, 352, 363, 379, 387, 395, 411, 415, 421, 427, 434, 444, 453, 463, 472, 481, 490, 506, 521, 537, 552, 567, 582, 590, 595, 604, 610, 614}

var ScalarOperators = [...]Operator{
	SubqueryOp,
	VariableOp,
	ConstOp,
	TrueOp,
	FalseOp,
	PlaceholderOp,
	TupleOp,
	ProjectionsOp,
	AggregationsOp,
	ExistsOp,
	AndOp,
	OrOp,
	NotOp,
	EqOp,
	LtOp,
	GtOp,
	LeOp,
	GeOp,
	NeOp,
	InOp,
	NotInOp,
	LikeOp,
	NotLikeOp,
	ILikeOp,
	NotILikeOp,
	SimilarToOp,
	NotSimilarToOp,
	RegMatchOp,
	NotRegMatchOp,
	RegIMatchOp,
	NotRegIMatchOp,
	IsOp,
	IsNotOp,
	ContainsOp,
	BitandOp,
	BitorOp,
	BitxorOp,
	PlusOp,
	MinusOp,
	MultOp,
	DivOp,
	FloorDivOp,
	ModOp,
	PowOp,
	ConcatOp,
	LShiftOp,
	RShiftOp,
	FetchValOp,
	FetchTextOp,
	FetchValPathOp,
	FetchTextPathOp,
	UnaryPlusOp,
	UnaryMinusOp,
	UnaryComplementOp,
	FunctionOp,
	CoalesceOp,
	UnsupportedExprOp,
}

var ConstValueOperators = [...]Operator{
	ConstOp,
	TrueOp,
	FalseOp,
}

var BooleanOperators = [...]Operator{
	TrueOp,
	FalseOp,
	AndOp,
	OrOp,
	NotOp,
}

var ComparisonOperators = [...]Operator{
	EqOp,
	LtOp,
	GtOp,
	LeOp,
	GeOp,
	NeOp,
	InOp,
	NotInOp,
	LikeOp,
	NotLikeOp,
	ILikeOp,
	NotILikeOp,
	SimilarToOp,
	NotSimilarToOp,
	RegMatchOp,
	NotRegMatchOp,
	RegIMatchOp,
	NotRegIMatchOp,
	IsOp,
	IsNotOp,
	ContainsOp,
}

var BinaryOperators = [...]Operator{
	BitandOp,
	BitorOp,
	BitxorOp,
	PlusOp,
	MinusOp,
	MultOp,
	DivOp,
	FloorDivOp,
	ModOp,
	PowOp,
	ConcatOp,
	LShiftOp,
	RShiftOp,
	FetchValOp,
	FetchTextOp,
	FetchValPathOp,
	FetchTextPathOp,
}

var UnaryOperators = [...]Operator{
	UnaryPlusOp,
	UnaryMinusOp,
	UnaryComplementOp,
}

var RelationalOperators = [...]Operator{
	ScanOp,
	ValuesOp,
	SelectOp,
	ProjectOp,
	InnerJoinOp,
	LeftJoinOp,
	RightJoinOp,
	FullJoinOp,
	SemiJoinOp,
	AntiJoinOp,
	InnerJoinApplyOp,
	LeftJoinApplyOp,
	RightJoinApplyOp,
	FullJoinApplyOp,
	SemiJoinApplyOp,
	AntiJoinApplyOp,
	GroupByOp,
	UnionOp,
	IntersectOp,
	ExceptOp,
}

var JoinOperators = [...]Operator{
	InnerJoinOp,
	LeftJoinOp,
	RightJoinOp,
	FullJoinOp,
	SemiJoinOp,
	AntiJoinOp,
	InnerJoinApplyOp,
	LeftJoinApplyOp,
	RightJoinApplyOp,
	FullJoinApplyOp,
	SemiJoinApplyOp,
	AntiJoinApplyOp,
}

var JoinApplyOperators = [...]Operator{
	InnerJoinApplyOp,
	LeftJoinApplyOp,
	RightJoinApplyOp,
	FullJoinApplyOp,
	SemiJoinApplyOp,
	AntiJoinApplyOp,
}

var EnforcerOperators = [...]Operator{
	SortOp,
}
