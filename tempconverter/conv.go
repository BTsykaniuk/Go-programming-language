package tempconverter


// Convert temperature from Celsius to Fahrenheit
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}


// Convert temperature from Celsius to Kelvin
func CtoK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}


// Convert Fahrenheit to Celsius
func FtoC(f Fahrenheit) Celsius {
	return Celsius((f+32)*5/9)
}
