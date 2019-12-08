package main

func CountNumberOfOrbits(values []string) (numberOfOrbits int) {
	var orbits []Orbit

	for i := range values {
		parentName, childName := ResolveParentChild(values[i])
		orbits = append(orbits, Orbit{parentName, childName})
	}

	for i := range orbits {
		parent := orbits[i].parent

		for {
			if parent == "" {
				break
			}
			parent = GetParent(parent, orbits)
			numberOfOrbits++
		}
	}

	return numberOfOrbits
}
