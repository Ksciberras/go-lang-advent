package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var num_map = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file %v", err)
	}

	lines_of_string := strings.Split(string(dat), "\n")
	sum := 0

	for _, line := range lines_of_string {

		numbers_and_letters := strings.Split(line, "")
		numbers := []string{}
		content := []byte(line)
		regex := regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine)")
		indexes := regex.FindAllIndex(content, -1)
		if len(indexes) != 0 {
			for _, index := range indexes {
				word := line[index[0]:index[1]]
				numbers_and_letters = insert(numbers_and_letters, index[0], num_map[word])
			}
		}

		for _, number_or_letter := range numbers_and_letters {
			number, err := strconv.Atoi(number_or_letter)
			if err == nil {
				numbers = append(numbers, strconv.Itoa(number))
			}
		}
		if len(numbers) == 0 {
			continue
		}
		numbers_slice := []string{numbers[0], numbers[len(numbers)-1]}

		final_number := join_numbers_and_convert_to_int(numbers)
		fmt.Println(numbers_slice, final_number, line)

		sum = sum + final_number
	}
	fmt.Println(sum)
}

func join_numbers_and_convert_to_int(numbers []string) int {
	if len(numbers) == 0 {
		return 0
	}

	numbers_slice := []string{numbers[0], numbers[len(numbers)-1]}
	joined_numbers := strings.Join(numbers_slice, "")
	final_number, _ := strconv.Atoi(joined_numbers)

	return final_number
}

func insert(slice []string, index int, value string) []string {
	if index < 0 || index > len(slice) {
		fmt.Println("Index out of range")
		return slice
	}

	slice = append(slice, "")
	copy(slice[index+1:], slice[index:])
	slice[index] = value
	return slice
}
