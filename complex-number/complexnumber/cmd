package complexnumber

func (c *Complex) Real() float64 {
	return c.real
}
func (c *Complex) Img() float64 {
	return c.img
}

func (c *Complex) Add(n Complex) {
	c.real += n.real
	c.img += n.img
}

func (c *Complex) Multiply(n Complex) {
	c.real = (c.real * n.real) - (c.img * n.img)
	c.img = (c.real * n.img) + (c.img * n.real)
}

func (c *Complex) Substract(n Complex) {
	c.real -= n.real
	c.img -= n.img
}

func (c *Complex) Divisionn(n Complex) {
	base := n.real*n.real + n.img*n.img
	c.Multiply(negation(n))
	c.real /= base
	c.img /= base
}

func (c *Complex) Print() {
	Print(*c)
}
