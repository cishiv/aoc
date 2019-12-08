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
)

func main() {
	data := readFile("./input-day7")
	val := len(data)
	// 2019/12/07 09:34:57 part1 took 15.255814ms
	fmt.Println(run(data, []int{0, 1, 2, 3, 4}, "part1"))
	// 2019/12/07 09:34:57 part2 took 17.091813ms
	fmt.Println(run(data, []int{5, 6, 7, 8, 9}, "part2"))
}

func run(prog []int, phase []int, part string) int {
	defer stopwatch(time.Now(), part)
	best := 0
	for _, phase := range permutations(phase) {
		sig := emulate(prog, phase)
		best = max(best, sig)
	}
	return best
}
func emulate(prog []int, phase []int) int {
	// buffer this one to receive result
	c1 := make(chan int, 1)
	c2 := make(chan int)
	c3 := make(chan int)
	c4 := make(chan int)
	c5 := make(chan int)

	halt := make(chan bool)

	// vroom vroom
	/**
		The amplifiers each feed each other the result of computation as an input
		i.e .. a1 -- > a2 --> a3 --> a4 --> a5 -- > a1
		until a halt signal is produced from a5
		so we create a concurrent pipe of in/out channels between the amps
		in c1 -> out c2 A1
		in c2 -> out c3 A2
		in c3 -> out c4 A3
		in c4 -> out c5 A4
		in c5 -> out c1 A5
	**/
	go vm(prog, c1, c2, halt)
	go vm(prog, c2, c3, halt)
	go vm(prog, c3, c4, halt)
	go vm(prog, c4, c5, halt)
	go vm(prog, c5, c1, halt)

	/**
	The program that the amplifiers execute is 'phase-modified' by the phase value we pass into it
	i.e it takes a phase input to determine how the input signal will be amplified
	sequentially this would involve passing the phase signal as input to the program, editing the op codes, and then passing the input signal in.
	since we have created a feedback loop using channels, we can just pass the phase signal in on the input channel
	*/
	c1 <- phase[0]
	c2 <- phase[1]
	c3 <- phase[2]
	c4 <- phase[3]
	c5 <- phase[4]

	// initial input
	c1 <- 0

	// monitor the halt channels to see if we're done
	for j := 0; j < 5; j++ {
		<-halt
	}

	return <-c1
}

func vm(data []int, in <-chan int, out chan<- int, halt chan<- bool) {
	// move it into memory
	prog := make([]int, len(data))
	copy(prog, data)
	// we'll manage the instruction pointer as i
	for i := 0; i < len(prog); {
		instruc := buildInstruction(prog[i])
		switch instruc.opcode {
		// ADD instruction
		case 1:
			val0 := fetchInstruction(prog, i+1, instruc.modes[0])
			val1 := fetchInstruction(prog, i+2, instruc.modes[1])
			tar := prog[i+3]
			prog[tar] = val0 + val1
			i += instruc.parameters + 1
		// MUL instruction
		case 2:
			val0 := fetchInstruction(prog, i+1, instruc.modes[0])
			val1 := fetchInstruction(prog, i+2, instruc.modes[1])
			tar := prog[i+3]
			prog[tar] = val0 * val1
			i += instruc.parameters + 1
		// REPLACE instruction (actually the input instr but logically it's a replace at reg param)
		case 3:
			tar := prog[i+1]
			prog[tar] = <-in
			i += instruc.parameters + 1
		// OUT instruction
		case 4:
			val := fetchInstruction(prog, i+1, instruc.modes[0])
			// move the pointer
			i += instruc.parameters + 1
			out <- val
		// MOV PNTR NOT EQUAL
		case 5:
			val0 := fetchInstruction(prog, i+1, instruc.modes[0])
			val1 := fetchInstruction(prog, i+2, instruc.modes[1])
			if val0 != 0 {
				// set instruction pointer to 2nd param
				i = val1
			} else {
				i += instruc.parameters + 1
			}
		// MOV PNTR EQUAL
		case 6:
			val0 := fetchInstruction(prog, i+1, instruc.modes[0])
			val1 := fetchInstruction(prog, i+2, instruc.modes[1])
			if val0 == 0 {
				i = val1
			} else {
				i += instruc.parameters + 1
			}
		// JMP LESSTHAN
		case 7:
			val0 := fetchInstruction(prog, i+1, instruc.modes[0])
			val1 := fetchInstruction(prog, i+2, instruc.modes[1])
			tar := prog[i+3]
			if val0 < val1 {
				prog[tar] = 1
			} else {
				prog[tar] = 0
			}
			i += instruc.parameters + 1
		// JMP MORETHAN
		case 8:
			val0 := fetchInstruction(prog, i+1, instruc.modes[0])
			val1 := fetchInstruction(prog, i+2, instruc.modes[1])
			tar := prog[i+3]
			if val0 == val1 {
				prog[tar] = 1
			} else {
				prog[tar] = 0
			}
			i += instruc.parameters + 1
		// SIG EXIT
		case 99:
			halt <- true
			return
		default:
			panic("Invalid op code")
		}
	}
}

func fetchInstruction(prog []int, in int, m mode) int {
	val := prog[in]
	switch m {
	// if it's position mode, return ref
	case pos:
		return prog[val]
	// otherwise return value
	case immediate:
		return val
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
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func stopwatch(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

// https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
