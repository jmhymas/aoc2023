package main

import (
	"fmt"
	"os"
	"regexp"
)

func getSumOfCalValues(input []byte) int {
	var numbers []int
	var sum = 0

	for _, value := range input {
		fmt.Printf(string(value))
		if value > 47 && value < 58 {
			numbers = append(numbers, int(value-48))
		} else if value == 10 {
			calibrationValue := numbers[0]*10 + numbers[len(numbers)-1]
			fmt.Printf("First: %v, Last: %v, CalValue: %v\n\n", numbers[0], numbers[len(numbers)-1], calibrationValue)
			sum += calibrationValue
			numbers = nil
		}
	}

	return sum
}

func replaceText(input []byte) []byte {
	numbersText := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	inputString := string(input)

	//if we just do a normal replace to swap text for the actual number, edge cases like "eightwo"
	//will result in only one of the two digits getting converted. To prevent this we leave the first
	//and last characters on either side of the replacements. So replacing eight in the above string
	//will produce e8two
	for i, v := range numbersText {
		re := regexp.MustCompile(v)
		s := fmt.Sprintf("%v%v%v", v[0:1], (i + 1), v[len(v)-1:])
		inputString = re.ReplaceAllString(inputString, s)
	}

	return []byte(inputString)

}

func main() {

	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	sumPart1 := getSumOfCalValues(input)

	formattedInput := replaceText(input)
	sumPart2 := getSumOfCalValues(formattedInput)

	fmt.Printf("Sum of calibration values in part 1: %v\n", sumPart1)
	fmt.Printf("Sum of calibration values in part 2: %v\n", sumPart2)
}
