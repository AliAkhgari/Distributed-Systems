package main

import (
	// "CA1/phase1"
	"CA1/phase2"
	"os"
	"strconv"
)

func main() {
	// phase1.Run("input.txt", "output.txt")

	n, _ := strconv.Atoi(os.Args[1])
	phase2.Run("input.txt", "output.txt", n)
}
