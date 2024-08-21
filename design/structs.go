package design

import "fmt"

type Bicycle struct {
	Brand      string
	Model      string
	Year       int
	Colour     string
	bikeStatus bool
}

func (bicycle *Bicycle) Run() {
	fmt.Println("the bicycle ", bicycle.Brand, "is running")
	bicycle.bikeStatus = true
}

func (bicycle *Bicycle) Stop() {
	fmt.Println("the bicycle ", bicycle.Brand, "is stopped")
	bicycle.bikeStatus = false
}

func (bicycle *Bicycle) Status() {
	fmt.Println("The bicycle is ", bicycle.Brand, bicycle.bikeStatus)
}
