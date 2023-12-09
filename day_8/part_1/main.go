package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	// "sort"
	// "strconv"
	// "strings"
)

func main() {
	content, open_error := os.Open("input.txt")
	// content, open_error := os.Open("example_input.txt")
	if open_error != nil {
		log.Fatal(open_error)
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	scanner_idx := 0
	instructions := []string{}
  network := make(map[string]map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		if scanner_idx == 0 {
			instructions = strings.Split(line, "")
    } else if line != "" {
      r := regexp.MustCompile(`([A-Z]+)`)
      matches := r.FindAllString(line, -1)
      temp_map := make(map[string]string)
      temp_map["L"] = matches[1]
      temp_map["R"] = matches[2]
      network[matches[0]] = temp_map
    }
		scanner_idx++
	}

	fmt.Println(network)

  steps := 0
  current_node := "AAA"
  for current_node != "ZZZ" {
    instruction_pos := steps % len(instructions)
    current_node = network[current_node][instructions[instruction_pos]]
    steps++
    fmt.Println(instruction_pos)
    fmt.Println(current_node)
  }
	fmt.Println(steps)
}
