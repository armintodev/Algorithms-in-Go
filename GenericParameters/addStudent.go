package main

import (
	"fmt"
	"sort"
)

type Stringer = interface {
	String() string
}

type Integer int

func (i Integer) String() string {
	return fmt.Sprintf("%d", i)
}

type String string

func (s String) String() string {
	return string(s)
}

type Student struct {
	Name string
	Id   int
	Age  float64
}

func (s Student) String() string {
	return fmt.Sprintf("%s %d %0.2f", s.Name, s.Id, s.Age)
}

func addStudent[T any](students []T, student T) []T {
	return append(students, student)
}

func useStudent(bySort bool) {
	var students []String
	result := addStudent[String](students, "Armin")
	result = addStudent[String](result, "Ali")
	result = addStudent[String](result, "Reza")

	if bySort {
		sort.Sort(OrderedSlice[String](result))
	}

	fmt.Println(result)

	var students1 []Integer
	result1 := addStudent[Integer](students1, 150)
	result1 = addStudent[Integer](result1, 151)
	result1 = addStudent(result1, 152)

	if bySort {
		sort.Sort(OrderedSlice[Integer](result1))
	}

	fmt.Println(result1)

	var students2 []Student
	result2 := addStudent[Student](students2, Student{Name: "Armin", Id: 150, Age: 19})
	result2 = addStudent[Student](result2, Student{Name: "Ali", Id: 151, Age: 25})
	result2 = addStudent(result2, Student{Name: "Reza", Id: 152, Age: 34})

	if bySort {
		PerformSort[Student](result2, func(s1, s2 Student) bool {
			return s1.Age < s2.Age
		})
	}

	fmt.Println(result2)
}
