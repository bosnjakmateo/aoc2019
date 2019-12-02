package main

func FindNounAndVerb(integers []int) int {
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			newState := append([]int{}, integers...)

			newState[1] = i
			newState[2] = j

			state := CalculatePreviousState(newState)

			if state == 19690720 {
				return 100*i + j
			}
		}
	}

	return 100*0 + 0
}
