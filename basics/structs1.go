package main

import "fmt"

// code along with https://www.youtube.com/watch?v=3fsvqo9pQyg
// and https://www.youtube.com/watch?v=i3o4G4bmqPc

const MilestoKm float64 = 1.60934
const usixteenbitmax float64 = 65535

type Car struct {
	gas_pedal      uint16
	brake_pedal    uint16
	steering_wheel int16
	top_speed_kmh  float64
}

func (c Car) kph() float64 {
	return float64(c.gas_pedal) / usixteenbitmax * c.top_speed_kmh
}

func (c Car) mph() float64 {
	return c.kph() * MilestoKm
}

// pointer receiver used when modifying a field in a struct.
func (c *Car) upgrade(new_speed float64) {
	c.top_speed_kmh = new_speed
}

func main() {
	a_car := Car{11500, 0, 1748, 205.0}
	b_car := Car{
		gas_pedal:      634,
		brake_pedal:    2000,
		steering_wheel: 236,
		top_speed_kmh:  250.0,
	}

	fmt.Println("Kph:", a_car.kph())
	fmt.Println("Mph:", b_car.mph())
	a_car.upgrade(275.0)
	fmt.Println("Kph:", a_car.kph())
}
