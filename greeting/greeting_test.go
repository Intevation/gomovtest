package greeting

import "testing"

func Test_SignOfLife(t *testing.T) {
	result := SignOfLife(2)
	if result != "Beeep" {
		t.Error("incorrect result: expected Beeep, got", result)
	}
}
