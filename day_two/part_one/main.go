package main

import (
  "fmt"
  "log"
  "os"
  "regexp"
  "bufio"
  "strconv"
  "strings"
  // "reflect"
)

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
    game := strings.Split(strings.Split(line, ":")[0], " ")[1]
    game_id, _ := strconv.Atoi(game)
    games := strings.Split(strings.Split(line, ":")[1], ";")

    blue_re := regexp.MustCompile(`\d+\sblue`)
    red_re := regexp.MustCompile(`\d+\sred`)
    green_re := regexp.MustCompile(`\d+\sgreen`)

    is_valid := true
    for _, slice := range games {
      red := red_re.FindAllString(slice, -1)
      green := green_re.FindAllString(slice, -1)
      blue := blue_re.FindAllString(slice, -1)

      if len(red) > 0 {
        red_value, _ := strconv.Atoi(strings.Split(red[0], " ")[0])
        if red_value > 12 {
          fmt.Print("Invalid Game!! --- ")
          fmt.Println(blue, red, green)
          is_valid = false
          break
        }
      }
      if len(green) > 0 {
        green_value, _ := strconv.Atoi(strings.Split(green[0], " ")[0])
        if green_value > 13 {
          fmt.Print("Invalid Game!! --- ")
          fmt.Println(blue, red, green)
          is_valid = false
          break
        }
      }
      if len(blue) > 0 {
        blue_value, _ := strconv.Atoi(strings.Split(blue[0], " ")[0])
        if blue_value > 14 {
          fmt.Print("Invalid Game!! --- ")
          fmt.Println(blue, red, green)
          is_valid = false
          break
        }
      }
      is_valid = true
      fmt.Print(game + ": Valid Game!! --- ")
      fmt.Println(blue, red, green)
    }
    if is_valid {
      fmt.Print(strconv.Itoa(all_sum))
      fmt.Print(" + ")
      fmt.Print(strconv.Itoa(game_id))
      fmt.Print(" = ")
      fmt.Println(all_sum + game_id)
      all_sum += game_id
    }
  }
  fmt.Println(all_sum)
}
