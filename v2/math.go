// Package math provides basic mathematical operations.
// This package includes functions for addition and other arithmetic operations.
package math

import "golang.org/x/exp/constraints"

// Number is combined the Integer and Float types
type Number interface {
	constraints.Integer | constraints.Float
}

// Add takes two integers and returns their sum.
// For more information on addition, visit [mathsisfun]
//
// [mathsisfun]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
