package phase1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func CapitalizeInput(input string) string {
	re := regexp.MustCompile(`(^|\.|\?|!)\s*(\w)`)
	capitalizedInput := re.ReplaceAllStringFunc(input, func(s string) string {
		return strings.ToUpper(s)
	})
	return capitalizedInput
}

func OrdinalizeNumbers(input string) string {
	re := regexp.MustCompile(`\d+(st|nd|rd|th)?`)
	ordinalizedNumbers := re.ReplaceAllStringFunc(input, func(s string) string {
		return ConvertToOrdinal(s)
	})
	return ordinalizedNumbers
}

func ConvertToOrdinal(s string) string {
	if strings.HasSuffix(s, "st") || strings.HasSuffix(s, "nd") || strings.HasSuffix(s, "rd") || strings.HasSuffix(s, "th") {
		return s
	}
	num, _ := strconv.Atoi(s)
	lastDigit := num % 10
	ordinalSuffix := "th"
	switch lastDigit {
	case 1:
		if num%100 != 11 {
			ordinalSuffix = "st"
		}
	case 2:
		if num%100 != 12 {
			ordinalSuffix = "nd"
		}
	case 3:
		if num%100 != 13 {
			ordinalSuffix = "rd"
		}
	}
	return s + ordinalSuffix
}

func FormatText(text string) string {
	capitalizedText := CapitalizeInput(text)
	result := OrdinalizeNumbers(capitalizedText)
	return result
}

func Run(inputFilePath string, outputFilePath string) {
	startTime := time.Now()
	inputFile, _ := os.Open(inputFilePath)
	defer inputFile.Close()

	outputFile, _ := os.Create(outputFilePath)
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)

	for scanner.Scan() {
		text := scanner.Text()
		result := FormatText(text)
		writer.WriteString(result + "\n")
	}

	writer.Flush()
	elapsed := time.Since(startTime)
	fmt.Println("total time taken ", elapsed.Seconds(), "seconds")
}
