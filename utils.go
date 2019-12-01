package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func parse(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("err")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		fmt.Println(val)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func handleErr(err error) {
	fmt.Println("Error")
}

func stopwatch(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
