// Package math provides basic math functions for two numbers
package math

// Inputs stores the the two numbers you want do math on
type Inputs struct {
	InputA int
	InputB int
}

// Add adds InputA and InputB
func (i Inputs) Add() int {
	return i.InputA + i.InputB
}

// Subtract subtracts InputB from InputA
func (i Inputs) Subtract() int {
	return i.InputA - i.InputB
}

// Multiply multiplies InputA and InputB
func (i Inputs) Multiply() int {
	return i.InputA * i.InputB
}

// Divide divides InputA and InputB
func (i Inputs) Divide() int {
	return i.InputA / i.InputB
}
