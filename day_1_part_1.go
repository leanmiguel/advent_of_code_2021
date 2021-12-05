package main

import (
	"fmt"
	"os"
	"strconv"
)

func parseArgs(args []string) []int {
	var err error
	nums := make([]int, len(args))
	for i := 0; i < len(args); i++ {
		if nums[i], err = strconv.Atoi(args[i]); err != nil {
			panic(err)
		}
	}

	return nums
}

func main() {

	commandArgs := os.Args[1:]
	data := parseArgs(commandArgs)

	increasedCount := 0

	for i := 0; i < len(data); i++ {

		currentVal := data[i]

		if i == 0 {
			continue
		}

		previousVal := data[i-1]
		increased := currentVal > previousVal

		if increased {
			increasedCount += 1
		}

	}

	fmt.Printf("There are %d measurements that are larger than the previous measurement\n", increasedCount)
}
