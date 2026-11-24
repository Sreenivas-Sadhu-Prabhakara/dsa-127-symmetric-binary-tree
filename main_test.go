package main

import "testing"

func TestIsSymmetric(t *testing.T) {
	sym := tn(1, tn(2, tn(3, nil, nil), tn(4, nil, nil)), tn(2, tn(4, nil, nil), tn(3, nil, nil)))
	if !IsSymmetric(sym) {
		t.Error("expected symmetric")
	}
	// Same values but un-mirrored positions.
	asym := tn(1, tn(2, nil, tn(3, nil, nil)), tn(2, nil, tn(3, nil, nil)))
	if IsSymmetric(asym) {
		t.Error("expected asymmetric: both 3s on the right")
	}
	if !IsSymmetric(nil) {
		t.Error("empty tree is symmetric")
	}
	if !IsSymmetric(tn(1, nil, nil)) {
		t.Error("single node is symmetric")
	}
	// Values must mirror too, not just shape.
	vals := tn(1, tn(2, nil, nil), tn(9, nil, nil))
	if IsSymmetric(vals) {
		t.Error("expected asymmetric values")
	}
}
