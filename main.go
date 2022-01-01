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
	"path/filepath"
	"strconv"
)

const defaultFilename = "input.txt"

// Default day to run
var day = "day_02"

func main() {
	day = get_day()
	log.Print("AoC 2021 - " + day)

	switch day {
		case "day_01": day_01()
		case "day_01_part2": day_01(3)
		case "day_02": day_02()
		default: log.Fatal(day + " is not yet implemented")
	}

}

/**
 * Day 2
 */
func day_02() {
	log.Print("Not yet implemented")

	// Regex - match (forward|down|up) \d+

	// Update depth or horizontal_position accordingly

	// Output information and multiplied value
}

/**
 * Day 1
 */
func day_01(window_length_specified ...int) {

	window_length := 1

	if len(window_length_specified) > 0 {
		window_length = window_length_specified[0]
	}

	file := get_input_file()
	defer file.Close()

	increases := 0

	number_list:= make([]float64, 0)

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Read lines into an array of floats
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}

		number_list = append(number_list, number)
	}

	// Now we loop through the number list
	// - Each number is the start of a new window
	// - Compare with the next window
	// - Stop when we don't have enough to compare
	//   (X from the end - eg. index = length - X
	//    where X is window_length

	number_list_length := len(number_list)
	last_index := number_list_length - window_length

	var sum1 float64
	var sum2 float64

	for i := 0; i < last_index; i++ {
		sum1 = 0
		sum2 = 0
		for s1 := 0; s1 < window_length; s1++ {
			sum1 += number_list[i+s1]
		}
		for s2 := 0; s2 < window_length; s2++ {
			sum2 += number_list[i+s2+1]
		}

		log.Printf("Sums: %f, %f", sum1, sum2)

		if sum2 > sum1 {
			increases++
		}
	}

	log.Printf("Increases: %d", increases)
}

/**
 * Get the day from arguments if specified
 */
func get_day() string {
	if len(os.Args) > 1 {
		day = os.Args[1]
	}
	return day
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
	if len(os.Args) < 3 {
		return filepath.FromSlash(day + "/" + defaultFilename)
	}
	return os.Args[2]
}
