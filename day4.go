package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("lib/cleanupt.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	encapsulates := 0

	for fileScanner.Scan() {
		m := make(map[int]int)
		m[1] = 0
		m[2] = 0
		m[3] = 0
		m[4] = 0

		createAnInt := ""
		i := 1
		for _, char := range fileScanner.Text() {
			if string(char) != "-" && string(char) != "," {
				createAnInt = createAnInt + string(char)
				if i == 4 {
					val, _ := strconv.Atoi(createAnInt)
					m[4] = val
				}
			} else {
				val, _ := strconv.Atoi(createAnInt)
				m[i] = val
				createAnInt = ""
				i = i + 1
			}
		}

		/*
			Input is now in:
			Elf 1:
			m[1] = start
			m[2] = end

			Elf 2:
			m[3] = start
			m[4] = end
		*/

		if m[1] <= m[3] {
			if m[3] <= m[2] {
				encapsulates = encapsulates + 1
				fmt.Println(m)
			}
		} else if m[3] <= m[1] {
			if m[1] <= m[4] {
				encapsulates = encapsulates + 1
				fmt.Println(m)
			}
		} else {
			fmt.Println(m)
		}
	}
	fmt.Println(encapsulates)
	readFile.Close()
}
