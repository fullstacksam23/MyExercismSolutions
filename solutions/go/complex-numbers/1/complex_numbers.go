package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
    real float64
    img float64
}
func (n Number) Real() float64 {
	return n.real
}

func (n Number) Imaginary() float64 {
	return n.img
}

func (n1 Number) Add(n2 Number) Number {
	return Number{
        real: n1.real+n2.real,
        img: n1.img+n2.img,
    }
}

func (n1 Number) Subtract(n2 Number) Number {
		return Number{
        real: n1.real-n2.real,
        img: n1.img-n2.img,
    }
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{
        real: (n1.real*n2.real)-(n1.img*n2.img),
        img: (n1.img*n2.real) + (n1.real*n2.img),
    }
}

func (n Number) Times(factor float64) Number {
	return Number{
        real: n.real*factor,
        img: n.img*factor,
    }
}

func (n1 Number) Divide(n2 Number) Number {
	return Number{
		real: (n1.real*n2.real + n1.img*n2.img) / (n2.real*n2.real + n2.img*n2.img),
		img:  (n1.img*n2.real - n1.real*n2.img) / (n2.real*n2.real + n2.img*n2.img),
	}
}

func (n Number) Conjugate() Number {
	return Number{
        real: n.real,
        img: -n.img,
    }
}

func (n Number) Abs() float64 {
	return math.Sqrt((n.real*n.real)+(n.img*n.img))
}

func (n Number) Exp() Number {
    factor := math.Exp(n.real)
	return Number{
        real: factor * math.Cos(n.img),
        img: factor * math.Sin(n.img),
    }
}
