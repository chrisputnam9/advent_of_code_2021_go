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
	"regexp"
	"strconv"
)

const defaultFilename = "input.txt"

// Default day to run
var day = "day_02_part2"

func main() {
	day = get_day()
	log.Print("AoC 2021 - " + day)

	switch day {
		case "day_01": day_01()
		case "day_01_part2": day_01(3)
		case "day_02": day_02()
		case "day_02_part2": day_02(true)
		case "day_03": day_03()
		default: log.Fatal(day + " is not yet implemented.  Specify a day argument such as 'day_02' or 'day_01_part2'")
	}

}

/**
 * Day 3
 */
func day_03() {

	file := get_input_file()
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	gamma := ""
	epsilon := ""

	// https://go.dev/tour/moretypes/13
	ones := make([]int, 0)
	zeros := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}

		// https://golangcookbook.com/chapters/strings/processing/
		log.Print(line);
	}

	log.Printf("Gamma Binary: %s", gamma)
	log.Printf("Epsilon Binary: %s", epsilon)
}

/**
 * Day 2
 */
func day_02(use_aim_specified ...bool) {

	use_aim := false
	if len(use_aim_specified) > 0 {
		use_aim = use_aim_specified[0]
	}

	// Regex to match navigation commands
	regex, err := regexp.Compile(`(?i)^\s*(forward|down|up)\s*(\d+)\s*$`)
	if err != nil {
		log.Fatal(err)
	}

	file := get_input_file()
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	aim := 0
	depth := 0
	horizontal_position := 0

	// Read lines and parse commands
	for scanner.Scan() {
		command := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}

		// Update depth or horizontal_position accordingly
		matches := regex.FindStringSubmatch(command)
		if len(matches) == 3 {

			amount, err := strconv.Atoi(matches[2])
			if err != nil {
				log.Fatal(err)
			}

			switch matches[1] {
				case "forward":
					horizontal_position+= amount
					if use_aim {
						depth+= (amount * aim)
					}
				case "down":
					if use_aim {
						aim+= amount
					} else {
						depth+= amount
					}
				case "up":
					if use_aim {
						aim-= amount
					} else {
						depth-= amount
					}
			}
		} else {
			log.Fatal("Invalid command: " + command)
		}

		log.Printf("Command: %s, A: %d, H:%d, D:%d", command, aim, horizontal_position, depth)
	}

	// Output information and multiplied value
	log.Print("-----------------------------------")
	log.Printf("Depth: %d", depth)
	log.Printf("Horizontal Position: %d", horizontal_position)
	log.Printf("Multiplied: %d", depth * horizontal_position)
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
