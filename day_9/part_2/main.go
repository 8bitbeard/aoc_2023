package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	content, open_error := os.Open("input.txt")
	// content, open_error := os.Open("example_input.txt")
	if open_error != nil {
		log.Fatal(open_error)
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	all_sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			numbers := extractNumbers(line)
			firstDigits := []int{numbers[0]}
			iterNumbers := numbers
			for !isAllZeros(iterNumbers) {
				fmt.Println(iterNumbers)
				iterNumbers = buildDiffSlice(iterNumbers)
				firstDigit := []int{iterNumbers[0]}
				firstDigits = append(firstDigit, firstDigits...)
			}
      // println(firstDigits)
			next_value := 0
			for _, num := range firstDigits {
        next_value = num - next_value
        // println(next_value)
			}
			fmt.Println(next_value)
			all_sum += next_value
		}
	}
	println(all_sum)
}

func extractNumbers(s string) []int {
	stringNumbers := strings.Split(s, " ")
	intNumbers := []int{}
	for _, stringNumber := range stringNumbers {
		intNumber, err := strconv.Atoi(stringNumber)
		if err != nil {
			log.Fatal(err)
		}
		intNumbers = append(intNumbers, intNumber)
	}
	return intNumbers
}

func buildDiffSlice(numbers []int) []int {
	diffSlice := []int{}
	for i := 1; i < len(numbers); i++ {
		diffSlice = append(diffSlice, numbers[i]-numbers[i-1])
	}
	return diffSlice
}

func isAllZeros(slice []int) bool {
	zeroSlice := make([]int, len(slice))
	return reflect.DeepEqual(zeroSlice, slice)
}
