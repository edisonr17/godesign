package design

import (
	"fmt"
	"testing"
)

func test_builder(t *testing.T) {

	// TODO
	chiguiro := NewChiguiroBuilder().weight(80).name("pikachu").age(10).build()
	fmt.Println(chiguiro)

	if chiguiro.name != "pikachu" {
		t.Error("name is not correct")
	}
	if chiguiro.age != 10 {
		t.Error("age is not correct")
	}

	if chiguiro.weight != 80 {
		t.Error("weigth is not correct")
	}
}
