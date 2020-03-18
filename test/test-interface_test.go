package test

import (
	"fmt"
	"testing"
)

/**
定义一个人Person接口
*/
type Person interface {
	introduceSelf()
}

/**
定义一个老师Teacher结构体
*/
type Teacher struct {
	role string
}

/**
老师实现introduceSelf接口
*/
func (t Teacher) introduceSelf() {
	fmt.Println("I am a Teacher")
}

type Student struct {
	role string
}

func (s Student) introduceSelf() {
	fmt.Println("I am a Student")
}

func TestInterface(t *testing.T) {
	fmt.Println("test-interface main")
	var person Person

	person = new(Teacher)
	person.introduceSelf()

	person = new(Student)
	person.introduceSelf()
}
