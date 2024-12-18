package main

import "fmt"

  //decleare a Struct
type Job struct {
  role string
  salary int
  experience int
  location string
}

func Struct() {
  var job1 Job

  //job1 specification
  job1.role = "backend developer"
  job1.salary = 1200000
  job1.experience = 1
  job1.location = "remote"
  
  //access the elements
  printJobDetails(job1)
}

func printJobDetails(jobs Job) {
  //access the elements
  fmt.Println("role >> ", jobs.role)
  fmt.Println("salary >> ", jobs.salary)
  fmt.Println("experience >> ", jobs.experience)
  fmt.Println("location >> ", jobs.location)
}
