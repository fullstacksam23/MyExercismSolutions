package cars
// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	carsPerHour := float64(productionRate)*successRate/100.0
    return carsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsPerMinute := int(CalculateWorkingCarsPerHour(productionRate, successRate))/60
    return carsPerMinute
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groups := carsCount/10
    rem := carsCount%10
    return uint(groups*95000+rem*10000)
}
