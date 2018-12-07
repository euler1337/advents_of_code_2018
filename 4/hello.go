package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"sort"
	"strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func categorizeInput(s string) string {

	i_hashtag := strings.Index(s, "#")

	if(i_hashtag > 0) {
		return "shift"
	}

	i_w := strings.Index(s, "w")

	if(i_w > 0) {
		return "wake"
	}

	return "sleep"
}

func parseTime(s string) (int) {
	i_colon := strings.Index(s, ":")
	
	hour := s[i_colon-3:i_colon-1]
	min := s[i_colon+1:i_colon+3]

	hInt, _ := strconv.Atoi(hour)
	mInt, _ := strconv.Atoi(min)

	if hInt == 23 {
		mInt = 0
	}

	if hInt == 01 {
		mInt = 60
	}


	return mInt
}

func parseGuard(s string) (string){
	i_hashtag := strings.Index(s, "#")
	i_b := strings.Index(s, "b")

	guardId := s[i_hashtag+1:i_b-1]

	return guardId
}

func calc1() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)


	//var m map[string]int
	//m = make(map[string]int)

	var inputs[]string
	for scanner.Scan() {
		input := scanner.Text()
		inputs = append(inputs, input)
	}
	sort.Strings(inputs)

	// Find the guard that has most minutes asleep
	// Find the minute where this guard is most often asleep

	mapSleepMinutes := make(map[string]int)
	mapSleepMinutesList := make(map[string][]int)
	currentGuardId := ""
	fallAsleepMin := 0

	for _, input_str := range inputs {
		inputType := categorizeInput(input_str)
		
		if inputType == "shift" {
			currentGuardId = parseGuard(input_str)
			fmt.Printf("shift : guard[%s] \n", currentGuardId)

			_, ok := mapSleepMinutes[currentGuardId]
			if !ok {
				mapSleepMinutes[currentGuardId] = 0
			}

		} 
		if inputType == "wake" {
			minute := parseTime(input_str)
			sleepTime := minute - fallAsleepMin
			mapSleepMinutes[currentGuardId] = mapSleepMinutes[currentGuardId] + sleepTime
			fmt.Printf("wakeup : min[%d], sleeptime=[%d]\n", minute, sleepTime)

			for i := fallAsleepMin; i < minute; i++ {
				mapSleepMinutesList[currentGuardId] = append(mapSleepMinutesList[currentGuardId], i) 
			}
		} 
	
		if inputType == "sleep" {
			minute := parseTime(input_str)
			fallAsleepMin = minute

			fmt.Printf("sleep : min[%d]\n", minute)
		} 
	}

	maxTime := 0
	maxTimeGuard := ""
	for guard, time := range mapSleepMinutes {
		fmt.Printf("GUARD: guard[%s], tot time=[%d]\n", guard, time)
		if time > maxTime {
			maxTime = time
			maxTimeGuard = guard
		}
	}
	
	fmt.Printf("TIME: guard[%s], tot time=[%d]\n", maxTimeGuard, maxTime)

	// Find the most popular minute:
	sleepMinutesMap := make(map[int]int)
	minuteList := mapSleepMinutesList[maxTimeGuard]
	
	for _, i := range minuteList {
		_, ok := sleepMinutesMap[i]
		if ok {
			sleepMinutesMap[i] = sleepMinutesMap[i]+1
		} else {
			sleepMinutesMap[i] = 1
		}
	}

	mostPopKey := -1
	mostPopVal := 0
	for key, val:= range sleepMinutesMap {
		fmt.Printf("MINUTE: minute[%d], sleepCount=[%d]\n", key, val)
		if (val > mostPopVal) {
			mostPopKey = key
			mostPopVal = val
		}
	}
	fmt.Printf("MINUTES: minute[%d], sleepCount=[%d]\n", mostPopKey, mostPopVal)

	GuardIdInt, _ := strconv.Atoi(maxTimeGuard)
	fmt.Printf("ANSWER IS: %d\n", mostPopKey * GuardIdInt)
}


func main() {
	calc1()
	//calc2()

}
