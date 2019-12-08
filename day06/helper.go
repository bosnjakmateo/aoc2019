package main

import "strings"

type Orbit struct {
	parent, child string
}

func GetParent(value string, orbits []Orbit) string {
	for i := range orbits {
		if orbits[i].child == value {
			return orbits[i].parent
		}
	}
	return ""
}

func ResolveParentChild(value string) (parentName, childName string) {
	split := strings.Split(value, ")")
	return split[0], split[1]
}
