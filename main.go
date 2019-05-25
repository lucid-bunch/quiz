package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, fileErr := os.Open("quiz.csv")
	if fileErr != nil {
		fmt.Println(fileErr)
	}
	defer file.Close()

	lines, linesErr := csv.NewReader(file).ReadAll()
	if linesErr != nil {
		fmt.Println(lines)
	}

	reader := bufio.NewReader(os.Stdin)
	correctAnswers := int(0)
	for _, line := range lines {
		question, answer := line[0], line[1]
		fmt.Printf("%s?:", question)
		input, _, inputErr := reader.ReadLine()
		if inputErr != nil {
			fmt.Println(inputErr)
		}
		if strings.EqualFold(string(input), answer) {
			correctAnswers++
		}
	}
	fmt.Printf("You got %d correct out of %d\n", correctAnswers, len(lines))
}
