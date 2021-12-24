/**
 * Count lines of input in a text file where
 * - number is greater than the number on the previous line
 */
package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	log.Print("AoC 2021 - Day 1")

	file := get_input_file()
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		log.Print(scanner.Text())
	}

	if error := scanner.Err(); error != nil {
		log.Fatal(error)
	}
}

func get_input_file() *os.File {
	filepath := get_input_filepath()

	file, error := os.Open(filepath)
	if error != nil {
		log.Fatal(error)
	}

	log.Print("File opened successfully")
	return file
}

func get_input_filepath() string {
	if len(os.Args) < 2 {
		log.Fatal("ERROR: Pass input filepath as first argument to script")
	}
	return os.Args[1]
}
