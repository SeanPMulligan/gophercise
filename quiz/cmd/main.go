package main

import(
  "bufio"
  "encoding/csv"
  "io"
  "os"
  "log"
  "fmt"
  "time"
  "flag"
)

func main() {
  fmt.Print("It ENTER to begin quiz")
  fmt.Scanf("%s\n")

  timer := setTimer()

  csvFile, _ := os.Open("../problems.csv")
  reader := csv.NewReader(bufio.NewReader(csvFile))
  var correctCount int

  go func() {
    <-timer.C
    fmt.Println("\nYour total score is", correctCount, "/", 13)
    os.Exit(0)
  }()

  runQuiz(reader, &correctCount)
  defer fmt.Println("Your total score is", correctCount, "/", 13)
}

func runQuiz(reader *csv.Reader, correctCount *int) {
  for {
    line, err := reader.Read()

    if err == io.EOF {
      break
    } else if err != nil {
      log.Fatal(err)
    }

    fmt.Println(line[0])

    fmt.Print("-> ")
    var answer string
    fmt.Scanf("%s\n", &answer)

    if answer == line[1] {
      *correctCount++
    }
  }
}

func setTimer() *time.Timer {
  var timePtr = flag.Int("timer", 30, "quiz timer")
  flag.Parse()

 return time.NewTimer(time.Duration(*timePtr) * time.Second)
}
