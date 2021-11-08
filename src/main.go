package main

import (
	"fmt"
	"io"
	"memory_dbase/src/ydbase"
	"strconv"
	"strings"
)

const RED = "\x1b[31m"
const GREEN = "\x1b[32m"
const RESET = "\x1b[0m"

func main() {
	doLogin()
	runDB()
}

func doLogin() {
	fmt.Print("#> ")
	uname := ``
	upwd := ``
	for {
		n, err := fmt.Scanln(&uname, &upwd)
		if err == io.EOF {
			break
		}

		if n > 0 {
			if !ydbase.Login(uname, upwd) {
				fmt.Println("Login Failed")
				fmt.Print("#> ")
			} else {
				fmt.Printf("\t%sWelcome to use ydbase 1.0%s\n", GREEN, RESET)
				fmt.Printf("%s$> %s", RED, RESET)
				break
			}
		}
	}
}

func runDB() {
	cmd := ``
	exp := ``

	for {
		n, err := fmt.Scanln(&cmd, &exp)
		//fmt.Printf("%sDebug only: %s%s%s", RED, cmd, exp, RESET)
		if err == io.EOF {
			break
		}

		if n > 0 {
			if cmd == "init" {
				if err := cmdInit(exp); err != nil {
					fmt.Println("Parse exp failed in `cmdInit()` ->", err)
					fmt.Printf("%s$> %s", RED, RESET)
				}
			} else if cmd == "insert" {
				if err := cmdInsert(exp); err != nil {
					fmt.Println("Parse exp failed in `cmdInsert()` ->", err)
					fmt.Printf("%s$> %s", RED, RESET)
				}
			} else if cmd == "delete" {
				if err := cmdDelete(exp); err != nil {
					fmt.Println("Parse exp failed in `cmdDelete()` ->", err)
					fmt.Printf("%s$> %s", RED, RESET)
				}
			} else if cmd == "select" {
				if err := cmdSelect(exp); err != nil {
					fmt.Println("Parse exp failed in `cmdSelect()` ->", err)
					fmt.Printf("%s$> %s", RED, RESET)
				}
			} else if cmd == "update" {
				if err := cmdUpdate(exp); err != nil {
					fmt.Println("Parse exp failed in `cmdUpdate()` ->", err)
					fmt.Printf("%s$> %s", RED, RESET)
				}
			} else if cmd == "count" {
				fmt.Printf("Operation success, %d rows in total\n", ydbase.Count())
				fmt.Printf("%s$> %s", RED, RESET)
			} else if cmd == "exit" {
				break
			} else {
				fmt.Println("Unsupported syntax, try `insert`, `delete`, `update`, `exit`, `select`")
				fmt.Printf("%s$> %s", RED, RESET)
			}
		}
	}

	fmt.Println("See you!")
}

func cmdInit(exp string) error {
	len1, err := strconv.Atoi(exp)
	if err != nil {
		return err
	}

	for i := 1; i <= len1; i += 1 {
		stu := ydbase.Student{
			Id:     "09" + strconv.Itoa(i),
			Name:   "Name" + strconv.Itoa(i),
			Age:    22 + i,
			Height: 1.7,
		}

		ydbase.Insert(stu)
	}

	fmt.Printf("Operation success, %d rows in total\n", ydbase.Count())
	fmt.Printf("%s$> %s", RED, RESET)
	return nil
}

/// insert Id, Name, Age, Height
/// For example:
/// insert 001, jack, 33, 1.81
func cmdInsert(exp string) error {
	fields := strings.Split(exp, ",")
	leng := len(fields)
	if leng != 4 {
		fmt.Printf("%s$> %s", RED, RESET)
		return fmt.Errorf("`insert` command need 4 args!")
	}

	age, err := strconv.Atoi(fields[2])
	if err != nil {
		return err
	}
	height64, err := strconv.ParseFloat(fields[3], 32)
	if err != nil {
		return err
	}
	height := float32(height64)
	stu := ydbase.Student{
		Id:     fields[0],
		Name:   fields[1],
		Age:    age,
		Height: height,
	}

	ydbase.Insert(stu)
	fmt.Printf("Operation success, %d rows in total\n", ydbase.Count())
	fmt.Printf("%s$> %s", RED, RESET)
	return nil
}

func cmdDelete(exp string) error {
	if exp == `` {
		return fmt.Errorf("Need a number!")
	}
	if false == ydbase.Delete(exp) {
		return fmt.Errorf("Failed to delete")
	}

	fmt.Printf("Operation success, %d rows in total\n", ydbase.Count())
	fmt.Printf("%s$> %s", RED, RESET)
	return nil
}

func cmdSelect(exp string) error {
	newexp := strings.Replace(exp, "%", " like ", 1)
	stu := ydbase.Query(newexp)
	fmt.Printf("Operation success, match %d rows in total\n", len(stu))
	if len(stu) > 0 {
		ydbase.Print(stu)
	}
	fmt.Printf("%s$> %s", RED, RESET)
	return nil
}

func cmdUpdate(exp string) error {
	fields := strings.Split(exp, ",")
	leng := len(fields)
	if leng != 4 {
		fmt.Printf("%s$> %s", RED, RESET)
		return fmt.Errorf("`insert` command need 4 args!")
	}

	qryString := "id==" + fields[0]
	stu := ydbase.Query(qryString)
	if len(stu) == 1 {
		stu[0].Name = fields[1]

		age, err := strconv.Atoi(fields[2])
		if err != nil {
			return err
		}
		height64, err := strconv.ParseFloat(fields[3], 32)
		if err != nil {
			return err
		}
		height := float32(height64)

		stu[0].Age = age
		stu[0].Height = height

		ydbase.Modify(stu[0])
		fmt.Println("Operation success, 1 row affected")
		fmt.Printf("%s$> %s", RED, RESET)
		return nil
	}

	return fmt.Errorf("Failed in ydbase.Query()!")
}
