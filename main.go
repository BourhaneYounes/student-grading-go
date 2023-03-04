package main

import (
	"bufio"
	"fmt"
	_ "io"
	"os"
	"strconv"
	"strings"
)
type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func check(e error){
	if e!= nil {
		panic(e)
	}
}
func convert(str string) int{
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

func parseCSV(filePath string) []student {
	students := make([]student, 0)
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for scanner.Scan(){
		line := strings.Split(scanner.Text(), ",")
		students = append(students, student{firstName: line[0],
							lastName: line[1],	
							university: line[2],	
							test1Score: convert(line[3]),
							test2Score: convert(line[4]),
							test3Score: convert(line[5]),
							test4Score: convert(line[6]),
		})
	}

	return students 
}

func calculateGrade(students []student) []studentStat {
	return nil
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	return studentStat{}
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	return nil
}

func main(){
	fmt.Println(parseCSV("./grades.csv"))
}
