package main

import (
	"strings"
	"testing"

	"github.com/andrei-m/aoc/advent"
	"github.com/stretchr/testify/assert"
)

/*
func Test_debug(t *testing.T) {
	f, err := os.Open("/path/to/input.txt")
	require.NoError(t, err)
	defer f.Close()

	m, start, end := mustParseInput(f)
	adjacencts = advent.AdjacentsFn(len(m[0]), len(m))
	pathGraph := getPathGraph(m, start)
	path := getShortestPaths(pathGraph, start, advent.Right)
	printPath(path, end)
}
*/

func Test_part1(t *testing.T) {
	t.Run("edge case 1", func(t *testing.T) {
		raw := `###########################
#######################..E#
######################..#.#
#####################..##.#
####################..###.#
###################..##...#
##################..###.###
#################..####...#
################..#######.#
###############..##.......#
##############..###.#######
#############..####.......#
############..###########.#
###########..##...........#
##########..###.###########
#########..####...........#
########..###############.#
#######..##...............#
######..###.###############
#####..####...............#
####..###################.#
###..##...................#
##..###.###################
#..####...................#
#.#######################.#
#S........................#
###########################
`
		sb := strings.NewReader(raw)
		m, start, end := mustParseInput(sb)
		adjacencts = advent.AdjacentsFn(len(m[0]), len(m))
		pathGraph := getPathGraph(m, start)
		path := getShortestPaths(pathGraph, start, advent.Right)
		assert.Equal(t, 21148, path[end].cost)
	})

	t.Run("edge case 2", func(t *testing.T) {
		raw := `####################################################
#......................................#..........E#
#......................................#...........#
#....................#.................#...........#
#....................#.................#...........#
#....................#.................#...........#
#....................#.................#...........#
#....................#.................#...........#
#....................#.................#...........#
#....................#.................#...........#
#....................#.................#...........#
#....................#.............................#
#S...................#.............................#
####################################################
`
		sb := strings.NewReader(raw)
		m, start, end := mustParseInput(sb)
		adjacencts = advent.AdjacentsFn(len(m[0]), len(m))
		pathGraph := getPathGraph(m, start)
		path := getShortestPaths(pathGraph, start, advent.Right)
		printMap(m, path, start, end)

		assert.Equal(t, 1021, path[advent.Point{X: 20, Y: 10}].cost, "this was previously incorrectly 3023 by not considering the full path traversal cost in path comparison")
		assert.Equal(t, 5078, path[end].cost)
	})
}
