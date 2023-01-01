package main

import (
	"bufio"
	"log"
	"fmt"
	"os"
)

func main(){
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