package temperature

func CtoF(c float64) float64 {
	r := (c * (9.0 / 5.0)) + 32
	return r
}

func FtoC(f float64) float64 {
	r := (f - 32) * (9.0 / 5.0)
	return r
}
