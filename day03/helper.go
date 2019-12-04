package main

import utils "aoc2019/util"

func CalculateIntersectionSteps(line1, line2 Line, x, y int) int {
	horizontal, vertical := resolveLineDirection(line1, line2)

	horizontalSteps := utils.Abs(horizontal.a.x - x)
	verticalSteps := utils.Abs(vertical.a.y - y)

	return line1.a.steps + line2.a.steps + horizontalSteps + verticalSteps
}

func CalculateManhattanDistance(point Point) int {
	return utils.Abs(point.x) + utils.Abs(point.y)
}

func CalculateLineFunction(line Line) (int, int, int) {
	pointA := line.a
	pointB := line.b

	a := pointB.y - pointA.y
	b := pointA.x - pointB.x
	c := a*pointA.x + b*pointA.y

	return a, b, c
}

func Touching(line1, line2 Line) bool {
	horizontal, vertical := resolveLineDirection(line1, line2)

	minX, maxX := utils.MinMax(horizontal.a.x, horizontal.b.x)
	minY, maxY := utils.MinMax(vertical.a.y, vertical.b.y)

	x := vertical.a.x
	y := horizontal.a.y

	if (minX <= x && x <= maxX) && (minY <= y && y <= maxY) {
		return true
	} else {
		return false
	}
}

func resolveLineDirection(line1 Line, line2 Line) (horizontal, vertical Line) {
	if line1.a.y == line1.b.y {
		horizontal = line1
		vertical = line2
	} else {
		horizontal = line2
		vertical = line1
	}

	return horizontal, vertical
}
