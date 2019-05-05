// Simple package for converting temperature

package tempconverter

import "fmt"


type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteNullC	Celsius = -273.15
	FreezingC		Celsius = 0
	BoolingC		Celsius = 100
)


// Method for Celsius type
func (c Celsius) String() string {
	return fmt.Sprintf("%g C", c)
}


// Method for Fahrenheit type
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g F", f)
}


// Method for Kelvin type
func (k Kelvin) String() string {
	return fmt.Sprintf("%g K", k)
}


