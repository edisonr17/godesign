package design

import "testing"

func sintleton(t *testing.T) {

	var s1 = GetChiguiroInstance()
	var s2 = GetChiguiroInstance()

	if s1 != s2 {
		t.Error("s1 and s2 are not the same instance")
	}
}
