package ydbase

type Student struct {
	Id     string
	Name   string
	Age    int
	Height float32
}

var studentTable = []Student{}

func init() {
	studentTable = make([]Student, 0, 100)
}
