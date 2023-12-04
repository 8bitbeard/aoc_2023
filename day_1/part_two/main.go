package main

import (
  "fmt"
  "log"
  "os"
  "regexp"
  "bufio"
  "strconv"
  "strings"
)

func main() {
  content, error := os.Open("../input.txt")
  // content, error := os.Open("test_input.txt")
  if error != nil {
    log.Fatal(error)
  }
  defer content.Close()

  scanner := bufio.NewScanner(content)

  re := regexp.MustCompile(`\d`)
  all_sum := 0
  for scanner.Scan() {
    line := scanner.Text()
    buff := ""
    for _, slice := range strings.Split(line, "") {
      buff += slice

      switch {
      case strings.Contains(buff, "one"):
        buff = strings.Replace(buff, "one", "1", -1)
        buff += slice
      case strings.Contains(buff, "two"):
        buff = strings.Replace(buff, "two", "2", -1)
        buff += slice
      case strings.Contains(buff, "three"):
        buff = strings.Replace(buff, "three", "3", -1)
        buff += slice
      case strings.Contains(buff, "four"):
        buff = strings.Replace(buff, "four", "4", -1)
        buff += slice
      case strings.Contains(buff, "five"):
        buff = strings.Replace(buff, "five", "5", -1)
        buff += slice
      case strings.Contains(buff, "six"):
        buff = strings.Replace(buff, "six", "6", -1)
        buff += slice
      case strings.Contains(buff, "seven"):
        buff = strings.Replace(buff, "seven", "7", -1)
        buff += slice
      case strings.Contains(buff, "eight"):
        buff = strings.Replace(buff, "eight", "8", -1)
        buff += slice
      case strings.Contains(buff, "nine"):
        buff = strings.Replace(buff, "nine", "9", -1)
        buff += slice
      }
      // buff = strings.Replace(buff, "one", "1", -1)
      // buff = strings.Replace(buff, "two", "2", -1)
      // buff = strings.Replace(buff, "three", "3", -1)
      // buff = strings.Replace(buff, "four", "4", -1)
      // buff = strings.Replace(buff, "five", "5", -1)
      // buff = strings.Replace(buff, "six", "6", -1)
      // buff = strings.Replace(buff, "seven", "7", -1)
      // buff = strings.Replace(buff, "eight", "8", -1)
      // buff = strings.Replace(buff, "nine", "9", -1)
    }
    fmt.Print(line + " ----- " + buff)
    if len(buff) > 0 {
      submatchall := re.FindAllString(buff, -1)
      fmt.Print(submatchall)
      concat := submatchall[0] + submatchall[len(submatchall)-1]
      number, err_num := strconv.Atoi(concat)
      if err_num != nil {
          log.Fatal(err_num)
      }

      fmt.Print(strconv.Itoa(all_sum) + " + " + strconv.Itoa(number) + " = ")

      all_sum += number
      fmt.Println(all_sum)
    }
  }

  fmt.Println(all_sum)
}
