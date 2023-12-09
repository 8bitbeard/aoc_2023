package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	content, open_error := os.Open("input.txt")
	if open_error != nil {
		log.Fatal(open_error)
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	scanner_idx := 0
	instructions := []string{}
	current_nodes := []string{}
	network := make(map[string]map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		if scanner_idx == 0 {
			instructions = strings.Split(line, "")
		} else if line != "" {
			r := regexp.MustCompile(`([A-Z]+)`)
			nodes := r.FindAllString(line, -1)
			temp_map := make(map[string]string)
			temp_map["L"] = nodes[1]
			temp_map["R"] = nodes[2]
			network[nodes[0]] = temp_map

			r2 := regexp.MustCompile(`.*A$`)
			if r2.MatchString(nodes[0]) {
				current_nodes = append(current_nodes, nodes[0])
			}
		}
		scanner_idx++
	}

	fmt.Println(instructions)
	fmt.Println(network)
	fmt.Println(current_nodes)

  result := 1
	for _, current_node := range current_nodes {
		steps := 0
		for !checkEndsWithZ(current_node) {
			instruction_pos := steps % len(instructions)
			current_node = network[current_node][instructions[instruction_pos]]
			steps++
		}
    result = LCM(result, steps)
    fmt.Println(result)
	}
}

func checkEndsWithZ(node string) bool {
	r := regexp.MustCompile(`.*Z$`)
	if !r.MatchString(node) {
		return false
	}
	return true
}

func GCD(a, b int) int {
  for b != 0 {
    t := b
    b = a % b
    a = t
  }
  return a
}

func LCM(a, b int, integers ...int) int {
  result := a * b / GCD(a, b)

  for i := 0; i < len(integers); i++ {
    result = LCM(result, integers[i])
  }

  return result
}
