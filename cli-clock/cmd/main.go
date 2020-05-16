package main

import (
  "fmt"
  "time"
  "github.com/inancgumus/screen"
)

func main() {
  spacer := [5]string{
    " ",
    " ",
    " ",
    " ",
    " ",
  }

  separator := [5]string{
    "   ",
    " ░ ",
    "   ",
    " ░ ",
    "   ",
  }

  numbers := [10][5]string{
    {
      "███",
      "█ █",
      "█ █",
      "█ █",
      "███",
    },
    {
      "██ ",
      " █ ",
      " █ ",
      " █ ",
      "███",
    },
    {
      "███",
      "  █",
      "███",
      "█  ",
      "███",
    },
    {
      "███",
      "  █",
      "███",
      "  █",
      "███",
    },
    {
      "█ █",
      "█ █",
      "███",
      "  █",
      "  █",
    },
    {
      "███",
      "█  ",
      "███",
      "  █",
      "███",
    },
    {
      "███",
      "█  ",
      "███",
      "█ █",
      "███",
    },
    {
      "███",
      "  █",
      "  █",
      "  █",
      "  █",
    },
    {
      "███",
      "█ █",
      "███",
      "█ █",
      "███",
    },
    {
      "███",
      "█ █",
      "███",
      "  █",
      "███",
    },
  }
  screen.Clear()

  for {
    screen.MoveTopLeft()

    now := time.Now()
    hour, min, sec := now.Hour(), now.Minute(), now.Second()

    clock := [6][5]string{
      numbers[hour/10], numbers[hour%10],
      numbers[min/10], numbers[min%10],
      numbers[sec/10], numbers[sec%10],
    }

    for i := 0; i < len(spacer); i++ {
      fmt.Print(clock[0][i])
      fmt.Print(spacer[i])
      fmt.Print(clock[1][i])
      fmt.Print(spacer[i])
      fmt.Print(separator[i])

      fmt.Print(clock[2][i])
      fmt.Print(spacer[i])
      fmt.Print(clock[3][i])
      fmt.Print(spacer[i])
      fmt.Print(separator[i])

      fmt.Print(clock[4][i])
      fmt.Print(spacer[i])
      fmt.Println(clock[5][i])
    }

    time.Sleep(1 * time.Second)
  }
}
