package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

type Student struct {
	Name     string
	Password string
}

func (s *Student) String() string {
	return fmt.Sprintf("Name: %s, Password: %s\n", s.Name, s.Password)
}

func (s *Student) Set(_s string) error {
	studentArr := strings.Split(_s, ",")
	if len(studentArr) == 2 {
		s.Name = studentArr[0]
		s.Password = studentArr[1]
		return nil
	}
	return errors.New("error")
}

func main() {
	name := flag.String("name", "admin", "user name")
	s := new(Student)
	flag.Var(s, "student", "student info")

	flag.Parse()

	fmt.Println("hello,", *name)
	fmt.Println(s)
}
