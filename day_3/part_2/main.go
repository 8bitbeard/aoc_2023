package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type EngineSchematic struct {
	cells   [][]string
	numbers []Number
	symbols []Symbol
}

func (es *EngineSchematic) computeNumberNeighbors() {
	new_numbers := []Number{}
	for _, number := range es.numbers {
		start_point := number.coord[0]
		end_point := number.coord[1]
		for x := start_point.x - 1; x <= end_point.x+1; x++ {
			for y := start_point.y - 1; y <= end_point.y+1; y++ {
				if x >= 0 && x < len(es.cells[0]) && y >= 0 && y < len(es.cells) {
					if x < start_point.x || x > end_point.x || y != start_point.y {
						number.neighbors = append(number.neighbors, Point{x: x, y: y})
					}
				}
			}
		}
		new_numbers = append(new_numbers, number)
	}
	es.numbers = new_numbers
}

func (es *EngineSchematic) computeSymbolNumberNeighbors() {
	new_symbols := []Symbol{}
	for _, symbol := range es.symbols {
		for _, number := range es.numbers {
			for _, point := range number.neighbors {
				if symbol.coord == point {
          symbol.neighbors = append(symbol.neighbors, number)
				}
			}
		}
    new_symbols = append(new_symbols, symbol)
	}
  es.symbols = new_symbols
}

type Number struct {
	coord     []Point
	value     int
	neighbors []Point
}

func (n Number) hasNeighbourSymbol(symbols []Symbol) bool {
	for _, neighbour_point := range n.neighbors {
		for _, symbol := range symbols {
			if symbol.coord == neighbour_point {
				return true
			}
		}
	}
	return false
}

type Symbol struct {
	coord     Point
	value     string
	neighbors []Number
}

type Point struct {
	x int
	y int
}

func get_numbers_and_symbols(chars []string, line_num int) ([]Number, []Symbol) {
	numbers := []Number{}
	symbols := []Symbol{}
	chars = append(chars, ".")

	end_idx := 0
	buff := ""
	for char_idx, char := range chars {
		if _, err := strconv.Atoi(char); err == nil {
			// fmt.Println(char)
			buff += char
			end_idx = char_idx
		} else {
			if char == "*" {
				symbol_coord := Point{x: char_idx, y: line_num}
				symbol := Symbol{coord: symbol_coord, value: char}
				symbols = append(symbols, symbol)
			}
			if buff != "" {
				start_idx := end_idx - len(buff) + 1
				number_start_coord := Point{x: start_idx, y: line_num}
				number_end_coord := Point{x: end_idx, y: line_num}
				value, _ := strconv.Atoi(buff)
				number := Number{coord: []Point{number_start_coord, number_end_coord}, value: value}
				numbers = append(numbers, number)
			}
			buff = ""
			end_idx = char_idx
		}
	}

	return numbers, symbols

}

func main() {
	content, error := os.Open("../input.txt")
	// content, error := os.Open("test_input.txt")
	if error != nil {
		log.Fatal(error)
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	scanner_index := 0
	engine_schematic := &EngineSchematic{}
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		numbers, symbols := get_numbers_and_symbols(chars, scanner_index)
		engine_schematic.cells = append(engine_schematic.cells, chars)
		engine_schematic.numbers = append(engine_schematic.numbers, numbers...)
		engine_schematic.symbols = append(engine_schematic.symbols, symbols...)
		scanner_index++

		// fmt.Println(engine_schematic)
		// slice = append(slice, split)
		// fmt.Println(slice[0][1])
	}
	engine_schematic.computeNumberNeighbors()
	engine_schematic.computeSymbolNumberNeighbors()

	// for _, symbol := range engine_schematic.symbols {
	//   fmt.Println(engine_schematic.cells[symbol.coord.y][symbol.coord.x])
	// }

	all_sum := 0
	for _, symbol := range engine_schematic.symbols {
    if len(symbol.neighbors) == 2 {
      fmt.Println(symbol)
      all_sum += symbol.neighbors[0].value * symbol.neighbors[1].value
    }
	}
	// fmt.Println(engine_schematic.cells)
	// fmt.Println(engine_schematic.numbers)
	// fmt.Println(engine_schematic.symbols)

	fmt.Println(all_sum)
	// fmt.Println(engine_schematic.cells[126][140])
}
