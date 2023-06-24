package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to the Quiz app")

	filename := checkFileFlag()
	
	// open the csv file 
	file, fileErr := os.Open(filename)

	if fileErr != nil {
		if(fileErr == os.ErrNotExist){
			fmt.Println("File Name doesn't exist")
		}
		fmt.Println("Error opening  file", fileErr)
	}

	// close the file 
	defer file.Close()

	// reader for the csv file
	r := csv.NewReader(file)

	totalQustions := 0
	correctAnswers := 0

	for{
		// read a line and print it 
		line, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error reading the file")
		}
		
		totalQustions += 1
		var userAnswer string

		//
		fmt.Printf("Ques %v: %v \n",totalQustions, line[0])
		fmt.Printf("Answer : ")
		fmt.Scan(&userAnswer)

		if(userAnswer == line[1]){
			correctAnswers +=1
		}
		
		fmt.Printf("Correct Answers : %v\n", correctAnswers)
	}
}


func checkFileFlag() string{
	// check for filename flag for questionaire 
	fileFlag := flag.String("f", "problems.csv", "Change the file for the questionaire")
	flag.Parse()
	fmt.Printf("%v\n", *fileFlag)
	return *fileFlag
}
