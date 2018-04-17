package main

import (
  "fmt"
  "strings"
  "strconv"
)

func main() {
  fmt.Println("hello, world")
  // doesn't seem to be any difference between ToTitle and ToUpper for English
  fmt.Println(strings.ToUpper("hello world"))
  // Title is probably what will be used in most exercises
  // e.g. "capitalize each word"
  fmt.Println(strings.Title("hello world"))

  fmt.Println(strconv.Itoa(2324))
  s := "988323"
  // ignore error handling in simple example
  val, _ := strconv.Atoi(s)
  fmt.Println(600 + val)
}
