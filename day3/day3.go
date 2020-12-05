package day3

//Solve solves day3
func Solve(mapLines []string, numRightPerDown int) int {
	xLength := len(mapLines[0])
	yLength := len(mapLines)
	x := 0
	treesHit := 0

	for y := 0; y < yLength; y++ {
		if string(mapLines[y][x]) == "#" {
			treesHit++
		}

		x += numRightPerDown
		x = x % xLength
	}

	return treesHit
}
