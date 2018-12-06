package main

import (
	"fmt"
	"os"
	"bufio"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func main() {
	

	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Find number of rows which contain a letter that occurs exactly twice
	// Find number of rows which contain a letter that occurs exactly three times
	// Multiply the

	var m map[string]int
	

	count2 := 0
	count3 := 0
	

	for scanner.Scan() {
		found2, found3 := false, false

		m = make(map[string]int)
		input_str := scanner.Text()
		for _, char := range input_str {
			key := string(char)
			_, ok := m[key]
			if ok {
				// char exists. increase count.
				m[key] = m[key] + 1

			} else {
				// char does not exist yet. Add it.
				m[key] = 1
			}
		}

		for k, v := range m { 
			fmt.Printf("key[%s] value[%d]\n", k, v)
			if !found2 && (v == 2) {
				count2 += 1
				found2 = true
			}
			if !found3 && (v == 3) {
				count3 += 1
				found3 = true
			}
		}
	}
	fmt.Printf("2 count [%d], 3 count [%d], mult [%d] \n", count2, count3, count2*count3)
}
