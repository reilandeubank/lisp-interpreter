package interpreter

import (
	// "fmt"
	"reflect"

	"github.com/reilandeubank/golisp/pkg/scanner"
	// "github.com/reilandeubank/golisp/pkg/parser"
)

func isTruthy(object interface{}) bool {
	if object == nil || object == 0.0 || object == 0 { // only false, nil, and 0 are falsey
		return false
	}

	if reflect.TypeOf(object) == reflect.TypeOf(false) {
		return object.(bool)
	}

	return true
}

func isEqual(a interface{}, b interface{}) bool {
	if a == nil && b == nil { // No implicit type conversion for equality, like Go
		return true
	} else if a == nil {
		return false
	}
	return a == b
}

func checkNumberOperand(operator scanner.Token, operand interface{}) (bool, error) {
	if reflect.TypeOf(operand) == reflect.TypeOf(0.0) || reflect.TypeOf(operand) == reflect.TypeOf(0) {
		return true, nil
	}
	return false, &RuntimeError{Token: operator, Message: "Operator must be a number"}
}

func checkNumberOperands(operator scanner.Token, left interface{}, right interface{}) error {
	if reflect.TypeOf(left) == reflect.TypeOf(0.0) && reflect.TypeOf(right) == reflect.TypeOf(0.0) {
		return nil
	}
	return &RuntimeError{Token: operator, Message: "Operators must be numbers"}
}

// func (i *Interpreter) cdr(k parser.Keyword) (interface{}, error) {
// 	var elems []parser.Expression
// 	for _, elem := range k.Args {
// 		newExpr, err := i.evaluate(elem)
// 		if err != nil {
// 			return nil, err
// 		}
// 		if expr, ok := newExpr.(parser.Expression); ok {
// 			elems = append(elems, expr)
// 		} else {
// 			return nil, &RuntimeError{Message: "type assertion failed, expected parser.Expression"}
// 		}
// 	}
// 	fmt.Println(elems)
// 	return parser.ListExpr{Head: elems[1], Tail: elems[2:]}, nil
// }