package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var mulReg = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func main() {
	content, err := Read("input.txt")
	if err != nil {
		fmt.Printf("Error al leer el archivo: %v\n", err)
		return
	}

	result := 0
	matches := mulReg.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		num1 := toInt(match[1])
		num2 := toInt(match[2])
		result += num1 * num2
	}
	fmt.Println(result)
}

func Read(fileName string) (string, error) {
	lines := ""
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error al abrir el archivo: %v\n", err)
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error al leer el archivo: %v\n", err)
		return "", err
	}

	return lines, nil
}

func toInt(x string) int {
	num, err := strconv.Atoi(x)
	if err != nil {
		fmt.Printf("Error al convertir: %s\n", x)
		panic(err)
	}
	return num
}
