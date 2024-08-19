package design

func AnonymousFunction(num1, num2 int16, operation string) int16 {
	sum := func(i int16, j int16) int16 {
		return i + j
	}

	return sum(num1, num2)

}
