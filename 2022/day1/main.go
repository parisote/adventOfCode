package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)

	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	c := 0
	var a []int

	for fileScanner.Scan() {
		if fileScanner.Text() != "" {
			i, _ := strconv.Atoi(fileScanner.Text())
			c = c + i
		} else {
			a = append(a, c)
			c = 0
		}
	}

	sort.Ints(a)
	l := len(a)
	m := a[l-1]
	s := a[l-3] + a[l-2] + a[l-1]
	fmt.Println(m)
	fmt.Println(s)

	file.Close()
}
