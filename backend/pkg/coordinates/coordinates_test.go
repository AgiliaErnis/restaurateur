package coordinates

import (
	"math"
	"strconv"
	"testing"
)

func FloatToString(inputNum float64) string {
	return strconv.FormatFloat(inputNum, 'f', 6, 64)
}

func TestToRadians(t *testing.T) {
	got := toRadians(10)
	want := 10 * (math.Pi / 180)
	if got != want {
		t.Errorf("got %q, wanted %q", FloatToString(got), FloatToString(want))
	}
}

func TestHaversine(t *testing.T) {
	got := int(math.Round(Haversine(50.078702, 14.439827, 50.076482, 14.439048))) //random coords next to praguecollege
	want := 253
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
