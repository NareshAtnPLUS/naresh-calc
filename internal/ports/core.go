package ports

type CalculationPort interface {
	Addition(x, y int) (int, error)
	Subtraction(x, y int) (int, error)
	Multiplication(x, y int) (int, error)
	Division(x, y int) (int, int, error)
}
