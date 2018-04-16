package main


import (
  "fmt"
)

func main() {
  var name string
  fmt.Print("Please enter your name: ")
  fmt.Scanf("%s", &name) // pointer used as second argument
  fmt.Println("Hello,", name)
}
