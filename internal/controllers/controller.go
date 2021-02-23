package controllers

import (
  "fmt"
  "github.com/sdomino/scribble"
)

type Controller struct {
    db *scribble.Driver
}

func NewController(db *scribble.Driver) (*Controller){
  return &Controller{
    db: db,
  }
}

func (c *Controller)RunMainTimeLoop() {
  fmt.Println("Running...")
}

func (c *Controller)StopGracefullShutDown() {
  fmt.Println("GoodBye")
}
