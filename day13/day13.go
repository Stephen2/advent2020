package day13

import (
	"math"
	"strconv"
	"strings"
)

func parseInput(rawBusInput []string) (int, []int) {
	now, _ := strconv.Atoi(rawBusInput[0])

	rawBuses := strings.Split(rawBusInput[1], ",")
	buses := []int{}

	for _, bus := range rawBuses {
		if bus != "x" {
			busId, _ := strconv.Atoi(bus)
			buses = append(buses, busId)
		}
	}

	return now, buses
}

func Solve(rawBusInput []string) int {
	now, buses := parseInput(rawBusInput)
	busArrives := map[int]int{}

	// Naive?
	for _, bus := range buses {
		arrivalTime := bus

		for arrivalTime < now {
			arrivalTime += bus
		}

		busArrives[bus] = arrivalTime
	}

	soonestBus := math.MaxInt32
	soonestArrivalTime := math.MaxInt32

	for bus, arrivalTime := range busArrives {
		if arrivalTime < soonestArrivalTime {
			soonestBus = bus
			soonestArrivalTime = arrivalTime
		}
	}

	return soonestBus * (soonestArrivalTime - now)
}

type BusData struct {
	id         int
	minsAfterT int
}

func SolvePart2(rawBusInput []string) int {
	busData := []BusData{}

	for i, bus := range strings.Split(rawBusInput[1], ",") {
		if bus != "x" {
			busId, _ := strconv.Atoi(bus)

			thisBusData := BusData{
				id:         busId,
				minsAfterT: i,
			}

			busData = append(busData, thisBusData)
		}
	}

	// CHEATED!!
	firstBus := busData[0]
	multiplier := firstBus.id
	time := 0

	for n := 1; n < len(busData); n++ {
		bus := busData[n]

		for {
			if (time+bus.minsAfterT)%bus.id == 0 {
				multiplier *= bus.id
				break
			}
			time += multiplier
		}
	}

	return time
}
