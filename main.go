package main

import (
	"bufio"
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

func average(score ...int) float32{
	var avg float32
	for _, nums := range score{
		avg += float32(nums)
	}
	return avg/4
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
	stats := make([]studentStat, len(students))
	for i, st := range students{
		finalscore := average(st.test1Score, st.test2Score, st.test3Score, st.test4Score)
		if (finalscore < 35) {
			stats[i] = studentStat{student: st, finalScore: finalscore, grade: F}
		}else if finalscore >= 35 && finalscore < 50 {
			stats[i] = studentStat{student: st, finalScore: finalscore, grade: C}
		}else if finalscore >=50 && finalscore < 70 {
			stats[i] = studentStat{student: st, finalScore: finalscore, grade: B}
		}else if finalscore >= 70 {
			stats[i] = studentStat{student: st, finalScore: finalscore, grade: A}
		}

	}
	return stats
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	var maxgrad studentStat
	for _, stats := range gradedStudents{
		if stats.finalScore > maxgrad.finalScore{
			maxgrad.finalScore = stats.finalScore
			maxgrad.student = stats.student
			maxgrad.grade = stats.grade
		}
	}
	return studentStat{maxgrad.student, maxgrad.finalScore, maxgrad.grade}
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	top := make(map[string]studentStat)
	for _, stats := range gs {
		if stats.finalScore > top[stats.student.university].finalScore {
			top[stats.student.university] = stats
		}
	}
	return top
}

