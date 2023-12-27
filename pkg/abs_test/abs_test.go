package abs_test

import (
	"Section_31_repo/pkg/abs"
	"testing"
)

func TestAbs(t *testing.T) {
	got := abs.Abs(-1.2)
	if got != 1 {
		t.Errorf("Abs(-1) = %.2f; want 1", got)
	}
}
