package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Quiz app")

	questionsFile := checkFileFlag()
	// open the csv file 
	file, fileErr := os.Open(questionsFile)
	fileExists := checkFileExists(fileErr)
	
	if (fileExists){
		// reader for the csv file
		r := csv.NewReader(file)
		totalQustions := 0
		correctAnswers := 0
		for {
			// read a question and print it 
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
			fmt.Printf("Your Answer Answer : ")
			fmt.Scan(&userAnswer)

			userResult := formatComparison(userAnswer)
			correctResult := formatComparison(line[1])

			if(userResult == correctResult){
				correctAnswers +=1
			}
		}
	}

	// close the file 
	defer file.Close()
}


func checkFileFlag() string {
	// check for filen provided using the flag , for questionaire
	// if no flag provided, default to problems.csv
	fileFlag := flag.String("f", "problems.csv", "Change the file for the questionaire")
	flag.Parse()
	return *fileFlag
}

func checkFileExists(fileErr error) bool {
	if fileErr != nil {
		fmt.Println("Error opening  file", fileErr)
		if(fileErr == os.ErrNotExist){
			fmt.Println("File Name doesn't exist")
		}
		return false
	}
	return true
}

func formatComparison(result string) string{
	returnResult := strings.ToLower(strings.Trim(result, " ")) 
	return returnResult
}