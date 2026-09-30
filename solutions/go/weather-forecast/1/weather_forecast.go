// Package weather provides weather forcast data for a city.
package weather

var (
    // CurrentCondition stores the current weather condition.
	CurrentCondition string
    // CurrentLocation stores the current location.
	CurrentLocation  string
)

// Forecast accepts a location and condition and returns a formatted string describing the weather at the location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
