package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	data := readFile("./input-day6")
	partOne(createEdges(data))
	partTwo(createEdges(data))
}

func partOne(edges map[string]string) {
	defer stopwatch(time.Now(), "part1")

	total := 0
	for o := range edges {
		for k, v := edges[o]; v; k, v = edges[k] {
			total++
		}
	}
	fmt.Println(total)
}

func partTwo(edges map[string]string) {
	defer stopwatch(time.Now(), "part2")

	you := map[string]int{}
	for k, v := edges["YOU"]; v; k, v = edges[k] {
		// Found the number of orbital transfers to you
		you[k] = len(you)
	}

	santa := 0
	for k, v := edges["SAN"]; v; k, v = edges[k] {
		// Huzzah!
		if _, v := you[k]; v {
			fmt.Println(santa + you[k])
			break
		}
		santa++
	}
}

func createEdges(data []string) map[string]string {
	edges := make(map[string]string)
	// configure edges
	for _, k := range data {
		vals := strings.Split(k, ")")
		edges[vals[1]] = vals[0]
	}
	return edges
}

// -- UTIL --

func readFile(filePath string) []string {
	file, err := os.Open(filePath)
	var data []string
	if err != nil {
		fmt.Printf("err")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		data = append(data, val)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func stopwatch(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
