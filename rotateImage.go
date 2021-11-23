package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"path/filepath"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("only one input file is required")
	}

	fileName := os.Args[1]
	fileExt := filepath.Ext(fileName)

	if fileExt != ".pbm" {
		log.Fatal("Provide only pbm file as an input")
	}

	fileBytes, err := ioutil.ReadFile(fileName)
	log.Printf("Processing data from file...")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(fileBytes), "\n")
	validatedData, _, noOfCols := validateFileData(sliceData)

	var inputSlice [][]int
	for i:= 2; i < len(validatedData) ; i++ {
		is := []int{}

		jValue := strings.Split(validatedData[i], " ")
		if len(jValue) > noOfCols {
			log.Fatal("Metric is not in correct size")
		} else {
			for k := 0; k < len(jValue); k++ {
				number, _ := strconv.Atoi(jValue[k])
				is = append(is, number)
			}
		}
		if len(is) > 0 {
			inputSlice = append(inputSlice, is)
		}
	}

	//printMatrix(inputSlice, len(inputSlice), len(inputSlice[0]))

	output := rotate(inputSlice, len(inputSlice), len(inputSlice[0]))

	writeDataToFile(output)
}

func printMatrix(arr [][]int, rows int, columns int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Printf("%v\t", arr[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

// rotate MxN metric
func rotate(matrix [][]int, rows int, columns int) [][]int{
	log.Printf("Processing metric data...")
	var rotatedMetric [][]int

	log.Printf("Generating rotating metrix...")
	for j := 0; j < columns; j++ {
		var temp []int
		for i := 0; i < rows; i++ {
			//fmt.Printf("%v\t", matrix[i][j])
			temp = append(temp, matrix[i][j])
		}
		//fmt.Println()
		rotatedMetric = append(rotatedMetric, temp)
	}
	return rotatedMetric
}

func validateFileData(data []string) ([]string, int, int) {
	log.Printf("Formatting data...")
	formattedData := removeEmptyStrings(data)
	var rows int
	var cols int

	log.Printf("Validating data...")

	if len(formattedData) == 0 {
		log.Fatal("Received empty file")
	}

	if formattedData[0] != "P1" {
		log.Fatal("Not a valid format, missing P1")
	}

	arrayDimension := strings.Split(formattedData[1], " ")
	if len(arrayDimension) > 2 {
		log.Fatal("it's looks like there is an error in matrix size declaration")
	} else {
		rows, _ = strconv.Atoi(arrayDimension[1])
		cols,_ = strconv.Atoi(arrayDimension[0])
		}
	return formattedData, rows, cols
}

func writeDataToFile(outputMetrix [][]int) {
	log.Printf("Writing data to new file")
	rows := len(outputMetrix)
	cols := len(outputMetrix[0])
	f, err := os.Create("pbmoutputfile.pbm")

	if err != nil {
		log.Fatal(err)
	}

	if _, err = f.WriteString("P1\n"); err != nil {
		log.Fatal(err)
	}

	if _, err = f.WriteString(strconv.Itoa(cols) +" " +strconv.Itoa(rows)+"\n"); err != nil {
		log.Fatal(err)
	}

	for _, value := range outputMetrix {
		_, err := fmt.Fprintln(f, value)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("Finished the process, new file %v is avaliable now" , f.Name())
}

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" || len(str) > 0 {
			if !strings.HasPrefix(str, "#") {
				r = append(r, str)
			}
		}
	}
	return r
}
