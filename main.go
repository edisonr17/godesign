package main

import (
	"fmt"

	"github.com/edisonr17/godesign/design"
)

func main() {
	fmt.Println("Hola, Aquí vamos a encontrar una recopilación de patrones de diseño en golang")
	//anonymous()
	design.Closures()
}

func anonymous() {
	fmt.Println(design.AnonymousFunction(1, 2, "+"))
	fmt.Println(design.AnonymousFunction(1, 2, "-"))
	fmt.Println(design.AnonymousFunction(1, 2, "*"))
	fmt.Println(design.AnonymousFunction(1, 0, "/"))
}
