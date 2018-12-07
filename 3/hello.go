package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func parseRow(s string) (int, int, int, int){
	i_at := strings.Index(s, "@")
	i_comma := strings.Index(s, ",")
	i_colon := strings.Index(s, ":")
	i_x := strings.Index(s, "x")

	xStart := s[i_at+2:i_comma]
	yStart := s[i_comma+1:i_colon]

	x := s[i_colon+2:i_x]
	y := s[i_x+1:]

	fmt.Printf("%s, xs[%s] ys[%s], x[%s], y[%s] \n", s, xStart, yStart, x, y)

	xStart_int, err := strconv.Atoi(xStart)
	yStart_int, err := strconv.Atoi(yStart)
	x_int, err := strconv.Atoi(x)
	y_int, err := strconv.Atoi(y)

	check(err)

	return xStart_int, yStart_int, x_int, y_int

}

func calc1() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)


	var m map[string]int
	m = make(map[string]int)

	for scanner.Scan() {
		input_str := scanner.Text()
		xs, ys, x, y := parseRow(input_str)

		for xcoord := 0; xcoord < x; xcoord++ {
			for ycoord := 0; ycoord < y; ycoord++ {
				key := strconv.Itoa(xs + xcoord) + "_" + strconv.Itoa(ys + ycoord)
				_, ok := m[key]
				if ok {
					// char exists. increase count.
					m[key] = m[key] + 1
	
				} else {
					// char does not exist yet. Add it.
					m[key] = 1
				}
			}
		}
	}

	count := 0
	for _, v := range m { 
		if(v > 1) {
			count+=1
		}
	}
	fmt.Printf("count[%d]", count)
}

func calc2() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)


	m := make(map[string][]int)

	index := 0
	
	// Create blacklist for all level indexes.  

	for scanner.Scan() {
		index++

		input_str := scanner.Text()
		xs, ys, x, y := parseRow(input_str)

		for xcoord := 0; xcoord < x; xcoord++ {
			for ycoord := 0; ycoord < y; ycoord++ {
				key := strconv.Itoa(xs + xcoord) + "_" + strconv.Itoa(ys + ycoord)
				_, ok := m[key]
				if ok {
					m[key] = append(m[key], index)
				} else {
					m[key] = append(m[key], index)	
				}			
			}
		}
	}

	blacklist := make(map[int]int)
	for _, v := range m { 
		if(len(v) > 1) {
			for _, element := range v  {
				blacklist[element] = blacklist[element] + 1
				//fmt.Printf("add element[%d] to blacklist", element)
			}
		}
	}

	// Find the index which is not in blacklist
	index_search := 0
	for {
		index_search++
		_, ok := blacklist[index_search]
		if !ok {
			fmt.Printf("index[%d] is not in blacklist", index_search)
			break;
		}
	}
}

func main() {
	//calc1()
	calc2()

}
