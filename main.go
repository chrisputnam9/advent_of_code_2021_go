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
var day = "day_03_part2"

func main() {
	day = get_day()
	log.Print("AoC 2021 - " + day)

	switch day {
		case "day_01": day_01()
		case "day_01_part2": day_01(3)
		case "day_02": day_02()
		case "day_02_part2": day_02(true)
		case "day_03": day_03("power_consumption")
		case "day_03_part2": day_03("life_support")
		default: log.Fatal(day + " is not yet implemented.  Specify a day argument such as 'day_02' or 'day_01_part2'")
	}

}

/**
 * Day 3
 */
func day_03(calculation string) {

	file := get_input_file()
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	number_list := make([][]int, 0)

	// Loop over input, counting 0s and 1s in each position
	//  - and adding lines to a slice for future reference
	for scanner.Scan() {
		line := scanner.Text()

		digits := make([]int, 0)

		for _, char := range line {
			switch char {
				case '0':
					digits = append(digits, 0)
				case '1':
					digits = append(digits, 1)
			}
		}

		number_list = append(number_list, digits)
	}

	switch (calculation) {
		case "power_consumption": day_03_power_consumption(number_list)
		case "life_support": day_03_life_support(number_list)
		default: log.Fatalf("%s calculation not yet implemented", calculation)
	}

}
func day_03_life_support(number_list [][]int) {

	oxygen_list := make([][]int, 0)
	copy(oxygen_list, number_list)
	co2_list := make([][]int, 0);
	copy(co2_list, number_list)

	oxygen:= ""
	co2 := ""

	// Loop through each position
	for i := 0; i < len(number_list[0]); i++ {

		zeros, ones := day_03_count_bits(number_list)
		count_zeros := zeros[i]
		count_ones := ones[i]

		oxygen_bit_keep := 1
		co2_bit_keep := 0

		// Check for most common bit
		if ( count_zeros > count_ones ) {
			oxygen_bit_keep = 0
			co2_bit_keep = 1
		} // otherwise, stick with default 1/0

		// TODO
		// We could improve this by stopping once we have only one
		oxygen_list = day_03_life_support_filter(oxygen_list, i, oxygen_bit_keep)
		co2_list = day_03_life_support_filter(co2_list, i, co2_bit_keep)

	}

	// Finalize
	// TODO

}

/**
 * Count the bits at each position and return total
 * @return zeros, ones
 */
func day_03_count_bits(number_list [][]int) ([]int, []int) {
	zeros := make([]int, 0)
	ones := make([]int, 0)

	for _, digit_list := range number_list {
		for i, digit := range digit_list {

			// Expand our slices if needed
			if ( len(ones) < (i + 1) ) {
				zeros = append(zeros, 0)
				ones = append(ones, 0)
			}

			switch digit {
				case 0:
					zeros[i]++
				case 1:
					ones[i]++
			}

			log.Printf("Digits: %v, Zeroes: %v, Ones: %v", digit_list, zeros, ones);
		}
	}

	return zeros, ones
}
/**
 * Filter number list based on a given position and what bit signals to keep that number
 * - eg. for position 1, bit 0 - keep all numbers with a 0 in position 1
 */
func day_03_life_support_filter(number_list [][]int, position, bit_to_keep int) [][]int {
	filtered := make([][]int, 0)

	for _, number := range number_list {
		digit := number[position]
		if digit == bit_to_keep {
			filtered = append(filtered, number)
		}
	}

	return filtered
}
/**
 * Day 3 part 2 - calculate power consumption
 */
func day_03_power_consumption(number_list [][]int){

	zeros, ones := day_03_count_bits(number_list)

	gamma := ""
	epsilon := ""

	// Now Build Gamma and Epsilon by getting most and least common bit in each position
	for i := 0; i < len(zeros); i++ {
		count_zeros := zeros[i]
		count_ones := ones[i]

		if ( count_zeros > count_ones ) {
			gamma += "0" // Most common bit
			epsilon += "1" // Least common bit
		} else if ( count_ones > count_zeros ) {
			gamma += "1" // Most common bit
			epsilon += "0" // Least common bit
		} else {
			log.Fatalf("There are the same number of ones as there are zeros at position %d - this scenario is not accounted for in the specs", i)
		}
	}

	gamma_decimal, err := strconv.ParseUint( gamma, 2, 32)
	if err != nil {
		log.Fatal(err)
	}

	epsilon_decimal, err := strconv.ParseUint( epsilon, 2, 32)
	if err != nil {
		log.Fatal(err)
	}

	multiplied := gamma_decimal * epsilon_decimal

	// Output information and multiplied value
	log.Print("-----------------------------------")
	log.Printf("Gamma Binary: %s", gamma)
	log.Printf("Epsilon Binary: %s", epsilon)
	log.Print("-----------------------------------")
	log.Printf("Gamma Decimal: %d", gamma_decimal)
	log.Printf("Epsilon Decimal: %d", epsilon_decimal)
	log.Print("-----------------------------------")
	log.Printf("Multiplied: %d", multiplied)
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
