package main

import "fmt"

type Student struct {
	Name         string
	Age          int
	AverageGrade float64
}

func isHonorStudent(s Student) bool {
	if s.AverageGrade >= 4.5 {
		return true
	}
	return false
}
func main() {
	pavel := Student{"Pavel", 27, 4.8}
	tom := Student{"Tom", 18, 1.0}

	fmt.Println(isHonorStudent(pavel))
	fmt.Println(isHonorStudent(tom))

}
