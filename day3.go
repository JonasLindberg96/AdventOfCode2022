package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	p := make(map[string]int)

	p["a"] = 1
	p["b"] = 2
	p["c"] = 3
	p["d"] = 4
	p["e"] = 5
	p["f"] = 6
	p["g"] = 7
	p["h"] = 8
	p["i"] = 9
	p["j"] = 10
	p["k"] = 11
	p["l"] = 12
	p["m"] = 13
	p["n"] = 14
	p["o"] = 15
	p["p"] = 16
	p["q"] = 17
	p["r"] = 18
	p["s"] = 19
	p["t"] = 20
	p["u"] = 21
	p["v"] = 22
	p["w"] = 23
	p["x"] = 24
	p["y"] = 25
	p["z"] = 26
	p["A"] = 27
	p["B"] = 28
	p["C"] = 29
	p["D"] = 30
	p["E"] = 31
	p["F"] = 32
	p["G"] = 33
	p["H"] = 34
	p["I"] = 35
	p["J"] = 36
	p["K"] = 37
	p["L"] = 38
	p["M"] = 39
	p["N"] = 40
	p["O"] = 41
	p["P"] = 42
	p["Q"] = 43
	p["R"] = 44
	p["S"] = 45
	p["T"] = 46
	p["U"] = 47
	p["V"] = 48
	p["W"] = 49
	p["X"] = 50
	p["Y"] = 51
	p["Z"] = 52

	readFile, err := os.Open("lib/rucksack.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	prioPoint := 0
	commonChars := ""
	groupPrioPoint := 0
	i := 1

	first := ""
	second := ""
	third := ""

	for fileScanner.Scan() {
		halfStringLenght := len(fileScanner.Text()) / 2
		compartmentOne := fileScanner.Text()[0:halfStringLenght]
		compartmentTwo := fileScanner.Text()[halfStringLenght:]

		for _, char := range compartmentOne {
			if strings.ContainsAny(compartmentTwo, string(char)) {
				if !strings.ContainsAny(commonChars, string(char)) {
					commonChars = string(commonChars + string(char))
					prioPoint = prioPoint + p[string(char)]
				}
			}
		}
		commonChars = ""

		if i%3 == 0 {
			third = fileScanner.Text()
			for _, char := range first {
				if strings.ContainsAny(second, string(char)) {
					if strings.ContainsAny(third, string(char)) {
						groupPrioPoint = groupPrioPoint + p[string(char)]
						break
					}
				}
			}
			first = ""
			second = ""
			third = ""
			i = 1

		} else {
			if i == 1 {
				first = fileScanner.Text()
			} else if i == 2 {
				second = fileScanner.Text()
			}
			i = i + 1
		}
	}
	readFile.Close()
	fmt.Println("Total Points part one: ", prioPoint)
	fmt.Println("Group points: ", groupPrioPoint)
}
