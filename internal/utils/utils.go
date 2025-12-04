package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// This function reads the wanted days file and returns it as a string slice split by linebreaks
func ReadDayInput(dayNumber uint) []string {
	fileName := fmt.Sprintf("%02d.txt", dayNumber)
	absPath, _ := filepath.Abs(filepath.Join("inputs", fileName))

	file, err := os.Open(absPath)
	if err != nil {
		log.Fatalf("Couldn't open file %s", absPath)
	}
	defer file.Close()

	r := bufio.NewReader(file)
	var output []string
	for {
		_, err := r.Peek(1)
		if err != nil {
			break
		}
		line, _, err := r.ReadLine()
		if err != nil {
			log.Fatalf("Something went wrong when reading line in input file: %v", err)
		}
		output = append(output, string(line))
	}
	return output
}

// ReplaceAtIndex swaps the rune at `index` in the string with `r`.
func ReplaceAtIndex(in string, r rune, index int) string {
	out := []rune(in)
	out[index] = r
	return string(out)
}
