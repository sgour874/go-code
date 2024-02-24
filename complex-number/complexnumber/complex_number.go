package complexnumber

import "fmt"

type Complex struct {
	real float64
	img  float64
}

func GetComplex(real float64, img float64) Complex {
	return Complex{
		real: real,
		img:  img,
	}
}

func Add(n1 Complex, n2 Complex) Complex {
	return Complex{
		n1.real + n2.real,
		n1.img + n2.img,
	}
}

func Multiply(n1 Complex, n2 Complex) Complex {
	var ans Complex
	ans.real = (n1.real * n2.real) - (n1.img * n2.img)
	ans.img = (n1.real * n2.img) + (n1.img * n2.real)
	return ans
}

func Substract(n1 Complex, n2 Complex) Complex {
	return Complex{
		n1.real - n2.real,
		n1.img - n2.img,
	}
}

func negation(n Complex) Complex {
	return Complex{
		n.real,
		n.img * (-1),
	}
}

func Divisionn(n1 Complex, n2 Complex) Complex {
	base := n2.real*n2.real + n2.img*n2.img
	numerator := Multiply(n1, negation(n2))

	return Complex{
		numerator.real / base,
		numerator.img / base,
	}

}

func Print(n Complex) {
	if n.img < 0 {
		fmt.Printf("%f - i%f", n.real, (-1)*n.img)
	} else if n.img > 0 {
		fmt.Printf("%f + i%f", n.real, n.img)
	} else {
		fmt.Printf("%f", n.real)
	}
	fmt.Println()
}
