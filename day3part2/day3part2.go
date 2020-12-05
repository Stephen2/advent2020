package day3part2

//Solve solves day3part2
func Solve(mapLines []string, numRight int, numDown int) int {
	xLength := len(mapLines[0])
	yLength := len(mapLines)
	x := 0
	treesHit := 0

	for y := 0; y < yLength; y += numDown {
		if string(mapLines[y][x]) == "#" {
			treesHit++
		}

		x += numRight
		x = x % xLength
	}

	return treesHit
}
