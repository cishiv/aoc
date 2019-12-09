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

// Size of the parameter set for each opcode
var parameterSetSize = map[int]int{
	// add
	1: 3,
	// mul
	2: 3,
	// in
	3: 1,
	// out
	4: 1,
	// jmp if true (!=0)
	5: 2,
	// jmp if false (==0)
	6: 2,
	// less than
	7: 3,
	// equals
	8: 3,
	// relative
	9: 1,
	// halt
	99: 0,
}

// What an instruction looks like
type instruction struct {
	opcode     int
	parameters int
	modes      []mode
}

// Instruction modes
type mode int

// pos = position mode i.e by reference
// immediate = immediate mode i.e by value
const (
	pos mode = iota
	immediate
	relative
)

func main() {
	memory := createMemory(readFile("./input-day9"))
	runIntcode(memory, 1, "part1-testmode")
	runIntcode(memory, 2, "part1-sigboost")

}

func createMemory(data []int) []int {
	mem := make([]int, 10000000)
	for i, v := range data {
		mem[i] = v
	}
	return mem
}

// concurrency would've been great but I don't want to deal with deadlock
func runIntcode(prog []int, input int, part string) {
	defer stopwatch(time.Now(), part)
	rel := 0
	// we'll manage the instruction pointer as i
	for i := 0; i < len(prog); {
		instruc := buildInstruction(prog[i])
		switch instruc.opcode {
		case 1:
			val0 := fetchInstruction(prog, i+1, instruc.modes[0], rel, false)
			val1 := fetchInstruction(prog, i+2, instruc.modes[1], rel, false)
			tar := 0
			if len(instruc.modes) > 2 {
				tar = fetchInstruction(prog, i+3, instruc.modes[2], rel, true)
			} else {
				tar = prog[i+3]

			}
			prog[tar] = val0 + val1
			i += instruc.parameters + 1
		case 2:
			val0 := fetchInstruction(prog, i+1, instruc.modes[0], rel, false)
			val1 := fetchInstruction(prog, i+2, instruc.modes[1], rel, false)
			tar := 0
			if len(instruc.modes) > 2 {
				tar = fetchInstruction(prog, i+3, instruc.modes[2], rel, true)
			} else {
				tar = prog[i+3]

			}
			prog[tar] = val0 * val1
			i += instruc.parameters + 1
		case 3:
			tar := 0
			if len(instruc.modes) > 0 {
				tar = fetchInstruction(prog, i+1, instruc.modes[0], rel, true)
			} else {
				tar = prog[i+1]
			}
			prog[tar] = input
			i += instruc.parameters + 1
		case 4:
			val := fetchInstruction(prog, i+1, instruc.modes[0], rel, false)
			// move the pointer
			i += instruc.parameters + 1
			fmt.Println(val)
		case 5:
			val0 := fetchInstruction(prog, i+1, instruc.modes[0], rel, false)
			val1 := fetchInstruction(prog, i+2, instruc.modes[1], rel, false)
			if val0 != 0 {
				// set instruction pointer to 2nd param
				i = val1
			} else {
				i += instruc.parameters + 1
			}
		case 6:
			val0 := fetchInstruction(prog, i+1, instruc.modes[0], rel, false)
			val1 := fetchInstruction(prog, i+2, instruc.modes[1], rel, false)
			if val0 == 0 {
				i = val1
			} else {
				i += instruc.parameters + 1
			}
		case 7:
			val0 := fetchInstruction(prog, i+1, instruc.modes[0], rel, false)
			val1 := fetchInstruction(prog, i+2, instruc.modes[1], rel, false)
			tar := 0
			if len(instruc.modes) > 2 {
				tar = fetchInstruction(prog, i+3, instruc.modes[2], rel, true)
			} else {
				tar = prog[i+3]
			}
			if val0 < val1 {
				prog[tar] = 1
			} else {
				prog[tar] = 0
			}
			i += instruc.parameters + 1
		case 8:
			val0 := fetchInstruction(prog, i+1, instruc.modes[0], rel, false)
			val1 := fetchInstruction(prog, i+2, instruc.modes[1], rel, false)
			tar := 0
			if len(instruc.modes) > 2 {
				tar = fetchInstruction(prog, i+3, instruc.modes[2], rel, true)
			} else {
				tar = prog[i+3]
			}
			if val0 == val1 {
				prog[tar] = 1
			} else {
				prog[tar] = 0
			}
			i += instruc.parameters + 1
		case 9:
			val := fetchInstruction(prog, i+1, instruc.modes[0], rel, false)
			rel += val
			i += instruc.parameters + 1
		// goodbye
		case 99:
			return
		}
	}
}

func fetchInstruction(prog []int, in int, m mode, rel int, write bool) int {
	val := prog[in]
	if !write {
		switch m {
		// if it's position mode, return ref
		case pos:
			return prog[val]
		// otherwise return value
		case immediate:
			return val
		case relative:
			return prog[val+rel]
		}
	} else {
		switch m {
		// if it's position mode, return ref
		case pos:
			return val
		// otherwise return value
		case immediate:
			return val
		case relative:
			return val + rel
		}
	}
	// that didn't work out...
	return 0
}

func buildInstruction(in int) instruction {
	var instr instruction
	// pos len - 1
	instr.opcode = in % 100
	instr.parameters = parameterSetSize[instr.opcode]
	for in /= 100; len(instr.modes) < instr.parameters; in /= 10 {
		// abusing mod to parse digits ^_^
		instr.modes = append(instr.modes, mode(in%10))
	}
	return instr
}

func readFile(filePath string) []int {
	file, err := os.Open(filePath)
	var prog []int
	if err != nil {
		fmt.Printf("err")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		codes := strings.Split(val, ",")
		for _, v := range codes {
			val, err := strconv.Atoi(v)
			handleErr(err)
			prog = append(prog, val)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return prog
}

func handleErr(err error) {
	if err != nil {
		fmt.Println("Error")
	}
}

func stopwatch(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
