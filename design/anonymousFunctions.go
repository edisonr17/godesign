package design

import "errors"

func AnonymousFunction(num1, num2 int16, operation string) (int16, error) {
	sum := func(i int16, j int16) (int16, error) {
		return i + j, nil
	}

	subtraction := func(i int16, j int16) (int16, error) {
		return i - j, nil
	}

	multiplication := func(i int16, j int16) (int16, error) {

		return i * j, nil
	}

	division := func(i int16, j int16) (int16, error) {
		if j == 0 {
			return -1, errors.New("Division by zero")
		}
		return i / j, nil
	}

	switch operation {
	case "+":
		return sum(num1, num2)
	case "-":
		return subtraction(num1, num2)
	case "*":
		return multiplication(num1, num2)
	case "/":
		return division(num1, num2)
	default:
		return 0, nil
	}

}
