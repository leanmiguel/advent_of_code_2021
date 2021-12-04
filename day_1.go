package main

import (
	"fmt"
)

func main() {

	data := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	increasedCount := 0

	for i := 0; i < len(data); i++ {

		currentVal := data[i]

		if i == 0 {
			fmt.Printf("%d (N/A - no previous measurement)\n", currentVal)
			continue
		}

		previousVal := data[i-1]
		increased := currentVal > previousVal
		statusString := "decreased"

		if increased {
			increasedCount += 1
			statusString = "increased"
		}

		fmt.Printf("%d (%s)\n", currentVal, statusString)

	}
}
