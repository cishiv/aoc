package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	partOne()
	partTwo()
}

func calcFuel(fuel float64, total float64) float64 {
	div := math.Floor(float64(fuel)/3.0) - 2.0
	if div <= 0.0 {
		return total
	} else {
		total += div
		return calcFuel(div, total)
	}
}

func partTwo() {
	defer stopwatch(time.Now(), "part 2")
	total := 0.0
	file, err := os.Open("./input-day1")
	handleErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		localTotal := 0.0
		val, err := strconv.ParseFloat(scanner.Text(), 64)
		handleErr(err)
		total += calcFuel(val, localTotal)
	}
	fmt.Println("part 2: " + fmt.Sprintf("%f", total))
}

func partOne() {
	defer stopwatch(time.Now(), "part 1")
	fuel := 0.0
	file, err := os.Open("./input-day1")
	handleErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.ParseFloat(scanner.Text(), 64)
		handleErr(err)
		fuel += math.Floor(val/3.0) - 2
	}
	fmt.Println("part 1: " + fmt.Sprintf("%f", fuel))
}

func handleErr(err error) {
	if err != nil {
		fmt.Println("error")
	}
}

func stopwatch(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
