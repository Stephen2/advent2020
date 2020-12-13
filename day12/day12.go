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

func moveShipToWaypoint(ns int, ew int, waypointNs int, waypointEw int, value int) (int, int) {
	newNs := ns + (waypointNs * value)
	newEw := ew + (waypointEw * value)

	return newNs, newEw
}

func SolvePart2(rawDirections []string) int {
	ns := 0
	ew := 0
	waypointNs := 1
	waypointEw := 10
	directionRegex := regexp.MustCompile(`(.)(\d+)`)

	for _, direction := range rawDirections {
		matches := directionRegex.FindStringSubmatch(direction)
		action := matches[1]
		value, _ := strconv.Atoi(matches[2])

		switch action {
		case "N":
			waypointNs, waypointEw = moveShip(waypointNs, waypointEw, north, value)
		case "S":
			waypointNs, waypointEw = moveShip(waypointNs, waypointEw, south, value)
		case "E":
			waypointNs, waypointEw = moveShip(waypointNs, waypointEw, east, value)
		case "W":
			waypointNs, waypointEw = moveShip(waypointNs, waypointEw, west, value)
		case "L":
			newWaypointNs, newWaypointEw := waypointNs, waypointEw

			switch value {
			case 90:
				newWaypointNs = waypointEw
				newWaypointEw = -waypointNs
			case 180:
				newWaypointNs = -waypointNs
				newWaypointEw = -waypointEw
			case 270:
				newWaypointNs = -waypointEw
				newWaypointEw = waypointNs
			}

			waypointNs = newWaypointNs
			waypointEw = newWaypointEw

		case "R":
			newWaypointNs, newWaypointEw := waypointNs, waypointEw

			switch value {
			case 90:
				newWaypointNs = -waypointEw
				newWaypointEw = waypointNs
			case 180:
				newWaypointNs = -waypointNs
				newWaypointEw = -waypointEw
			case 270:
				newWaypointNs = waypointEw
				newWaypointEw = -waypointNs
			}

			waypointNs = newWaypointNs
			waypointEw = newWaypointEw

		case "F":
			ns, ew = moveShipToWaypoint(ns, ew, waypointNs, waypointEw, value)
		}
	}

	return Abs(ns) + Abs(ew)
}
