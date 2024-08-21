package design

import "fmt"

func counterByContext() func() int32 {
	count := int32(0)

	return func() int32 {
		count++
		return count
	}
}

func Closures() {
	counter := counterByContext()
	fmt.Println("contador 1", counter) // Output: (1, "")

	newCounter := counterByContext()
	fmt.Println("contador 1", counter()) // Output: (1, "")
	fmt.Println("contador 2", counter()) // Output: (2, "")

	fmt.Println("new contador 1", newCounter())
	fmt.Println("new contador 2", newCounter())
	fmt.Println("new contador 3", newCounter())
}
