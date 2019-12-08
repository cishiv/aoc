package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type image struct {
	length int
	height int
	layers []layer
}

type layer struct {
	pixels []int
	vals   map[int]int
}

func main() {
	data := readFile("./input-day8")
	im := buildImage(data, 25, 6)
	partOne(im, "part1")
	partTwo(im, "part2")

}

func buildImage(data []int, length int, height int) image {
	var layers []layer
	prev := 0
	val := len(data)
	im := image{length: length, height: height, layers: layers}
	for i := 0; i <= val; i += im.length * im.height {
		layerMap := make(map[int]int)
		// fmt.Println(strconv.Itoa(prev) + " " + strconv.Itoa(i))
		layeri := data[prev:i]
		var pixels []int
		for _, v := range layeri {
			layerMap[v] = layerMap[v] + 1
			pixels = append(pixels, v)
		}
		layero := layer{pixels: pixels, vals: layerMap}
		im.layers = append(im.layers, layero)
		// fmt.Println(layeri)
		prev = i
	}
	return im
}

func partOne(im image, part string) {
	defer stopwatch(time.Now(), part)
	least := math.MaxInt64
	fewest := 0
	for i, v := range im.layers {
		// fmt.Println(i)
		for range v.vals {
			if v.vals[0] < least {
				fewest = i
				least = v.vals[0]
			}
		}
	}
	fmt.Println(im.layers[fewest].vals[1] * im.layers[fewest].vals[2])
}

func partTwo(im image, part string) {
	defer stopwatch(time.Now(), part)
	decode(parse(im), im.height, im.length)
}

func parse(im image) []int {

	resultIm := make([]int, im.length*im.height)

	for i := range resultIm {
		resultIm[i] = 2
	}

	for k := len(im.layers) - 1; k >= 0; k-- {
		for i := 0; i < len(im.layers[k].pixels); i++ {
			if !(im.layers[k].pixels[i] == 2) {
				resultIm[i] = im.layers[k].pixels[i]
			}
		}
	}

	return resultIm

}

func decode(imRep []int, rows int, cols int) {
	pntr := cols
	pntrFirst := 0
	for i := 0; i < rows; i++ {
		for _, v := range imRep[pntrFirst:pntr] {
			if v == 1 {
				fmt.Print(" ")
			} else {
				fmt.Print("█")
			}
		}
		pntrFirst = pntr
		pntr += cols
		fmt.Println()
	}
}

// -- UTILS

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
		codes := strings.Split(val, "")
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
