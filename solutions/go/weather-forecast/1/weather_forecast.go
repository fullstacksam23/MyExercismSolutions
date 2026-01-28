// Package weather provides the tools to forecast the weather at a location.
package weather

var (
    // CurrentCondition represents the weather contdition currently.
	CurrentCondition string
    // CurrentLocation represents the current location where we're figuring out the weather.
	CurrentLocation  string
)
//Forecast takes city and condition as params and returns the weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
