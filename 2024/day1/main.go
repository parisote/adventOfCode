package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const _separator = "   "

func main() {
	lineA, lineB, err := Read("input.txt")
	if err != nil {
		fmt.Printf("Error al leer el archivo: %v\n", err)
		return
	}

	lineA = quickSort(lineA)
	lineB = quickSort(lineB)

	if len(lineA) != len(lineB) {
		fmt.Println("Las lineas no tienen el mismo numero de elementos")
		return
	}

	var sum int
	for i := 0; i < len(lineA); i++ {
		diff := lineA[i] - lineB[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}

	fmt.Println(sum)
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2] //mid
	var left, middle, right []int

	for _, x := range arr {
		if x < pivot {
			left = append(left, x)
		} else if x == pivot {
			middle = append(middle, x)
		} else {
			right = append(right, x)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, middle...), right...)
}

func Read(fileName string) ([]int, []int, error) {
	lineA := make([]int, 0)
	lineB := make([]int, 0)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error al abrir el archivo: %v\n", err)
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines := strings.Split(line, _separator)
		a, b := lines[0], lines[1]
		lineA = append(lineA, toInt(a))
		lineB = append(lineB, toInt(b))
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error al leer el archivo: %v\n", err)
		return nil, nil, err
	}

	return lineA, lineB, nil
}

func toInt(linea string) int {
	num, err := strconv.Atoi(linea)
	if err != nil {
		fmt.Printf("Error al convertir la linea: %v\n", err)
		panic(err)
	}
	return num
}
