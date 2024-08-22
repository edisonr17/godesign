package design

import "fmt"

type BikeAction interface {
	Run()
	Stop()
	Status()
	SetStatus(bool)
}

type Bike struct {
	Brand      string
	Model      string
	Year       int
	Colour     string
	bikeStatus bool
}

func (b *Bike) Run() {
	fmt.Println("riding")
}

func (b *Bike) Stop() {
	fmt.Println("stoping")
}

func (b *Bike) Status() {
	fmt.Println("bike is ", b.bikeStatus)
}

func (b *Bike) SetStatus(status bool) {
	b.bikeStatus = status
}
