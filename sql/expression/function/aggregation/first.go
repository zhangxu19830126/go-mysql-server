package aggregation

import (
	"fmt"

	"github.com/src-d/go-mysql-server/sql"
	"github.com/src-d/go-mysql-server/sql/expression"
)

// First agregation returns the first of all values in the selected column.
// It implements the Aggregation interface.
type First struct {
	expression.UnaryExpression
}

// NewFirst returns a new First node.
func NewFirst(e sql.Expression) *First {
	return &First{expression.UnaryExpression{Child: e}}
}

// Type returns the resultant type of the aggregation.
func (f *First) Type() sql.Type {
	return f.Child.Type()
}

func (f *First) String() string {
	return fmt.Sprintf("FIRST(%s)", f.Child)
}

// TransformUp implements the Transformable interface.
func (f *First) TransformUp(fn sql.TransformExprFunc) (sql.Expression, error) {
	child, err := f.Child.TransformUp(fn)
	if err != nil {
		return nil, err
	}
	return fn(NewFirst(child))
}

// NewBuffer creates a new buffer to compute the result.
func (f *First) NewBuffer() sql.Row {
	return sql.NewRow(nil)
}

// Update implements the Aggregation interface.
func (f *First) Update(ctx *sql.Context, buffer, row sql.Row) error {
	if buffer[0] != nil {
		return nil
	}

	v, err := f.Child.Eval(ctx, row)
	if err != nil {
		return err
	}

	if v == nil {
		return nil
	}

	buffer[0] = v

	return nil
}

// Merge implements the Aggregation interface.
func (f *First) Merge(ctx *sql.Context, buffer, partial sql.Row) error {
	return nil
}

// Eval implements the Aggregation interface.
func (f *First) Eval(ctx *sql.Context, buffer sql.Row) (interface{}, error) {
	return buffer[0], nil
}
