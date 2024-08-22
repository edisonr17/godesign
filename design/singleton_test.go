package design

import "testing"

func singleton(t *testing.T) {

	var chiguiro1 = NewChiguiroInstance()
	chiguiro1.SetWeight(80)

	var chiguiro2 = NewChiguiroInstance()

	if chiguiro1.GetWeight() != chiguiro2.GetWeight() {
		t.Error("s1 and s2 are not the same instance")
	}
}
