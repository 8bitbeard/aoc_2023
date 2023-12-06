package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start_value int
	end_value   int
}

func (r Range) is_in_range(value int) bool {
	if value >= r.start_value && value <= r.end_value {
		return true
	}
	return false
}

type Mapping struct {
	configs [][]int
}

func (m *Mapping) append_config(config []int) {
	m.configs = append(m.configs, config)
}

func (m Mapping) get_input_from_output(output int) int {
	for _, config := range m.configs {
		s := config[1]
		d := config[0]
		r := config[2]

		if output >= d && output <= d+r {
			return s + (output - d)
		}
	}
	return output
}

func main() {
	content, error := os.Open("input.txt")
	// content, error := os.Open("test_input.txt")
	if error != nil {
		log.Fatal(error)
	}
	defer content.Close()

	var (
		seeds                   []Range
		seed_to_soil            Mapping
		soil_to_fertilizer      Mapping
		fertilizer_to_water     Mapping
		water_to_light          Mapping
		light_to_temperature    Mapping
		temperature_to_humidity Mapping
		humidity_to_location    Mapping
	)

	scanner := bufio.NewScanner(content)

	var mapping *Mapping
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "seeds:") {
			values := strings.Split(strings.Split(line, ": ")[1], " ")
			for i := 0; i < len(values); i += 2 {
				int_start, _ := strconv.Atoi(values[i])
				int_range, _ := strconv.Atoi(values[i+1])
				seeds = append(seeds, Range{start_value: int_start, end_value: int_start + int_range})
			}
		} else if strings.Contains(line, "seed-to-soil map:") {
			mapping = &seed_to_soil
			continue
		} else if strings.Contains(line, "soil-to-fertilizer map") {
			mapping = &soil_to_fertilizer
			continue
		} else if strings.Contains(line, "fertilizer-to-water map:") {
			mapping = &fertilizer_to_water
			continue
		} else if strings.Contains(line, "water-to-light map:") {
			mapping = &water_to_light
			continue
		} else if strings.Contains(line, "light-to-temperature map:") {
			mapping = &light_to_temperature
			continue
		} else if strings.Contains(line, "temperature-to-humidity map:") {
			mapping = &temperature_to_humidity
			continue
		} else if strings.Contains(line, "humidity-to-location map:") {
			mapping = &humidity_to_location
			continue
		} else if line == "" {
			continue
		} else {
			values := strings.Split(line, " ")
			d, _ := strconv.Atoi(values[0])
			s, _ := strconv.Atoi(values[1])
			r, _ := strconv.Atoi(values[2])
			config := []int{d, s, r}
			fmt.Println(config)

			mapping.append_config(config)
		}
	}

	location := 0
	for {
		humidity := humidity_to_location.get_input_from_output(location)
		temperature := temperature_to_humidity.get_input_from_output(humidity)
		light := light_to_temperature.get_input_from_output(temperature)
		water := water_to_light.get_input_from_output(light)
		fertilizer := fertilizer_to_water.get_input_from_output(water)
		soil := soil_to_fertilizer.get_input_from_output(fertilizer)
		seed := seed_to_soil.get_input_from_output(soil)

		for _, r := range seeds {
			if r.is_in_range(seed) {
				fmt.Println(location)
				return
			}
		}
		location++
	}
}
