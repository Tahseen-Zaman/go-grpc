package arithmetic

import "fmt"

type ArithmeticAdapter struct{}

func NewArithmeticAdapter() *ArithmeticAdapter {
	return &ArithmeticAdapter{}
}
func (arith *ArithmeticAdapter) Addition(a, b int) (int, error) {
	return a + b, nil
}
func (arith *ArithmeticAdapter) Subtraction(a, b int) (int, error) {
	return a - b, nil
}
func (arith *ArithmeticAdapter) Multiplication(a, b int) (int, error) {
	return a * b, nil
}
func (arith *ArithmeticAdapter) Division(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
