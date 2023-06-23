package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the Quiz app")

	// open the csv file 
	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Error opening  file", err)
	}

	// close the file 
	defer file.Close()


	// reader for the csv file
	r := csv.NewReader(file)

	// read a line and pirnt it 
	line, err := r.Read()
	if err != nil {
		fmt.Println("error reading the file")
	}

	fmt.Println(line)

	
}