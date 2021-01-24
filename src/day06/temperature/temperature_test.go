package temperature

import "testing"

type Temperature float64

type temperatureTest struct {
	i        float64
	expected Temperature
}

var CtoFTests = []temperatureTest{
	{4.1, 39.38},
	{10, 50},
	{-10, 14},
}

var FtoCTests = []temperatureTest{
	{32, 0},
	{50, 32.4},
	{5, -48.6},
}

func TestCtoF(t *testing.T) {
	for _, tt := range CtoFTests {
		actual := CtoF(tt.i)
		if actual != float64(tt.expected) {
			t.Errorf("expected: %v, actual: %v", tt.expected, actual)
		}
	}
}

func TestFtoC(t *testing.T) {
	for _, tt := range FtoCTests {
		actual := FtoC(tt.i)
		if actual != float64(tt.expected) {
			t.Errorf("expected: %v, actual: %v", tt.expected, actual)
		}
	}
}
