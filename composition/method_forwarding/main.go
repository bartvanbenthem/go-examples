// Composition: Method forwarding
package main

import "fmt"

// Go will do method forwarding for you with struct embedding. To embed a type
// in a structure, specify the type without a field name, as shown in the following listing.
type report struct {
	sol int
	temperature
	location
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func main() {
	//All the methods on the temperature type are automatically made accessible through the report type:
	report := report{
		sol:         15,
		location:    location{-4.5895, 137.4417},
		temperature: temperature{high: -1.0, low: -78.0},
	}
	fmt.Printf("average %vº C\n", report.average())
	// Though no field name was specified, a field still exists with the same name as the
	// embedded type. You can access the temperature field as follows:
	fmt.Printf("average %vº C\n", report.temperature.average())

	// Embedding doesn’t only forward methods. Fields of an inner structure are accessible
	// from the outer structure. In addition to report.temperature.high, you can access the high
	// temperature with report.high as follows:
	fmt.Printf("%vº C\n", report.high)
	report.high = 32
	fmt.Printf("%vº C\n", report.temperature.high)
	// As you can see, changes to the report.high field are reflected in report.temperature.high. It’s
	// just another way to access the same data.

	// You can embed any type in a structure, not just structures.
}
