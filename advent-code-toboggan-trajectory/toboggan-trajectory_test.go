package toboggantrajectory

import (
	"testing"
)

func TestFindTreesExample(t *testing.T) {
	mapContent := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

	result := findTrees(mapContent)
	expected := 7

	if result != expected {
		t.Errorf("When getting the number of trees, got: %v, want: %v.", result, expected)
	}
}

func TestFindTreesExercise(t *testing.T) {
	mapContent := `.....#.........#...#..##....#..
.#........#...#........#.......
......#......#..#...#....#.#..#
...#.#####.#.......##.#........
...........#......#..#.....#...
#.#..#...#.#...#.##.....#.....#
....#..#....#...#.#...#.##.....
##...#..........##..######.....
.....#...#......#.............#
........##....#...##..#....#...
...#...#.........#.#..........#
..#.#.....##..........#........
##.......................#.....
#..#...##...##.#.........##....
.#....#.#####....#...#...#.....
#......#......###..#........#.#
.#....##..##.###.#.......#.....
.#..#.........##....#.#....#...
........#..................#...
.......#..#..#............#....
........#...................##.
.#......#......#.####......#...
..###.#..#..#.........#........
..#...........###..#.....#.##..
...#.##.#....#................#
#.....#.............#.#........
.#..............#.........#....
##.................#..........#
.#..#....#.###....##..#..#...#.
##........#......##.....#....##
#......#..#........#......#.#..
....#.##.#.............#...##..
.#...#...#..#............##...#
.#..#...#..#..#....##..#.#.#...
#....#...##.#.#......#........#
#..#..#...#.#.....#..##.#......
.....#..#.#..#.##.......#..###.
#......#......#...#............
.....#......#......#..#.##..#.#
......#..##..#.....#....#......
..#..#...#..#...#....###.#.#...
.................#..#..........
......#...##..#.....#...##.....
..#...............#...#.#.....#
.#....#.##.##..#.........##....
...###....##...#......#......##
....#...#.....#.........#..##..
..###.........#..#..#...#......
...##.....#.........#.......#..
.....#.................#.#.....
.#.###.#..#...#..##....#....##.
....#.....##.........#.#.......
.#.#....#..#................#..
..#.#......#......#........#...
#........#....#..#..#..#....#.#
#...........##..#....#..####...
.....#.......#.#...#.#....###..
.......#....#.......#..........
.............#.....#...........
#....#......#...#..##.#........
....#.......#.#.......#....###.
.####.#...........#.#.#...#.#..
#..##....##.#......#...........
...##...#.#.....#.....#........
...#.............#.....#...#...
...#.....#..#.....##...###..#.#
....##..#..##..#..#...#.....#..
........#...................##.
....#.......#.....#.......#....
....##.........#.#.............
......#..#........#.#...#......
.#..#...#...........#......#..#
.#....#.#........#............#
......#...................#...#
##...#.......................#.
........###.......#.......#..#.
...........##.............#....
..##...#.....#....#......#....#
................###...##...#.#.
..#.#.....#....##...#..##......
.....................#.#......#
.......#....##.#..#........##..
.##....#......#....#.........#.
#............#.........#..#.#..
....#...........#..#....#....##
.......#..#.....##.........#...
.##..........#.#.#....#..#.....
........#....##.##.#......#....
....##..##......##.....#.###...
......##.#....##.#.#....#......
..#..#..........#.....##.....##
#........#.##...#.#....#....###
........##............#........
##.##..##.#..#...##............
....#..#....#...........#....#.
..#.......#.#.......#...#......
.#..........##.....#..#...#...#
.................##.#...#...##.
##.............#......#....#...
..........#.#....#.............
...##..#.#.....#.....#.#.......
...##...##.#......#.#...#......
..#..#.....##..##..........##..
......##........##.......#....#
....#..####..#...##........#...
#.......#....#.......##.......#
........#..........#.........#.
.....#....#.........#.#.#.....#
..##.....#....#....#..#......#.
....#..#.##...#..#.....#......#
........###.........#..###...#.
.....#.......#.....#.#.#.......
...##.....#....##.....#.#.#...#
#.##....#.##.....#.#.#........#
.##..#.......#...#.#.......#...
.#..........#.............#....
.#...#...#......#..##..........
.......................#.#....#
............###....#..##.#..#..
...#.#......##....#..#.........
..#...#....#....#.#............
..#.#..###...............##....
.....##...#.....#........#..#.#
...........#......#..#...#.##.#
#...##......##...#..#...#..#...
..##....#............#......#.#
.#.#..#...#...#.#...#...##..##.
..#.#....#.......#.#.#.#.#.##..
....###.##..#...##....#........
.#...............#........#....
...#..#........##...#.##.......
........#..#..#......##........
##....#....#............#......
#....#...#.###.#.###.......#...
...#.###.##....#.........#...##
..#......##.#.....#..#.......#.
##.............#..#..##....#.#.
#...#...##........#.#.......#..
........#..#.....#.#..#..#.#...
#..##.........#.#.#.##...#....#
............#...#....#..#....#.
.....#.......#......##..#......
.#.....................#......#
...................#....#.#....
.....#....#.....##.............
#....##.#....##..#....##....#..
....#..........#..........#....
.....#.#...............#..##...
...#......###.......#..##......
#.#.#....##..#......#.##.#.....
.#...###..#.....##.........#.#.
..#...#.............#....#.....
#..#.............#.....#.....#.
.#.........#.#...#..#....#...#.
#....#......#....#.#..........#
.........................#.....
...................#...........
#.#...#......#....#............
.#..#........#...##....#....#..
..#......#..#..........##......
#.#....#....##....#.........#..
...#.#.#.#..#....##..#....#..#.
..#..............#.....##......
....#.........#...#.....#..#...
..#..................#.#.......
.....##.##........#.#....#..###
..#.#...#.....#..##..##.#.#.#..
.....#......#............#.....
.#.......#....##...............
...#.................#.....#...
...#.#..#.#...##........#....##
..........##...................
#........#..........#.#........
................#..##.##.#....#
....##..#.#.#...#...#....#.#.#.
..#.........#......##....#.....
.##.........#.....#.#..........
...##...###...........#......##
..#........#......#.....##.#...
###.....#.#.#...#.......#....#.
..##...#....###..##.#.#..##....
..###...##.......#.#..#....#..#
..#...............###....#..#..
...........#....#.##..#........
.#...#..#.#...##..#....#...#..#
..#............#......#.....#..
.#...#...#.#...#.#.............
...####.........#....##....#.#.
.....##...#........#.#......#..
...####...#.#..#.#.#.#.........
........#.##.#..#.......#......
......##......#.........#.#....
..#.#...#....#.....###.....##..
#.#.##..........#...##..#..#.#.
.....#................#.#..#..#
.........#........#.....#..#..#
......#...........#...........#
..#........#.#.........#...##..
.....####.....#....##.#........
....#...#........#.......#...#.
...#..#....#.....##....###.....
........#..#..#.#.#............
#..#......#..#....#....#.#.#..#
.........#...#......##.........
..#....#............#..#.....#.
#............#.#...#......#...#
..#..##...#........#.........##
.#...#....##...#.......#..##...
#..#.##......#........##...#...
...#..........#...#..#..#....#.
##..#........##..##...#..###.#.
............##...............#.
#......#...#....#.........#...#
................#..#.#.........
.....#...#...#...##.......#...#
..##.###...#...#.#..##.#.#...#.
#...##..........#....##.#.#.#..
.#.........#..........#........
.......#.#...............#.....
...#...#............#..........
.........#..#..........#.......
.........#..#...#....#.##....#.
..#............#......#....#.##
...#...#.#........#......#..#.#
........#......##...##...#..#.#
.......###......#............#.
#.....#...##.#.#...#.......#.#.
..#......#..............##....#
..#............##.......#.#.#.#
...#.#.....#.#.#........####...
...#................#..........
..#...#....#....#......#..#...#
.###......#..............#.#..#
......#......#..........##..#..
...##.#...........#.#.....##.#.
.#...#......#..........#.......
....#...#....#..........#.#....
..................##..#.....#.#
###.................#......##..
.....#.....#............#.#..#.
.....#........#...#....#.#.....
#.#...#........................
.#...#.......#..#.......#......
.......#.#.....###.#...#.#.....
#...#.#...........##...#.......
.#.......#.....#..#..#....#....
...#..##.....#..#..#.....#.....
...#................###......#.
#..#...##.###..#..##.......#...
.#.#.#........#.#.............#
#.......#..#.......#.....##...#
.#.#.#............#..#....#.#..
...#.#.##.#......##.....#....##
#............###...#....#......
.....#..#..#.#.........##.#....
.#.##........#.#.#...#.......#.
..###..#..#.#...#.##...###.....
#............#.............#...
.#.##.....#..#.......#...#...#.
.#...#........###...####.......
.#.#..##..#.....#.#..#.........
....#.#.#............##..#...#.
###.##......#.#.....#.....#....
.........#...##.....##....#....
..#................#.........#.
#.......###..##..##............
.....#...#.............#..#..#.
..........#...................#
....#....#...........#.........
.##.......##.##.........##.....
#......#.#....#....#...#.#.#...
..#.##..#.###.#.##....#..#.....
#....##.#...#..................
.......#...#...........#...#...
....###.#...#..#...............
##.#.#..#.#......#.#......#...#
.............#.....#.##....#...
#.............###....#...#.##.#
#..#.##.............#.##...#...
.#.#......#.........#...#......
.#.........#.#.#.....##.#.#....
.................#........#....
....##.#.#..#.........#........
#...##......##....#.#..#......#
..........##...##..#......##...
..........#..#.#..##..#..#.....
..#..#.....##........#...#.#...
#..........#.#.#..............#
#..........##.....#.#...##....#
.....#...#..#..#...##.#.......#
.##.#...............#.#...#....
..........#.....#......#.......
.....#.#......##...#.......#...
...........#.#...#.....#....#.#
.###.#........##....#.##...#...
#....#.##....#.###..##.#.......
##...........#..##.........#...
....#.##...#...#.....#.#..#....
........#.#.#..#.#...........##
..........#.##...#....#......#.
.##.....#.#.....##.#.......#.#.
.#..#....#.#.....#.##.#....#..#
#.......#..#..........##....#..
.#........#...#..#.#...#....#..
#......##...#...##..#.#.......#
.#......#.##.#............##.#.
.#....#.....##..##..........#..
..###..#..#...#...#.#.#..##....
.#.#.##...#..#...........#....#
....#......#.......##...#.#.#..
.......#..#...##..#.........#..
....#..#.#.......##........#..#
........#.#....#.##..#.......#.
.....#.......#.#...#.#.........
........#...#....#.#....###..#.
......#..#.##..##..#...#.#.....
.#.#.....#.....#....#...#...#..
...#..#...#..#......#..#.#.....
...##...#...........#..#......#
..#...#####..#.#.##....##......
...........#......#.#..#.......
..#....##..#.##.......#.#.#..#.
..#..#........#...#.......#....`

	result := findTrees(mapContent)
	expected := 189

	if result != expected {
		t.Errorf("When getting the number of trees, got: %v, want: %v.", result, expected)
	}
}