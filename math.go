// Package math provides basic math functions.
package math

// Inputs setups the two numbers you want do math on
type Inputs struct {
	InputA int
	InputB int
}

// Add adds InputA and InputB
func Add(i Inputs) int {
	return i.InputA + i.InputB
}

// Subtract subtracts InputB from InputA
func Subtract(i Inputs) int {
	return i.InputA - i.InputB
}

// Multiply multiplies InputA and InputB
func Multiply(i Inputs) int {
	return i.InputA * i.InputB
}

// Divide divides InputA and InputB
func Divide(i Inputs) int {
	return i.InputA / i.InputB
}
