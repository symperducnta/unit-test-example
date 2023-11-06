package calculator

type calculator struct{}

func NewCalculator() *calculator {
	return &calculator{}
}

func (*calculator) Sum(num1, num2 int) int {
	return num1 + num2
}
