package figuras

import "math"

type Circulo struct {
	Radio float64
}

func (cir *Circulo) area() float64 {
	return math.Phi * (cir.Radio * cir.Radio)
}

func (cir *Circulo) perimetro() float64 {
	return 2 * math.Phi * cir.Radio
}
