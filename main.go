package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type questionaire struct {
	ques string
	ans string
}

func main() {
	fmt.Println("Welcome to the Quiz app")

	var totalQuestions int
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
		lines,_ := r.ReadAll()
		listOfQuestions := createQuestionaire(lines)
		totalQuestions = len(listOfQuestions)
		
		for index, question := range listOfQuestions{
			// read a question and print it 
			fmt.Printf("Ques %v: %v \n",(index+1), question.ques)

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
				correctResult := formatComparison(question.ans)

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

func createQuestionaire(lines [][]string) []questionaire {
	var questions []questionaire

	for _,line := range lines {
		question := questionaire{ques: line[0], ans: line[1]}
		questions = append(questions, question)
	}
	return questions
}