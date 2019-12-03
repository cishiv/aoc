package main

import (
	"os"
	"bufio"
	"time"
	"fmt"
	"log"
	"strings"
	"strconv"
	"math"
) 

const UP = "U"
const DOWN = "D"
const LEFT = "L"
const RIGHT = "R"

type Position struct {
	X int
	Y int
}

func (p Position) GetX() int {
	return p.X
}

func (p Position) GetY() int {
	return p.Y
}

func main() {
	wire0_path := parse("./wire0")
	wire1_path := parse("./wire1")
	coords0 := buildCoordinateSet(wire0_path)
	coords1 :=  buildCoordinateSet(wire1_path)
	intersections := findIntersections(coords0, coords1)
	// fmt.Println(findShortest(intersections))
	leastSteps(coords0, coords1, intersections)

}

func partOne() {

}
func parse(filePath string) []string {
	var path []string
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("err")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		path = strings.Split(val, ",")
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return path
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func stopwatch(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func buildCoordinateSet(pathSet []string) []Position {
	X := 0
	Y := 0
	var positions []Position
	for _,v := range pathSet {
		switch pos0 := strings.Split(v, "")[0]; pos0 {
		case UP:
			s,err := strconv.Atoi(strings.Replace(v, UP, "", -1))
			handleErr(err)
			for i := 0; i < s; i++ {
				Y += 1
				positions = append(positions, Position{X:X, Y:Y})
			}
		case DOWN:
			s,err := strconv.Atoi(strings.Replace(v, DOWN, "", -1))
			handleErr(err)
			for i := 0; i < s; i++ {
				Y -= 1
				positions = append(positions, Position{X:X, Y:Y})
			}
		case LEFT:
			s,err := strconv.Atoi(strings.Replace(v, LEFT, "", -1))
			handleErr(err)
			for i := 0; i < s; i++ {
				X -= 1
				positions = append(positions, Position{X:X, Y:Y})
			}
		case RIGHT:
			s,err := strconv.Atoi(strings.Replace(v, RIGHT, "", -1))
			handleErr(err)
			for i := 0; i < s; i++ {
				X += 1
				positions = append(positions, Position{X:X, Y:Y})
			}
		}

	}
	return positions
}

func findIntersections(wire0 []Position, wire1 []Position) []Position {
	var intersections []Position
	outer := len(wire0)
	inner := len(wire1)
	for i := 0; i < outer; i++ {
		for j := 0; j < inner; j++ {
			if(wire0[i] == wire1[j]) {
				intersections = append(intersections, wire0[i])
			}
 		}
	}
	return intersections
}

func manhattan(pos0 Position, pos1 Position) float64 {
	val := math.Abs(float64(pos0.GetX()) - float64(pos1.GetX())) + math.Abs(float64(pos0.GetY()) - float64(pos1.GetY()))
	return val
}

func findShortest(interestions []Position) float64 {
	center :=  Position{X:0, Y:0}
	min := math.MaxFloat64
	for _,v := range interestions {
		dist := manhattan(v, center)
		if min > dist {
			min = dist
		}
	}
	return min
}

func leastSteps(wire0 []Position, wire1 []Position, intersections []Position) {
	step0 := 0 
	step1 := 1
	steps := 0 
	min := math.MaxInt64
	for _, v := range intersections {
		for i, v0 := range wire0 {
			if (v == v0) {
				step0 = i+1
			}
		}
		for j, v1 := range wire1 {
			if (v == v1) {
				step1 = j+1
			}
		}
		steps = step0 + step1
		if (min > steps) {
			min = steps
		}
	}
	fmt.Println(min)
}