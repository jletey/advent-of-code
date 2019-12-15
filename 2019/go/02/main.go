package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func runProgram(program []int) []int {
	var position = 0;
	for program[position] != 99 {
		var operation = program[position];
		var num1 = program[program[position + 1]];
		var num2 = program[program[position + 2]];
		if operation == 1 {
			var sum = num1 + num2;
			program[program[position + 3]] = sum;
		}
		if operation == 2 {
			var mult = num1 * num2;
			program[program[position + 3]] = mult;
		}
		position += 4;
	}
	return program;
}

func main() {
	// Read in the puzzle input
	file, _ := ioutil.ReadFile("input");
	var temp = strings.Split(string(file), ",");
	var program []int;
	for _, str := range temp {
		var num, _ = strconv.Atoi(str);
		program = append(program, num);
	}
	// Part 1
	program[1] = 12;
	program[2] = 2;
	var output = runProgram(program);
	fmt.Println("Part 1:", output[0]);
}