package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	file, err := os.Open("./input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	horizontalPosition := 0
	depth := 0

	for scanner.Scan() {

		text := strings.Split(scanner.Text(), " ")
		action := text[0]
		movementValue, err := strconv.Atoi(text[1])
		checkError(err)

		if action == "forward" {
			horizontalPosition += movementValue
		} else if action == "up" {
			depth -= movementValue
		} else {
			depth += movementValue
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("the depth is %d, the horizontal position is %d", depth, horizontalPosition)
}
