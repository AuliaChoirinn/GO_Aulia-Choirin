package main

import (
    "fmt"
)

type Student struct {
    name  string
    score int
}

func (s *Student) setName(name string) {
    s.name = name
}

func (s *Student) setScore(score int) {
    s.score = score
}

func main() {
    students := make([]Student, 5)
    var totalScore, minScore, maxScore int

    for i := 0; i < 5; i++ {
        var name string
        var score int
        fmt.Printf("Input %d student's name: ", i+1)
        fmt.Scanln(&name)
        fmt.Printf("Input %d student's score: ", i+1)
        fmt.Scanln(&score)

        student := Student{name: name, score: score}
        students[i] = student
        totalScore += score

        if i == 0 {
            minScore = score
            maxScore = score
        } else {
            if score < minScore {
                minScore = score
            }
            if score > maxScore {
                maxScore = score
            }
        }
    }

    avgScore := float64(totalScore) / float64(len(students))
    fmt.Printf("\nAverage score: %.2f\n", avgScore)

    fmt.Println("Student with minimum score:")
    for _, student := range students {
        if student.score == minScore {
            fmt.Printf("Name: %s, Score: %d\n", student.name, student.score)
        }
    }

    fmt.Println("Student with maximum score:")
    for _, student := range students {
        if student.score == maxScore {
            fmt.Printf("Name: %s, Score: %d\n", student.name, student.score)
        }
    }
}
