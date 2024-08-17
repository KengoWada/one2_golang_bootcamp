package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
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

func (s student) String() string {
	return fmt.Sprintf("%s %s from %s", s.firstName, s.lastName, s.university)
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func sliceAtoi(stringNumbers []string) ([]int, error) {
	var intNumbers []int
	for _, stringNumber := range stringNumbers {
		intNumber, err := strconv.Atoi(stringNumber)
		if err != nil {
			return nil, err
		}
		intNumbers = append(intNumbers, intNumber)
	}

	return intNumbers, nil
}

func average(scores ...int) float32 {
	var avg float32
	for _, score := range scores {
		avg += float32(score)
	}
	return avg / float32(len(scores))
}

func parseCSV(filePath string) []student {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to open file: "+filePath, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse csv file: "+filePath, err)
	}

	var students []student
	for _, record := range records[1:] {
		testScores, err := sliceAtoi(record[3:])
		if err != nil {
			log.Fatal("Invalid test score values")
		}

		var student = student{
			firstName:  record[0],
			lastName:   record[1],
			university: record[2],
			test1Score: testScores[0],
			test2Score: testScores[1],
			test3Score: testScores[2],
			test4Score: testScores[3],
		}
		students = append(students, student)
	}

	return students
}

func calculateGrade(students []student) []studentStat {
	var stats []studentStat
	for _, student := range students {
		finalScore := average(student.test1Score, student.test2Score, student.test3Score, student.test4Score)
		var grade Grade
		switch {
		case finalScore >= 70:
			grade = A
		case finalScore < 70 && finalScore >= 50:
			grade = B
		case finalScore < 50 && finalScore >= 35:
			grade = C
		default:
			grade = F
		}
		stats = append(stats, studentStat{student, float32(finalScore), grade})
	}
	return stats
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	topStudent := studentStat{}
	for _, gradedStudent := range gradedStudents {
		if gradedStudent.finalScore > topStudent.finalScore {
			topStudent = gradedStudent
		}
	}

	return topStudent
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	var topperPerUniversity = make(map[string]studentStat)

	for _, gradedStudent := range gs {
		_, ok := topperPerUniversity[gradedStudent.university]
		if !ok {
			topperPerUniversity[gradedStudent.university] = gradedStudent
			continue
		}

		var currentUniTopper = topperPerUniversity[gradedStudent.university]
		if gradedStudent.finalScore > currentUniTopper.finalScore {
			topperPerUniversity[gradedStudent.university] = gradedStudent
		}
	}

	return topperPerUniversity
}
