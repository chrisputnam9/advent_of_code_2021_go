/**
 * Count lines of input in a text file where
 * - number is greater than the number on the previous line
 */
package main

import (
	"bufio"
	"errors"
	"log"
	"os"
)

const DefaultFilename = "input.txt"

func main() {
	log.Print("AoC 2021 - Day 1")

	file := get_input_file()
	defer file.Close()

	previous := nil
	increases := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		log.Print(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func get_input_file() *os.File {
	filepath := get_input_filepath()

	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		log.Fatal("ERROR: '" + filepath + "' does not exist. Create default "+DefaultFilename+" or pass other input filepath as first argument to script")
	}

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("File opened successfully")
	return file
}

/**
 * Get path from arguments or fall back to DefaultFilename
 */
func get_input_filepath() string {
	if len(os.Args) < 2 {
		return DefaultFilename
	}
	return os.Args[1]
}
