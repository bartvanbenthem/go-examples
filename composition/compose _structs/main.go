// Composition: compose structs
package main

import "fmt"

// you can group related fields together with structures and composition. The
// following listing defines a report structure composed of structures for temperature and location.
// Structs inside of structs:
type report struct {
	sol         int
	temperature temperature
	location    location
}
type temperature struct {
	high, low celsius
}
type location struct {
	lat, long float64
}
type celsius float64

// you can further organize your code by hanging methods from each type.
// For example, to calculate the average temperature, you can write a method like the one shown in the next listing.
func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func main() {
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}

	// The temperature type and average method can be used independently of the weather report as follows:
	fmt.Printf("average %vº C\n", t.average())

	// When you create a weather report, the average method is accessible by chaining off the temperature field:
	report := report{sol: 15, temperature: t, location: bradbury}
	fmt.Printf("average %vº C\n", report.temperature.average())
}
