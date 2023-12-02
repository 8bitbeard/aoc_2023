package main

import (
  "fmt"
  "log"
  "os"
  "regexp"
  "bufio"
  "strconv"
)

func main() {
  content, error := os.Open("input.txt")
  if error != nil {
    log.Fatal(error)
  }
  defer content.Close()

  scanner := bufio.NewScanner(content)

  re := regexp.MustCompile(`\d`)
  all_sum := 0
  for scanner.Scan() {
    line := scanner.Text()
    if len(line) > 0 {
      submatchall := re.FindAllString(line, -1)
      fmt.Print(submatchall)
      concat := submatchall[0] + submatchall[len(submatchall)-1]
      number, err_num := strconv.Atoi(concat)
      if err_num != nil {
          log.Fatal(err_num)
      }

      all_sum += number
      fmt.Println(number)
    }
  }

  fmt.Println(all_sum)
}
