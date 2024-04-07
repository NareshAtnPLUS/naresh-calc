package calc

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (calc Adapter) Addition(x, y int) (int, error) {
	sum := x + y
	return sum, nil
}
func (calc Adapter) Subtraction(x, y int) (int, error) {
	difference := x - y
	return difference, nil
}
func (calc Adapter) Multiplication(x, y int) (int, error) {
	product := x * y
	return product, nil
}
func (calc Adapter) Division(x, y int) (int, int, error) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder, nil
}
