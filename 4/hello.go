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

	// 2

	// Which minute is any guard most often asleep on?
	globalMaxKey := 0
	globalMaxVal := -1
	
	for guard, listVal := range mapSleepMinutesList {
		
		sleepMinutesMap := make(map[int]int)

		//fmt.Printf("Guard[%s], minute=%d\n", guard, len(sleepMinutesMap))

		for _, i := range listVal {
			_, ok := sleepMinutesMap[i]
			if ok {
				sleepMinutesMap[i] = sleepMinutesMap[i]+1
			} else {
				sleepMinutesMap[i] = 1
			}
		}

		fmt.Printf("Guard[%s], Number of keys = %d\n", guard, len(sleepMinutesMap))

		forThisGuardMaxKey := 0
		forThisGuardMaxVal := 0

		for key, val := range sleepMinutesMap {
			if val > forThisGuardMaxVal {
				forThisGuardMaxVal = val
				forThisGuardMaxKey = key
			}
		}

		fmt.Printf("GUARD = %s, MIN = %d COUNT: %d\n",guard, forThisGuardMaxKey, forThisGuardMaxVal)

		if forThisGuardMaxVal > globalMaxVal {
			globalMaxVal = forThisGuardMaxKey
			globalMaxKey, _ = strconv.Atoi(guard)
		}
	}
	fmt.Printf("B) KEY = %d, MIN = %d ANSWER IS: %d\n",globalMaxKey, globalMaxVal, globalMaxKey * globalMaxVal)

	//Became to messy. Looked at the logs and found the answer. it was guard 1097 * minute 21. Lol O_o
}

func main() {
	calc1()
	//calc2()

}
