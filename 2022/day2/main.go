package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//67-90 tijera -> 3
	//66-89 papel -> 2
	//65-88 piedra -> 1
	// 6 win
	// 3 draw
	// 0 lose
	// 90 -> win
	// 89 -> draw
	// 88 -> lose
	m := make(map[int]map[int]int)
	m[65] = make(map[int]int)
	m[65][88] = 3 //4
	m[65][89] = 4 //8
	m[65][90] = 8 //3

	m[66] = make(map[int]int)
	m[66][88] = 1 //1
	m[66][89] = 5 //5
	m[66][90] = 9 //9

	m[67] = make(map[int]int)
	m[67][88] = 2 //7
	m[67][89] = 6 //2
	m[67][90] = 7 //6

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	total := 0
	for fileScanner.Scan() {
		s := fileScanner.Text()
		total += m[int(s[0])][int(s[2])]
	}

	fmt.Println(total)
	file.Close()
}
