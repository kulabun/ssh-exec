package main

import (
	"fmt"
	"github.com/kulabun/ssh-exec/commands"
)

func main() {
  err := commands.Execute()
  if err != nil {
    fmt.Println(err)
  }
}
