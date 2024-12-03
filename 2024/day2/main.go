package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const _separator = " "

func main() {
	levels, err := Read("input.txt")
	if err != nil {
		fmt.Printf("Error al leer el archivo: %v\n", err)
		return
	}

	result := 0
	for _, level := range levels {
		if isSafeWithDampener(level) {
			result++
		}
	}
	fmt.Println(result)
}

// PART2
func isSafeWithDampener(level []int) bool {
	if isSafe(level) {
		return true
	}

	for i := range level {
		tempLevel := make([]int, len(level)-1)

		index := 0
		for _, v := range level[:i] {
			tempLevel[index] = v
			index++
		}

		for _, v := range level[i+1:] {
			tempLevel[index] = v
			index++
		}

		if isSafe(tempLevel) {
			return true
		}
	}

	return false
}

// PART1
func isSafe(level []int) bool {
	if len(level) <= 1 {
		return true
	}

	isIncreasing := level[1] > level[0]
	for i := 0; i < len(level)-1; i++ {
		rest := level[i+1] - level[i]
		if abs(rest) > 3 || rest == 0 {
			return false
		}

		if isIncreasing && rest < 0 || !isIncreasing && rest > 0 {
			return false
		}
	}

	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Read(fileName string) ([][]int, error) {
	levels := make([][]int, 0)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error al abrir el archivo: %v\n", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		level := strings.Split(line, _separator)
		levels = append(levels, make([]int, 0))
		for _, l := range level {
			levels[count] = append(levels[count], toInt(l))
		}
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error al leer el archivo: %v\n", err)
		return nil, err
	}

	return levels, nil
}

func toInt(linea string) int {
	num, err := strconv.Atoi(linea)
	if err != nil {
		fmt.Printf("Error al convertir la linea: %v\n", err)
		panic(err)
	}
	return num
}
