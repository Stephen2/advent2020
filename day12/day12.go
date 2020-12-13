package day12

import (
	"regexp"
	"strconv"
)

type Direction int

const (
	north Direction = 0
	east  Direction = 90
	south Direction = 180
	west  Direction = 270
)

func Abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func turnShip(shipFacing Direction, action string, value int) Direction {
	newDirection := int(shipFacing)

	switch action {
	case "L":
		newDirection = (360 + newDirection - value) % 360
	case "R":
		newDirection = (newDirection + value) % 360
	}

	return Direction(newDirection)
}

func moveShip(ns int, ew int, direction Direction, value int) (int, int) {
	switch direction {
	case north:
		return ns + value, ew
	case east:
		return ns, ew + value
	case south:
		return ns - value, ew
	case west:
		return ns, ew - value
	}

	panic("NOOooo")
}

func Solve(rawDirections []string) int {
	shipFacing := east
	ns := 0
	ew := 0
	directionRegex := regexp.MustCompile(`(.)(\d+)`)

	for _, direction := range rawDirections {
		matches := directionRegex.FindStringSubmatch(direction)
		action := matches[1]
		value, _ := strconv.Atoi(matches[2])

		switch action {
		case "N":
			ns, ew = moveShip(ns, ew, north, value)
		case "S":
			ns, ew = moveShip(ns, ew, south, value)
		case "E":
			ns, ew = moveShip(ns, ew, east, value)
		case "W":
			ns, ew = moveShip(ns, ew, west, value)
		case "L", "R":
			shipFacing = turnShip(shipFacing, action, value)
		case "F":
			ns, ew = moveShip(ns, ew, shipFacing, value)
		}
	}

	return Abs(ns) + Abs(ew)
}
