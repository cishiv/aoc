package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// ADD The OpCode for the ADD Instruction
const ADD = "1"

// MUL The OpCode for the MUL Instruction
const MUL = "2"

// HALT The OpCode for the HALT Instruction
const HALT = "99"

func main() {
	// part 1
	partOne("./input-day2")
	// part 2
	partTwo("./input-day2")
}

func partOne(filePath string) {
	defer stopwatch(time.Now(), "part 1")
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
		codes = assignInputs("12", "2", codes)
	}
	resultProgram := compute(codes)
	fmt.Println(resultProgram[0])
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func partTwo(filePath string) {
	defer stopwatch(time.Now(), "part 2")
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
				codes = assignInputs(strconv.FormatInt(int64(i), 10), strconv.FormatInt(int64(j), 10), codes)
				resultProgram := compute(codes)
				if resultProgram[0] == "19690720" {
					fmt.Println(100*i + j)
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func compute(instr []string) []string {
	codes := instr
	pos := 0
	execution := true
	for execution {
		if codes[pos] == HALT {
			execution = false
			break
		} else {
			if codes[pos] == ADD {
				val0 := getPosition(pos+1, codes)
				val1 := getPosition(pos+2, codes)
				index, err := strconv.Atoi(codes[pos+3])
				handleErr(err)
				codes[index] = strconv.FormatInt(int64(add(val0, val1)), 10)
			}
			if codes[pos] == MUL {
				val0 := getPosition(pos+1, codes)
				val1 := getPosition(pos+2, codes)
				index, err := strconv.Atoi(codes[pos+3])
				handleErr(err)
				codes[index] = strconv.FormatInt(int64(mult(val0, val1)), 10)
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
	return result
}

func assignInputs(val0 string, val1 string, codes []string) []string {
	result := codes
	result[1] = val0
	result[2] = val1
	return result
}

func mult(val0 int, val1 int) int {
	return val0 * val1
}

func add(val0 int, val1 int) int {
	return val0 + val1
}

func stopwatch(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
