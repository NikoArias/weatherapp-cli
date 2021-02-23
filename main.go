package main

import(
  "fmt"
  "log"
  "os"
  "os/signal"
  "syscall"

  "github.com/nikoarias/mathpal-cli/internal/controllers"
  "github.com/nikoarias/mathpal-cli/internal/utils/simpledb"

)

func main() {
  fmt.Println("Welcome")
  done := make(chan os.Signal, 1)
  signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

  db, err := simpledb.ConnectDB("croplog_db")
  if err != nil {
    log.Fatal(err)
  }

  app := controllers.NewController(db)

  go app.RunMainTimeLoop()

  <-done

  app.StopGracefullShutDown()
}
