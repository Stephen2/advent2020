package day11

type ChairStatus int

const (
	floor    ChairStatus = 0
	empty    ChairStatus = 1
	occupied ChairStatus = 2
)

func getChairStatus(rawSeats []string) [][]ChairStatus {
	chairStatus := make([][]ChairStatus, len(rawSeats))

	for row, line := range rawSeats {
		chairStatus[row] = make([]ChairStatus, len(line))

		for col, seat := range line {
			status := floor

			switch seat {
			case 'L':
				status = empty
			case '#':
				status = occupied
			}

			chairStatus[row][col] = status
		}
	}

	return chairStatus
}

func processMovements(chairStatus [][]ChairStatus) (newChairStatus [][]ChairStatus, changeCount int) {
	newChairStatus = make([][]ChairStatus, len(chairStatus))

	for row, line := range chairStatus {
		newChairStatus[row] = make([]ChairStatus, len(line))

		for col, seat := range line {
			if seat == floor {
				newChairStatus[row][col] = floor
				continue
			}

			occupiedSeats := 0

			for rowMod := -1; rowMod <= 1; rowMod++ {
				for colMod := -1; colMod <= 1; colMod++ {
					rowToCheck := row + rowMod
					colToCheck := col + colMod

					// Don't check your own position dumbo
					if rowToCheck == row && colToCheck == col {
						continue
					}

					if rowToCheck >= 0 &&
						rowToCheck < len(chairStatus) &&
						colToCheck >= 0 &&
						colToCheck < len(line) &&
						chairStatus[rowToCheck][colToCheck] == occupied {
						occupiedSeats++
					}
				}
			}

			if seat == empty && occupiedSeats == 0 {
				newChairStatus[row][col] = occupied
				changeCount++
			} else if seat == occupied && occupiedSeats >= 4 {
				newChairStatus[row][col] = empty
				changeCount++
			} else {
				newChairStatus[row][col] = seat
			}
		}
	}

	return newChairStatus, changeCount
}

type Direction struct {
	rowMod int
	colMod int
}

var nw = Direction{rowMod: -1, colMod: -1}
var n = Direction{rowMod: -1, colMod: 0}
var ne = Direction{rowMod: -1, colMod: 1}
var e = Direction{rowMod: 0, colMod: 1}
var w = Direction{rowMod: 0, colMod: -1}
var sw = Direction{rowMod: 1, colMod: -1}
var s = Direction{rowMod: 1, colMod: 0}
var se = Direction{rowMod: 1, colMod: 1}

var allDirections = []Direction{nw, n, ne, e, w, sw, s, se}

func processMovementsBasedOnLOS(chairStatus [][]ChairStatus) (newChairStatus [][]ChairStatus, changeCount int) {
	newChairStatus = make([][]ChairStatus, len(chairStatus))

	for row, line := range chairStatus {
		newChairStatus[row] = make([]ChairStatus, len(line))

		for col, seat := range line {
			if seat == floor {
				newChairStatus[row][col] = floor
				continue
			}

			occupiedSeats := 0

			for _, direction := range allDirections {
				rowToCheck := row + direction.rowMod
				colToCheck := col + direction.colMod

				for {
					if rowToCheck < 0 || rowToCheck >= len(chairStatus) ||
						colToCheck < 0 || colToCheck >= len(line) ||
						chairStatus[rowToCheck][colToCheck] == empty {
						break
					}

					if chairStatus[rowToCheck][colToCheck] == occupied {
						occupiedSeats++
						break
					}

					rowToCheck = rowToCheck + direction.rowMod
					colToCheck = colToCheck + direction.colMod
					continue
				}
			}

			if seat == empty && occupiedSeats == 0 {
				newChairStatus[row][col] = occupied
				changeCount++
			} else if seat == occupied && occupiedSeats >= 5 {
				newChairStatus[row][col] = empty
				changeCount++
			} else {
				newChairStatus[row][col] = seat
			}
		}
	}

	return newChairStatus, changeCount
}

func countOccupiedSeats(chairStatus [][]ChairStatus) (occupiedSeats int) {
	for _, line := range chairStatus {
		for _, seat := range line {
			if seat == occupied {
				occupiedSeats++
			}
		}
	}

	return occupiedSeats
}

func Solve(rawSeats []string) int {
	chairStatus := getChairStatus(rawSeats)
	processes := 0

	for {
		newChairStatus, changeCount := processMovements(chairStatus)
		chairStatus = newChairStatus

		if changeCount == 0 {
			break
		}

		processes++
	}

	println("Processed:", processes, "times")
	return countOccupiedSeats(chairStatus)
}

func SolvePart2(rawSeats []string) int {
	chairStatus := getChairStatus(rawSeats)
	processes := 0

	for {
		newChairStatus, changeCount := processMovementsBasedOnLOS(chairStatus)
		chairStatus = newChairStatus

		if changeCount == 0 {
			break
		}

		processes++
	}

	println("Processed:", processes, "times")
	return countOccupiedSeats(chairStatus)
}
