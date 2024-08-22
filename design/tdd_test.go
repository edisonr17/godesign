package design

import "testing"

func test_sum(t *testing.T) {
	n1 := 1
	n2 := 4
	expect := 5
	res := sum(n1, n2)
	if res != expect {
		t.Errorf("sum(%d, %d) = %d; want %d", n1, n2, res, expect)
	}
}
