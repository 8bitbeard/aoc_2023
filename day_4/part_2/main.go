package main

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	// "math"
	"os"
	"strconv"
	"strings"
	// "sort"
)

type Card struct {
	number         int
	winning_values []int
	card_values    []int
	points         int
}

func (c *Card) computeCardPoints() {
	matches := 0
	for _, cn := range c.card_values {
		for _, wn := range c.winning_values {
			if cn == wn {
				matches++
				break
			}
		}
	}
	// fmt.Print(matches)
	c.points = matches
}

func main() {
	content, error := os.Open("../input.txt")
	// content, error := os.Open("test_input.txt")
	if error != nil {
		log.Fatal(error)
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	cards := []Card{}
	n := 199
	// n := 6
	array := make([]int, n)
	for i := range array {
		array[i] = 1
	}
	for scanner.Scan() {
		line := scanner.Text()
    re := regexp.MustCompile(`[0-9]+`)
		card_number, _ := strconv.Atoi(re.FindString(strings.Split(line, ":")[0]))
		values := strings.Split(strings.Split(line, ":")[1], "|")
		winning_values := strings.Split(strings.Trim(values[0], " "), " ")
		card_values := strings.Split(strings.Trim(values[1], " "), " ")
		// fmt.Println(card_numbers)

		card := Card{}
		card.number = card_number
		for _, value := range card_values {
			int_num, err := strconv.Atoi(value)
			if err == nil {
				card.card_values = append(card.card_values, int_num)
			}
		}
		for _, values := range winning_values {
			int_num, err := strconv.Atoi(values)
			if err == nil {
				card.winning_values = append(card.winning_values, int_num)
			}
		}
		// test := card.card_numbers
		// sort.Ints(test)
		// fmt.Println(test)
		card.computeCardPoints()

		fmt.Print(card.number)
		fmt.Print(" -- ")
		fmt.Print(card.points)
		fmt.Print(" -- ")

		for i := card.number; i < card.points+card.number; i++ {
			// fmt.Print(i)
			// fmt.Print(array[i])
      fmt.Println(array)
			array[i] += array[card.number-1]
		}
		cards = append(cards, card)

		fmt.Println(array)

		// fmt.Println(card.winning_numbers)
		// fmt.Println(card.card_numbers)
		// fmt.Println(card.points)

		// fmt.Println(card)
	}
  all_sum := 0
  for _, val := range array {
    all_sum += val
  }
	fmt.Println(all_sum)
}
