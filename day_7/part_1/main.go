package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	hand_type string
	cards     string
	bid       int
}

func (h *Hand) categorizeHand() {
	cardCount := make(map[rune]int)
	for _, char := range h.cards {
		cardCount[char]++
	}
	switch len(cardCount) {
	case 1:
		h.hand_type = "Five of a kind"
	case 2:
		for _, value := range cardCount {
			if value > 3 {
				h.hand_type = "Four of a kind"
				break
			}
		}
		if h.hand_type == "" {
			h.hand_type = "Full house"
		}
	case 3:
		for _, value := range cardCount {
			if value > 2 {
				h.hand_type = "Three of a kind"
				break
			}
		}
		if h.hand_type == "" {
			h.hand_type = "Two pair"
		}
	case 4:
		h.hand_type = "One pair"
	case 5:
		h.hand_type = "High card"
	}
}

// const alphabet = "J23456789TQKA"
const alphabet = "23456789TJQKA"

// var hand_power := [7]string{"High card", "One pair", "Two pair", "Three pair", "Full house", "Four of a kind", "Five of a kind"}
var hand_power = [7]string{"High card", "One pair", "Two pair", "Three of a kind", "Full house", "Four of a kind", "Five of a kind"}

func customLess(h1 Hand, h2 Hand, alphabet string) bool {
	s1 := strings.ToUpper(h1.cards)
	s2 := strings.ToUpper(h2.cards)

	for h1.hand_type == h2.hand_type {
		for i := 0; i < len(s1) && i < len(s2); i++ {
			idx1 := strings.IndexRune(alphabet, rune(s1[i]))
			idx2 := strings.IndexRune(alphabet, rune(s2[i]))

			if idx1 != idx2 {
				return idx1 < idx2
			}
		}
		return len(s1) < len(s2)
	}

	var idx1 int
	var idx2 int
	for i, value := range hand_power {
		if value == h1.hand_type {
			idx1 = i
		}
		if value == h2.hand_type {
			idx2 = i
		}
	}
	return idx1 < idx2
}

func main() {
  // content, open_error := os.Open("input.txt")
  content, open_error := os.Open("example_input.txt")
  if open_error != nil {
    log.Fatal(open_error)
  }
  defer content.Close()

  scanner := bufio.NewScanner(content)

  hands := []Hand{}
  for scanner.Scan() {
    line := scanner.Text()
    splittedLine := strings.Split(line, " ")
    cards := splittedLine[0]
    bid, _ := strconv.Atoi(splittedLine[1])
    hand := Hand{cards: cards, bid: bid}
    hand.categorizeHand()
    hands = append(hands, hand)
  }

  sort.Slice(hands, func(i, j int) bool {
    return customLess(hands[i], hands[j], alphabet)
  })

  all_sum := 0
  for i, hand := range hands {
    all_sum += hand.bid * (i + 1)
  }

  fmt.Println(hands)
  fmt.Println(all_sum)
}
