package ydbase

import (
	"strconv"
	"strings"
)

func Insert(stu Student) bool {
	lenTable := len(studentTable)
	capTable := cap(studentTable)
	if lenTable < capTable {
		studentTable = append(studentTable, stu)
	}
	studentTable[lenTable] = stu
	return true
}

func Count() int {
	return len(studentTable)
}

func Modify(stu Student) bool {
	for i, v := range studentTable {
		if v.Id == stu.Id {
			studentTable[i] = stu
			break
		}
	}
	return true
}

func Delete(id string) bool {
	for i, v := range studentTable {
		if v.Id == id {
			studentTable = append(studentTable[:i], studentTable[i+1:]...)
			return true
		}
	}
	return false
}

func Query(exp string) []Student {
	stu := make([]Student, 0)
	where := make([]string, 3)

	if strings.Contains(exp, ">=") {
		where = strings.Split(exp, ">=")
		if len(where) == 2 {
			for _, v := range studentTable {
				if strings.ToUpper(strings.Trim(where[0], " ")) == "AGE" {
					age, _ := strconv.Atoi(where[1])
					if v.Age >= age {
						stu = append(stu, v)
					}
				}
			}
		}
	}

	if strings.Contains(exp, "<=") {
		where = strings.Split(exp, "<=")
		if len(where) == 2 {
			for _, v := range studentTable {
				if strings.ToUpper(strings.Trim(where[0], " ")) == "AGE" {
					age, _ := strconv.Atoi(where[1])
					if v.Age <= age {
						stu = append(stu, v)
					}
				}
			}
		}
	}

	if strings.Contains(exp, "==") {
		where = strings.Split(exp, "==")
		if len(where) == 2 {
			for _, v := range studentTable {
				if strings.ToUpper(strings.Trim(where[0], " ")) == "ID" {
					if v.Id == strings.Trim(where[1], " ") {
						stu = append(stu, v)
					}
				}

				if strings.ToUpper(strings.Trim(where[0], " ")) == "NAME" {
					if v.Name == strings.Trim(where[1], " ") {
						stu = append(stu, v)
					}
				}
			}
		}
	}

	if strings.Contains(exp, "like") {
		where = strings.Split(exp, "like")
		if len(where) == 2 {
			for _, v := range studentTable {
				if strings.ToUpper(strings.Trim(where[0], " ")) == "ID" {
					if strings.Contains(v.Id, strings.Trim(where[1], " ")) {
						stu = append(stu, v)
					}
				}

				if strings.ToUpper(strings.Trim(where[0], " ")) == "NAME" {
					if strings.Contains(v.Name, strings.Trim(where[1], " ")) {
						stu = append(stu, v)
					}
				}
			}
		}
	}
	return stu
}
