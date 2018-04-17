package main

// Create a program which parses a query to do distance conversions. For example, from a terminal:
// $ distance_converter 50mi km
// Should produce:
// 80.47km
// It should support miles (mi), kilometers (km), feet (ft) and meters (m).

import (
  "fmt"
  "strings"
)

// use a map to avoid a lengthy switch() where each condition would need
// to be specified.
var conversionMap = map[string]float64 {
  "mi2km": 1.60934,
  "km2mi": 0.621371192,
  "mi2ft": 5280,
  "km2m": 1000,
  "mi2m": 1609.34,
  "ft2mi": 0.000189394,
  "m2km": .001,
  "ft2m": .3048,
  "m2ft": 3.28084,
  "ft2km": 0.0003048,
  "km2ft": 3280.84,
  "m2mi": 0.000621371,
}

// conversion to os.Args would need split or regexp to break apart
// 1st user input arg: for example, 50mi -> 50, mi

func main() {
  var quantity float64
  var units_from string
  var units_to string

  fmt.Scanf("%f%s %s", &quantity, &units_from, &units_to)

  units_from = strings.ToLower(units_from)
  units_to = strings.ToLower(units_to)

  // also catching when both are zero value because of bad scan
  if units_from == units_to {
    panic("Units are identical or not specified.")
  } else if quantity == 0 {
    panic("No quantity specified.")
  }

  var converterkey string
  converterkey = units_from + "2" + units_to

  var converter float64
  converter, ok := conversionMap[converterkey]
  if ok == true {
    fmt.Printf("%-.2f%s\n", quantity * converter, units_to)
  } else {
    fmt.Println("Units supported: mi km ft m")
  }
}
