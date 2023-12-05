package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
  "math"
  // "sort"
)

type Card struct {
	winning_numbers []int
	card_numbers    []int
	points          int
}

func (c *Card) computeCardPoints() {
  matches := 0.0
  for _, cn := range c.card_numbers {
    for _, wn := range c.winning_numbers {
      if cn == wn {
        matches++
        break
      }
    }
  }
  // fmt.Print(matches)
  c.points = int(math.Pow(2, matches - 1))
}

func main() {
	content, error := os.Open("../input.txt")
	// content, error := os.Open("test_input.txt")
	if error != nil {
		log.Fatal(error)
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

  all_sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(strings.Split(line, ":")[1], "|")
    winning_numbers := strings.Split(strings.Trim(numbers[0], " "), " ")
    card_numbers := strings.Split(strings.Trim(numbers[1], " "), " ")
    // fmt.Println(card_numbers)

    card := Card{}
    for _, number := range card_numbers {
      int_num, err := strconv.Atoi(number)
      if err == nil {
        card.card_numbers = append(card.card_numbers, int_num)
      }
    }
    for _, number := range winning_numbers {
      int_num, err := strconv.Atoi(number)
      if err == nil {
        card.winning_numbers = append(card.winning_numbers, int_num)
      }
    }
    // test := card.card_numbers
    // sort.Ints(test)
    // fmt.Println(test)
    card.computeCardPoints()
    // fmt.Println(card.winning_numbers)
    // fmt.Println(card.card_numbers)
    // fmt.Println(card.points)

    all_sum += card.points
    fmt.Println(all_sum)

    // fmt.Println(card)
	}
  // fmt.Println(all_sum)
}
