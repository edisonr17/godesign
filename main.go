package main

import (
	"fmt"

	"github.com/edisonr17/godesign/design"
)

func main() {
	fmt.Println("Hola, Aquí vamos a encontrar una recopilación de patrones de diseño en golang")
	//anonymous()
	//design.Closures()
	//exampleStruct()
	exampleInterface()
}

func exampleInterface() {
	trek := design.Bike{
		Brand:  "Trek",
		Model:  "Emonda",
		Year:   2020,
		Colour: "Black",
	}
	StartRide(&trek)
	StopRide(&trek)

}

func StopRide(bike design.BikeAction) {
	bike.Stop()
	bike.SetStatus(false)
}

func StartRide(bike design.BikeAction) {
	bike.Run()
	bike.Stop()
	bike.SetStatus(false)
}

func anonymous() {
	fmt.Println(design.AnonymousFunction(1, 2, "+"))
	fmt.Println(design.AnonymousFunction(1, 2, "-"))
	fmt.Println(design.AnonymousFunction(1, 2, "*"))
	fmt.Println(design.AnonymousFunction(1, 0, "/"))
}

func exampleStruct() {
	trek := design.Bicycle{
		Brand:  "Trek",
		Model:  "Emonda",
		Year:   2020,
		Colour: "Black",
	}

	trek.Run()
	trek.Status()
	trek.Stop()
	trek.Status()

	giant := design.Bicycle{
		Brand:  "Giant",
		Model:  "Escape",
		Year:   2020,
		Colour: "Red",
	}

	giant.Run()
	giant.Status()
	giant.Stop()
	giant.Status()
}
