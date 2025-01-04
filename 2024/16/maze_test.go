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
	path := getShortestPaths(pathGraph, start)
	printPath(path, end)
}
*/

func Test_part1(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
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
		path := getShortestPaths(pathGraph, start)
		endNode := node{loc: end, dir: advent.Up}
		assert.Equal(t, 21148, path[endNode].cost)
	})

	t.Run("case 2", func(t *testing.T) {
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
		path := getShortestPaths(pathGraph, start)
		printMap(m, path, start, end)

		n := node{loc: advent.Point{X: 20, Y: 10}, dir: advent.Up}
		assert.Equal(t, 1021, path[n].cost, "this was previously incorrectly 3023 by not considering the full path traversal cost in path comparison")
		endNode := node{loc: end, dir: advent.Up}
		assert.Equal(t, 5078, path[endNode].cost)
	})

	t.Run("case 3", func(t *testing.T) {
		raw := `##########################################################################################################
#.........#.........#.........#.........#.........#.........#.........#.........#.........#.........#...E#
#.........#.........#.........#.........#.........#.........#.........#.........#.........#.........#....#
#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#
#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#
#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#
#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#....#
#....#.........#.........#.........#.........#.........#.........#.........#.........#.........#.........#
#S...#.........#.........#.........#.........#.........#.........#.........#.........#.........#.........#
##########################################################################################################
`
		sb := strings.NewReader(raw)
		m, start, end := mustParseInput(sb)
		adjacencts = advent.AdjacentsFn(len(m[0]), len(m))
		pathGraph := getPathGraph(m, start)
		path := getShortestPaths(pathGraph, start)
		printMap(m, path, start, end)
		endNode := node{loc: end, dir: advent.Up}
		assert.Equal(t, 41210, path[endNode].cost)
	})

	t.Run("case 4", func(t *testing.T) {
		raw := `##########
#.......E#
#.##.#####
#..#.....#
##.#####.#
#S.......#
##########
`
		sb := strings.NewReader(raw)
		m, start, end := mustParseInput(sb)
		adjacencts = advent.AdjacentsFn(len(m[0]), len(m))
		pathGraph := getPathGraph(m, start)
		path := getShortestPaths(pathGraph, start)
		printMap(m, path, start, end)

		n := node{loc: advent.Point{X: 4, Y: 1}, dir: advent.Right}
		assert.Equal(t, 4009, path[n].cost, "previous node should (3,1) rather than (4,2) - this is globally optimal")
		nSuboptimal := node{loc: advent.Point{X: 4, Y: 1}, dir: advent.Up}
		assert.Equal(t, 3015, path[nSuboptimal].cost, "this is optimal for (4,1), but suboptimal on the path to E")

		endNode := node{loc: end, dir: advent.Right}
		assert.Equal(t, 4013, path[endNode].cost)
	})
}
