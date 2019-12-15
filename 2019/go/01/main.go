package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func fuel(mass int) int {
	return int(mass/3) - 2;
}

func sumArray(array []int) int {
	var sum int;
	for _, element := range array {
		sum += element;
	}
	return sum;
}

func main() {
	// Read in the puzzle input
	file, _ := ioutil.ReadFile("input");
	var temp = strings.Split(string(file), "\n");
	var moduleMasses []int;
	for _, str := range temp {
		var mass, _ = strconv.Atoi(str);
		moduleMasses = append(moduleMasses, mass);
	}
	// Part 1
	var fuelRequirements []int;
	for _, mass := range moduleMasses {
		fuelRequirements = append(fuelRequirements, fuel(mass));
	}
	fmt.Println("Part 1:", sumArray(fuelRequirements));
	// Part 2
	var newFuelRequirements []int;
	for _, mass := range moduleMasses {
		var moduleFuel = 0;
		var currentMass = mass;
		for currentMass > 0 {
			currentMass = fuel(currentMass);
			if currentMass > 0 {
				moduleFuel += currentMass;
			}
		}
		newFuelRequirements = append(newFuelRequirements, moduleFuel);
	}
	fmt.Println("Part 2:", sumArray(newFuelRequirements));
}