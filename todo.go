package main

import (
	"time"
)

type isAlive struct {
	Name		string		`json:"name"`
	Version 	string		`json:"version"`
	Date 		time.Time	`json:"time now"`
} 

type Todo struct {
	Id 		int		`json:"id"`
	Name		string		`json:"name"`
	Completed 	string		`json:"completed"`
	Due		time.Time	`json:"due"`
}

type Todos []Todo 

