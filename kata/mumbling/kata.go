package kata

import (
  "strings"
)

func Accum(s string) string {
    sl := strings.Split(s,"")
    for i, letter := range sl  {
       sl[i] = strings.Title(strings.Repeat(strings.ToLower(letter), i + 1))
    }
    return strings.Join(sl, "-")
}
