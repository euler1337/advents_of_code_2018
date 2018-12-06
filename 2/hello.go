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

func calc1() {
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

func CalcNumberOfDifferentChars(e1 string ,  e2 string) (int, int) {
	l1 := len(e1)
	l2 := len(e2)



	if l1 != l2 {
		return -1, 0
	}

	diffIndex := 0
	diffCount := 0
	for index := range e1 {
		if e1[index] != e2[index] {
			diffCount+=1
			diffIndex = index
		}
	}

	return diffCount, diffIndex
}

func calc2() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Find number of rows which contain a letter that occurs exactly twice
	// Find number of rows which contain a letter that occurs exactly three times
	// Multiply the

	inputs := make([]string, 0)

	for scanner.Scan() {
		input_str := scanner.Text()
		inputs = append(inputs, input_str)
	}

	for _, e1 := range inputs {
		for _, e2 := range inputs {
			diffCount, diffIndex := CalcNumberOfDifferentChars(e1, e2)
			if diffCount == 1 {
				fmt.Printf("FOUND ONE: e1 = %s, e2 = %s \n", e1, e2)
				fmt.Printf("Index: %d, ANSWER ---------> %s%s <-------------\n", diffIndex, e1[0:diffIndex], e1[diffIndex+1:])
			}
		}
	}



}


func main() {
	//calc1()
	calc2()

}
