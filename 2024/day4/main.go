package main

import (
	"bufio"
	"fmt"
	"os"
)

const target = "XMAS"

func main() {
	grid := readFile()

	count := 0
	directions := [][2]int{
		{0, 1},   // right
		{1, 0},   // down
		{1, 1},   // diagonal down-right
		{1, -1},  // diagonal down-left
		{0, -1},  // left
		{-1, 0},  // up
		{-1, 1},  // diagonal up-right
		{-1, -1}, // diagonal up-left
	}

	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, dir := range directions {
				if checkWord(grid, i, j, dir, target) {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}

func checkWord(grid [][]rune, row, col int, dir [2]int, target string) bool {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) {
		return false
	}

	endRow := row + dir[0]*(len(target)-1)
	endCol := col + dir[1]*(len(target)-1)
	if endRow < 0 || endRow >= len(grid) || endCol < 0 || endCol >= len(grid[0]) {
		return false
	}

	for i := 0; i < len(target); i++ {
		currRow := row + dir[0]*i
		currCol := col + dir[1]*i
		if grid[currRow][currCol] != rune(target[i]) {
			return false
		}
	}

	return true
}

func readFile() [][]rune {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	return grid
}
