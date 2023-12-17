package abs_test

import (
	"testing"

	"Section_32_repo/pkg/abs"
)

func TestAbs(t *testing.T) {
	got := abs.Abs(-1.2)
	if got != 1 {
		t.Errorf("Abs(-1) = %.2f; want 1", got)
	}
}
