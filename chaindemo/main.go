package main

import "fmt"

// Student 学生类
type Student struct {
	Name string
}

// Learn 学生类的学习方法
func (s *Student) Learn() *Student {
	fmt.Printf("%s热爱学习\n", s.Name)
	return s
}

// DoHomework 学生类的交作业方法
func (s *Student) DoHomework() *Student {
	fmt.Printf("%s喜欢交作业\n", s.Name)
	return s
}

func main() {
	s := &Student{Name: "张三"}
	s.Learn().DoHomework()
}
