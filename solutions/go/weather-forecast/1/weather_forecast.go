// Package weather connects the application to
// external weather services to fetch meteorological data.
package weather



var (
    // CurrentCondition stores the current weather condition of the location.
	CurrentCondition string 
    // CurrentLocation stores the name of the current location.
	CurrentLocation  string 
)

// Forecast returns the current condition of the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
