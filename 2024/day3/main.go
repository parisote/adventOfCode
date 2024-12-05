package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	doCommand   = "do()"
	dontCommand = "don't()"
	mulCommand  = "mul"
	mulPattern  = `mul\((\d{1,3}),(\d{1,3})\)`
)

var mulReg = regexp.MustCompile(mulPattern)

func main() {
	content, err := readFile("input.txt")
	if err != nil {
		fmt.Printf("Error al leer el archivo: %v\n", err)
		os.Exit(1)
	}

	result := processContent(content)
	fmt.Println(result)
}

func processContent(content string) int {
	result := 0
	doFlag := true

	for i := 0; i < len(content); i++ {
		doFlag = updateFlag(content, i, doFlag)
		result += processMultiplication(content, i, doFlag)
	}

	return result
}

func updateFlag(content string, i int, currentFlag bool) bool {
	if content[i] != 'd' || i+4 >= len(content) {
		return currentFlag
	}

	switch {
	case i+4 < len(content) && content[i:i+4] == doCommand:
		return true
	case i+7 < len(content) && content[i:i+7] == dontCommand:
		return false
	}
	return currentFlag
}

func processMultiplication(content string, i int, doFlag bool) int {
	if !doFlag || i+3 >= len(content) || content[i:i+3] != mulCommand {
		return 0
	}

	if i+12 >= len(content) || !mulReg.MatchString(content[i:i+12]) {
		return 0
	}

	matches := mulReg.FindStringSubmatch(content[i : i+12])
	num1 := toInt(matches[1])
	num2 := toInt(matches[2])
	return num1 * num2
}

func readFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("error al leer el archivo: %w", err)
	}
	return string(content), nil
}

func toInt(x string) int {
	num, err := strconv.Atoi(x)
	if err != nil {
		fmt.Printf("Error al convertir: %s\n", x)
		panic(err)
	}
	return num
}
