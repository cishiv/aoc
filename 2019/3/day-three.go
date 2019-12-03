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
	fmt.Println(findShortest(intersections))

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
			Y += s
			positions = append(positions, Position{X:X, Y:Y})
		case DOWN:
			s,err := strconv.Atoi(strings.Replace(v, DOWN, "", -1))
			handleErr(err)
			Y -= s
			positions = append(positions, Position{X:X, Y:Y})
		case LEFT:
			s,err := strconv.Atoi(strings.Replace(v, LEFT, "", -1))
			handleErr(err)
			X -= s
			positions = append(positions, Position{X:X, Y:Y})
		case RIGHT:
			s,err := strconv.Atoi(strings.Replace(v, RIGHT, "", -1))
			handleErr(err)
			X += s
			positions = append(positions, Position{X:X, Y:Y})
		}

	}
	return positions
}

func findIntersections(wire0 []Position, wire1 []Position) []Position {
	var intersections []Position
	outer := len(wire0) - 1
	inner := len(wire1) - 1
	for i := 0; i < outer; i++ {
		for j := 0; j < inner; j++ {
			if(intersect(wire0[i], wire0[i+1], wire1[j], wire1[j+1])) {
				pos := findPoint(wire0[i], wire0[i+1], wire1[j], wire1[j+1])
				if (pos.GetX() == 0 && pos.GetY() == 0) {

				} else {
					intersections = append(intersections, pos)
				}
			}
		}
	}
	return intersections
}

func onSegment(p Position, q Position, r Position) bool {
	if (float64(q.GetX()) <= math.Max(float64(p.GetX()), float64(r.GetX())) && float64(q.GetX()) >= math.Min(float64(p.GetX()), float64(r.GetY())) && float64(q.GetY()) <= math.Max(float64(p.GetY()), float64(r.GetY())) && float64(q.GetY()) >= math.Min(float64(p.GetY()), float64(r.GetY()))) {
		return true
	}
	return false
}

func orientation(p Position, q Position, r Position) int {
	val := (q.GetY() - p.GetY()) * (r.GetX() - q.GetX()) - (q.GetX() - p.GetX()) * (r.GetY() - q.GetY())
	if val == 0 {
		return 0
	}
	if val < 0 {
		return 2
	}
	if val > 0 {
		return 1
	}
	return -1
}

func intersect(p1 Position, p2 Position, q1 Position, q2 Position) bool {
	o1 := orientation(p1, q1, p2)
	o2 := orientation(p1, q1, q2)
    o3 := orientation(p2, q2, p1) 
	o4 := orientation(p2, q2, q1) 
	
	if (o1 != o2 && o3 != o4) {
		return true
	}

    if (o1 == 0 && onSegment(p1, p2, q1)) {
		return true
	}

    if (o2 == 0 && onSegment(p1, q2, q1)) {
		return true
	} 
  
    if (o3 == 0 && onSegment(p2, p1, q2)) {
		return true; 
	}

    if (o4 == 0 && onSegment(p2, q1, q2)) {
		return true
	}

    return false
}

func findPoint(A Position, B Position, C Position, D Position) Position { 
	a1 := B.GetY() - A.GetY()
	b1 := A.GetX() - B.GetX() 
	c1 := a1*(A.GetX()) + b1*(A.GetY())
   
	// Line CD represented as a2x + b2y = c2 
	a2 := D.GetY() - C.GetY() 
	b2 := C.GetX() - D.GetX() 
	c2 := a2*(C.GetX())+ b2*(C.GetY())
   
	determinant := a1*b2 - a2*b1
	if (determinant != 0) {
		x := (b2*c1 - b1*c2)/determinant
		y := (a1*c2 - a2*c1)/determinant
		return Position{X:x, Y:y} 
	} else {
		return Position{X:0, Y:0}
	}
} 

func manhattan(pos0 Position, pos1 Position) float64 {
	val := math.Abs(float64(pos0.GetX()) - float64(pos1.GetX())) + math.Abs(float64(pos0.GetY()) - float64(pos1.GetY()))
	return val
}

func findShortest(interestions []Position) float64 {
	center :=  Position{X:0, Y:0}
	min := math.MaxFloat64
	for _,v := range interestions {
		dist := manhattan(center, v)
		if min > dist {
			min = dist
		}
	}
	return min
}