package main

import "fmt"

// declare a Struct
type Job struct {
	role       string
	salary     int
	experience int
	location   string
}

func Struct() {
	var job1 Job

	//job1 specification
	job1.role = "backend developer"
	job1.salary = 1200000
	job1.experience = 1
	job1.location = "remote"

	//access the elements with function
	printJobDetails(job1)

	fmt.Println("-------------------")

	//access the elements with method
	job1.printJobDetails()

}

func printJobDetails(jobs Job) {
	//access the elements
	fmt.Println("role >> ", jobs.role)
	fmt.Println("salary >> ", jobs.salary)
	fmt.Println("experience >> ", jobs.experience)
	fmt.Println("location >> ", jobs.location)
}

// method of a struct
func (j Job) printJobDetails() {

	fmt.Println("role >> ", j.role)
	fmt.Println("salary >> ", j.salary)
	fmt.Println("experience >> ", j.experience)
	fmt.Println("location >> ", j.location)
}
