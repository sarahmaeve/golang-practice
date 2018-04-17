package main

// Create a program which parses a query to do distance conversions. For example, from a terminal:
// $ distance_converter 50mi km
// Should produce:
// 80.47km
// It should support miles (mi), kilometers (km), feet (ft) and meters (m).

import (
  "fmt"
  "os"
  "strings"
)

// Probably could be expressed better than a zillion constants

// next step -- use os.Args

const MilestoKm float64 = 1.60934
const KmtoMiles float64 = 0.621371192
const MilestoFeet float64 = 5280
const KmtoMeters float64 = 1000
const MilestoMeters float64 = MilestoKm * KmtoMeters
const KmtoFeet float64 = KmtoMiles * MilestoFeet
const FeettoMiles float64 = 1 / MilestoFeet
const MeterstoKm float64 = 1 / KmtoMeters
const FeettoMeters float64 = .3048
const MeterstoFeet float64 = 1 / FeettoMeters

func main() {
  var quantity float64
  var units_from string
  var units_to string

  fmt.Scanf("%f%s %s", &quantity, &units_from, &units_to)

  units_from = strings.ToLower(units_from)
  units_to = strings.ToLower(units_to)

  // also catching when both are zero value because of bad scan
  if units_from == units_to {
    fmt.Println("Units are identical.")
    os.Exit(1)
  }

  var converter float64
  switch {
    // add more later
  case units_from == "mi" && units_to == "km": converter = MilestoKm
  case units_from == "km" && units_to == "mi": converter = KmtoMiles
   }

  if converter != 0 {
    fmt.Printf("%-.2f%s\n", quantity * converter, units_to)
  }

}
