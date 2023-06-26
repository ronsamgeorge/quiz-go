package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to the Quiz app")

	totalQuestions := 0
	correctAnswers := 0

	answerChan := make(chan string)

	questionsFile, timer := checkFlags()
	// open the csv file 
	file, fileErr := os.Open(questionsFile)
	fileExists := checkFileExists(fileErr)
	
	if (fileExists){
		// reader for the csv file
		r := csv.NewReader(file)

		quizTimer := time.NewTimer( time.Duration(timer) * time.Second)
		
		for{
			// read a question and print it 
			line, err := r.Read()
			if err != nil {
				// exit if end of file
				if err == io.EOF {
					break
				}
				fmt.Println("error reading the file")
			}
			
			totalQuestions += 1 	// increment the question count
			
			
			fmt.Printf("Ques %v: %v \n",totalQuestions, line[0])
			// take input from user as a go routine
			go getUserAnswer(answerChan)
			select {
			case <- quizTimer.C:
				fmt.Println(" \n Your time is up")
				displayResult(totalQuestions, correctAnswers)
				return
			
			case userAnswerInput := <- answerChan :
				// format the input and the file answer for comparison
				userResult := formatComparison(userAnswerInput)
				correctResult := formatComparison(line[1])

				if(userResult == correctResult){
					correctAnswers +=1
				}
			}
		}
	}
	displayResult(totalQuestions, correctAnswers)
	// close the file 
	defer file.Close()
}


func checkFlags() (string, int) {
	// check for filen provided using the flag , for questionaire
	// if no flag provided, default to problems.csv
	fileFlag := flag.String("f", "problems.csv", "Change the file for the questionaire")
	timeFlag := flag.Int("t", 30, "set time for quiz")
	flag.Parse()
	return *fileFlag , *timeFlag
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

func displayResult( totalQues int, correctAns int) {
	fmt.Println("##########################")
	fmt.Printf("Your SCORE : %v / %v \n", correctAns, totalQues)
	fmt.Println("##########################")
}

func getUserAnswer(answerChan chan string){
	var userAnswerInput string
	fmt.Printf("Your Answer here : ")
	fmt.Scan(&userAnswerInput)

	answerChan <- userAnswerInput
}