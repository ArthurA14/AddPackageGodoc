package abs

// import (
// 	"testing"
// 	"math"
// )

func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// func TestAbs(t *testing.T) {
// 	got := Abs(-1)
// 	if got != 1 {
// 		t.Errorf("Abs(-1) = %v; want 1", got)
// 	}
// }
