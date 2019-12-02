package main

func CalculateModuleFuel(modules []int) (sum int) {
	for i := range modules {
		sum += (modules[i] / 3) - 2
	}

	return sum
}
