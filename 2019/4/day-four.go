package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	range0 := 265275
	range1 := 781584
	fmt.Println(partOne(range0, range1))
	fmt.Println(partTwo(range0, range1))
}

func partOne(range0 int, range1 int) int {
	defer stopwatch(time.Now(), "part1")
	tot := 0
	for j := range0; j <= range1; j++ {
		if cond0(j) {
			if cond1(j) {
				tot++
			}
		}
	}
	return tot
}

func cond0(val0 int) bool {
	s := strconv.Itoa(val0)
	s1 := strings.Split(s, "")
	for i := 0; i < len(s1)-1; i++ {
		val0, err := strconv.Atoi(s1[i])
		handleErr(err)
		val1, err := strconv.Atoi(s1[i+1])
		if val0 == val1 {
			return true
		}
	}
	return false
}

func cond1(val0 int) bool {
	s := strconv.Itoa(val0)
	s1 := strings.Split(s, "")
	for i := 0; i < len(s1)-1; i++ {
		val0, err := strconv.Atoi(s1[i])
		handleErr(err)
		val1, err := strconv.Atoi(s1[i+1])
		handleErr(err)
		if val0 > val1 {
			return false
		}
	}
	return true
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func partTwo(range0 int, range1 int) int {
	defer stopwatch(time.Now(), "part2")
	tot := 0
	for j := range0; j <= range1; j++ {
		if cond2(j) {
			if cond1(j) {
				tot++
			}
		}
	}
	return tot
}

func cond2(val0 int) bool {
	holds := false
	m0 := make(map[int]int)
	s := strconv.Itoa(val0)
	s1 := strings.Split(s, "")
	for i := 0; i < len(s1)-1; i++ {
		val0, err := strconv.Atoi(s1[i])
		handleErr(err)
		val1, err := strconv.Atoi(s1[i+1])
		if val0 == val1 {
			m0[val0]++
		}
	}
	for i, _ := range m0 {
		m0[i]++
		if m0[i] == 2 {
			holds = true
		}
	}
	return holds
}

func stopwatch(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
