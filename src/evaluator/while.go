package evaluator

import (
	"github.com/AvicennaJr/Nuru/ast"
	"github.com/AvicennaJr/Nuru/object"
)

func evalWhileExpression(we *ast.WhileExpression, env *object.Environment) object.Object {
	condition := Eval(we.Condition, env)
	if isError(condition) {
		return condition
	}
	if isTruthy(condition) {
		evaluated := Eval(we.Consequence, env)
		if isError(evaluated) {
			return evaluated
		}
		if evaluated != nil && evaluated.Type() == object.BREAK_OBJ {
			return evaluated
		}
		evalWhileExpression(we, env)
	}
	return NULL
}