package word

import "testing"

func TestZero(t *testing.T) {
	if !IsZero(1) {
		t.Error(`不是0`)
	}

	if IsZero(0) {
		t.Error(`是0`)
	}
}
