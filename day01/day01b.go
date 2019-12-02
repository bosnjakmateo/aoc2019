package main

func CalculateAllNonNegativeFuel(modules []int) (sum int) {
	for i := range modules {
		sum += CalculateFuelForFuel(modules[i])
	}

	return sum
}

func CalculateFuelForFuel(fuel int) int {
	if fuel < 6 {
		return 0
	}

	neededFuel := (fuel / 3) - 2

	return neededFuel + CalculateFuelForFuel(neededFuel)
}
