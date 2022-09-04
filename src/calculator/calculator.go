package calculator

type IntCalculator struct {
	a      int
	b      int
	result int
}

func CreateIntCalculator(a, b int) *IntCalculator {
	calculator := IntCalculator{
		a:      a,
		b:      b,
		result: 0,
	}
	return &calculator
}

func (c *IntCalculator) Add() *IntCalculator {
	c.result = c.a + c.b
	return c
}

func (c *IntCalculator) Subtract() *IntCalculator {
	c.result = c.a - c.b
	return c
}

func (c *IntCalculator) Multiply() *IntCalculator {
	c.result = c.a * c.b
	return c
}

func (c *IntCalculator) Divide() *IntCalculator {
	if c.b != 0 {
		c.result = c.a / c.b
	}
	return c
}

func (c *IntCalculator) Result() int {
	return c.result
}

func (c *IntCalculator) SetValues(a, b int) *IntCalculator {
	c.a = a
	c.b = b
	return c
}
