package main

import . "fmt"

func main() {
	type students struct {
		name    string
		subject string
		score   int
		grade   string
	}
	var classes = map[string]students{}
	for {
		var className string
		var studentName string
		var subName string
		var subScore int
		var subGrade string
		Print("Enter class name: ")
		Scanln(&className)
		Print("Enter student name: ")
		Scanln(&studentName)
		Print("Enter subject name: ")
		Scanln(&subName)
		Print("Enter score : ")
		Scanln(&subScore)

		student_class := classes[className]
		student_class.name = studentName
		student_class.subject = subName
		student_class.score = subScore
		if subScore >= 70 {
			subGrade = "A"
		} else if subScore >= 60 && subScore < 70 {
			subGrade = "B"
		}
		student_class.grade = subGrade
		classes[className] = student_class
		Println(classes)
	}

}
