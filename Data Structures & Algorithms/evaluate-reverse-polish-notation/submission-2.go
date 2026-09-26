import(
	"slices"
)

func evalRPN(tokens []string) int {

	operators := []string{"+", "-", "/", "*"}
	numbers := make([]int, 0)
	var Result int

	if len(tokens) == 1 {
		Result , _ = strconv.Atoi(tokens[0])
		return Result
	}

	for i, num := range tokens {
		if slices.Contains(operators, num) {
			switch num {
			case "+":
				Result = numbers[len(numbers)-2] + numbers[len(numbers)-1]
			case "-":
				Result = numbers[len(numbers)-2] - numbers[len(numbers)-1]
			case "*":
				Result = numbers[len(numbers)-2] * numbers[len(numbers)-1]
			case "/":
				Result = numbers[len(numbers)-2] / numbers[len(numbers)-1]
			}
			numbers = numbers[:len(numbers)-1]
			numbers[len(numbers)-1] = Result

		} else {
			number, _ := strconv.ParseInt(tokens[i], 10, 64)
			numbers = append(numbers, int(number))

		}

	}
	return Result

}