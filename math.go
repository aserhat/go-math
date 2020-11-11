package math

type Inputs struct {
	InputA int
	InputB int
}

func Add(i Inputs) int {
	return i.InputA + i.InputB
}

func Subtract(i Inputs) int {
	return i.InputA - i.InputB
}
func Multiply(i Inputs) int {
	return i.InputA * i.InputB
}
func Divide(i Inputs) int {
	return i.InputA / i.InputB
}
