package main

import (
	"errors"
	"strconv"
)

type Point struct {
	x     int
	y     int
	steps int
}

type Line struct {
	a Point
	b Point
}

func CalculateMinManhattanDistance(wire1, wire2 []string) int {
	intersections := calculateAllIntersections(wire1, wire2)

	minManhattan := CalculateManhattanDistance(intersections[0])

	for i := range intersections {
		if CalculateManhattanDistance(intersections[i]) < minManhattan {
			minManhattan = CalculateManhattanDistance(intersections[i])
		}
	}

	return minManhattan
}

func CalculateMinSteps(wire1, wire2 []string) int {
	intersections := calculateAllIntersections(wire1, wire2)

	minSteps := intersections[0].steps

	for i := range intersections {
		if intersections[i].steps < minSteps {
			minSteps = intersections[i].steps
		}
	}

	return minSteps
}

func calculateAllIntersections(wire1, wire2 []string) (intersections []Point) {
	lines1 := calculateLines(wire1)
	lines2 := calculateLines(wire2)

	for i := range lines1 {
		for j := range lines2 {
			intersection, err := calculateIntersection(lines1[i], lines2[j])

			if err == nil {
				intersections = append(intersections, intersection)
			}
		}
	}

	if intersections != nil && intersections[0].x == 0 && intersections[0].y == 0 {
		return intersections[1:]
	} else {
		return intersections
	}
}

func calculateLines(input []string) (lines []Line) {
	var point = Point{0, 0, 0}

	for i := range input {
		newPoint, err := createPoint(input[i], point)

		if err != nil {
			panic(err)
		}

		lines = append(lines, Line{point, newPoint})

		point = newPoint
	}

	return lines
}

func calculateIntersection(line1 Line, line2 Line) (Point, error) {
	if !Touching(line1, line2) {
		return Point{}, errors.New("lines are not touching")
	}

	a1, b1, c1 := CalculateLineFunction(line1)
	a2, b2, c2 := CalculateLineFunction(line2)

	determinant := a1*b2 - a2*b1

	if determinant == 0 {
		return Point{}, errors.New("lines are parallel")
	}

	x := (b2*c1 - b1*c2) / determinant
	y := (a1*c2 - a2*c1) / determinant

	return Point{
		x:     x,
		y:     y,
		steps: CalculateIntersectionSteps(line1, line2, x, y),
	}, nil
}

func createPoint(input string, point Point) (Point, error) {
	value, _ := strconv.Atoi(input[1:])
	direction := string(input[0])

	switch direction {
	case "R":
		return Point{point.x + value, point.y, point.steps + value}, nil
	case "U":
		return Point{point.x, point.y + value, point.steps + value}, nil
	case "L":
		return Point{point.x - value, point.y, point.steps + value}, nil
	case "D":
		return Point{point.x, point.y - value, point.steps + value}, nil
	}

	return Point{0, 0, 0}, errors.New("unable to create point")
}
