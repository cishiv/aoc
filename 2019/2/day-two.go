package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

var baseMemory []int
var activeMemory []int

func main() {
	// part 1
	// partOne("./input-day2")
	// part 2
	partTwo("./input-day2")
}

func partOne(filePath string) {
	var codes []string
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("err")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		codes = strings.Split(val, ",")
		fmt.Println(codes)
		codes[1] = "12"
		codes[2] = "2"
	}
	fmt.Println(codes)
	resultProgram := compute(codes)
	fmt.Println(resultProgram[0])
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func partTwo(filePath string) {
	var codes []string
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("err")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		for i := 1; i <= 99; i++ {
			for j := 1; j <= 99; j++ {
				codes = strings.Split(val, ",")
				codes[1] = strconv.FormatInt(int64(i), 10)
				codes[2] = strconv.FormatInt(int64(j), 10)
				resultProgram := compute(codes)
				if (resultProgram[0] == "19690720") {
					fmt.Println(100 * i + j)
					//fmt.Println("FOUND IT")
					break
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
func mult(val0 int, val1 int) int {
	return val0 * val1
}

func add(val0 int, val1 int) int {
	return val0 + val1
}

func compute(instr []string) []string {
	codes := instr
	pos := 0 
	execution := true
	for execution {
		if (codes[pos] == "99") {
			//fmt.Println(strconv.FormatInt(int64(pos), 10) + " HALT")
			execution = false
			break
		} else {
			if (codes[pos] == "1") {
				val0 := getPosition(pos+1, codes)
				val1 := getPosition(pos+2, codes)
				index, err :=  strconv.Atoi(codes[pos+3])
				handleErr(err)
				codes[index] = strconv.FormatInt(int64(add(val0,val1)), 10)
				//fmt.Println("Adding " + strconv.FormatInt(int64(val0), 10) + " and " + strconv.FormatInt(int64(val1), 10) + " equal " + codes[index])
			}
			if (codes[pos] == "2") {
				val0 := getPosition(pos+1, codes)
				val1 := getPosition(pos+2, codes)
				index, err :=  strconv.Atoi(codes[pos+3])
				handleErr(err)
				codes[index] = strconv.FormatInt(int64(mult(val0,val1)), 10)
				//fmt.Println("Multiplying " + strconv.FormatInt(int64(val0), 10) + " and " + strconv.FormatInt(int64(val1), 10) + " equal " + codes[index])
			}
		}
		pos = pos + 4
	}
	return codes
}

func handleErr(err error) {
	if err != nil {
		fmt.Println("error")
	}
}

func getPosition(val int, codes []string) int {
	resultPos, err := strconv.Atoi(codes[val])
	handleErr(err)
	result, err := strconv.Atoi(codes[resultPos])
	handleErr(err)
	return  result
}
