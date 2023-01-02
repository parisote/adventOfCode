package main

import (
	"bufio"
	"log"
	"fmt"
	"os"
	"flag"
)

var input string

func main(){
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	if part == 1{
		part1(fileScanner)
	} else {
		part2(fileScanner)
	}

}

func part1(fileScanner *bufio.Scanner) {
	fileScanner.Split(bufio.ScanLines)

	total := 0

	for fileScanner.Scan() {
		s := fileScanner.Text()
		m := make(map[byte]int)
		for i := 0; i < (len(s)/2); i++ {
			m[s[i]] = 1
		}
		for i := len(s)/2; i < len(s); i++ {
			if _, ok := m[s[i]]; ok {
				total += getValue(s[i])				
				break
			}			
		}
	}

	fmt.Println(total)
}

func getValue(item byte) int{
	var priority int
    if item >= 'a' && item <= 'z' {
    	priority = int(item - 'a') + 1
    } else {
        priority = int(item - 'A') + 27
    }
	
	return priority
}

func part2(fileScanner *bufio.Scanner) {
	fileScanner.Split(bufio.ScanLines)

	total := 0
	i := 0
	m := make(map[int]map[byte]bool)
	for fileScanner.Scan() {
		s := fileScanner.Text()
		
		m[i] = make(map[byte]bool)
		for j:= 0; j < len(s); j++ {
			m[i][s[j]] = true
		}
		i++

		if i == 3 {
			fmt.Println(m[0])
			fmt.Println(m[1])
			fmt.Println(m[2])
			fmt.Println("VERIFICO CUAL")
			for k, _ := range m[0] {
				if m[1][k] && m[2][k] {
					total += getValue(k)
					break
				} 
			}
			m = make(map[int]map[byte]bool)
			i = 0
		}
	}

	fmt.Println(total)
}
