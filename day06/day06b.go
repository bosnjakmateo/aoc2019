package main

func CountOrbitsToSanta(values []string) (numberOfOrbits int) {
	var orbits []Orbit

	for i := range values {
		parentName, childName := ResolveParentChild(values[i])

		orbits = append(orbits, Orbit{parentName, childName})
	}

	youObjects := getObjectsToCom("YOU", orbits)
	sanObjects := getObjectsToCom("SAN", orbits)

	for i := range youObjects {
		for j := range sanObjects {
			if youObjects[i] == sanObjects[j] {
				return i + j
			}
		}
	}

	return -1
}

func getObjectsToCom(start string, orbits []Orbit) (objects []string) {
	for i := range orbits {
		if orbits[i].child != start {
			continue
		}

		parent := orbits[i].parent
		for {
			if parent == "" {
				return objects
			}
			objects = append(objects, parent)
			parent = GetParent(parent, orbits)
		}
	}

	return objects
}
