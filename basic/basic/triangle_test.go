package main

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{8, 15, 17},
	}

	for _, tt := range tests {
		if actual := caTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("caTriangle(%d,%d) got %d ,not %d", tt.a, tt.b, actual, tt.c)
		}
	}

}
