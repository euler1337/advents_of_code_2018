package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readFiles(sum int, m map[int]int) (int, bool) {
	finished := false
	finish_val := 0

	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if string(scanner.Text()[0]) == "+" {
			x, err := strconv.Atoi(scanner.Text()[1:])
			check(err)
			sum += x
		}

		if string(scanner.Text()[0]) == "-" {
			x, err := strconv.Atoi(scanner.Text()[1:])
			check(err)
			sum -= x
		}

		i, ok := m[sum]

		if ok {
			//Finished
			finished = true
			finish_val = sum
			i+=1
			fmt.Println("FIRST REOCCURING: ")
			fmt.Println(sum)
			return finish_val, finished
		} else {
			m[sum] = sum
		}
	}
	return sum, finished
}


func main() {
	
	sum := 0
	finished := false
	m := make(map[int]int)
	for {
		sum, finished = readFiles(sum,m)
		if finished {
			break
		}
	}

	// answer to 1
	//fmt.Println("TOTAL SUM: ")
	//fmt.Println(sum)
}