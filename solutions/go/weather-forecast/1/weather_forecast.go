// Package weather provides functionality related to weather information.
package weather

// CurrentCondition stores the current weather condition.
var CurrentCondition string 

// CurrentLocation stores the current location for weather forecasting location.
var CurrentLocation string 

// Forecast provides the current weather forecast for the given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
