/**
 * Count lines of input in a text file where
 * - number is greater than the number on the previous line
 */
package main

import (
	"bufio"
	"errors"
	"log"
	"strconv"
	"os"
)

const defaultFilename = "input.txt"

func main() {
	log.Print("AoC 2021 - Day 1")

	file := get_input_file()
	defer file.Close()

	first_line := true
	var previous float64 = -1 // Probably could be int based on input, but not detailed in specs
	var increases int = 0

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		current, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}

		if ! first_line {
			if current > previous {
				increases++
			}
		}

		previous = current
		first_line = false
	}

	log.Printf("Increases: %d", increases)

}

func get_input_file() *os.File {
	filepath := get_input_filepath()

	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		log.Fatal("ERROR: '" + filepath + "' does not exist. Create default "+defaultFilename+" or pass other input filepath as first argument to script")
	}

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("File opened successfully")
	return file
}

/**
 * Get path from arguments or fall back to defaultFilename
 */
func get_input_filepath() string {
	if len(os.Args) < 2 {
		return defaultFilename
	}
	return os.Args[1]
}
