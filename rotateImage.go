package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("USAGE : %s <target_filename> \n", os.Args[0])
		os.Exit(0)
	}

	fileName := os.Args[1]

	fileBytes, err := ioutil.ReadFile(fileName)
	log.Printf("Processing data from file...")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(fileBytes), "\n")

	validateFileData(sliceData)
	inputArray := [10][6]int{}
	// arrayDimension := strings.Split(sliceData[2], " ")

	for i, value := range sliceData {

		jValue := strings.Split(value, " ")

		for k := 0; k < len(jValue); k++ {
			number, _ := strconv.Atoi(jValue[k])
			inputArray[i][k] = number
		}
	}

	fmt.Printf("%v\n", inputArray)
}

func validateFileData(data []string) string {
	log.Printf("Validating data from file...")
	// for i, value := range data {
	// 	// fmt.Println(i, value)
	// }
	return ""
}
