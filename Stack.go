package main

import (
	"fmt"
)

type Stack interface {
	Len() int
	Top() (interface{}, error)
	Push(interface{}) error
	Pop() error
}

type Ints struct {
	elements []int
}

func NewIntStack(elements []int) *Ints {
	return &Ints{elements}
}

func (i *Ints) Len() int {
	return len(i.elements)
}

func (i *Ints) Top() (result interface{}, err error) {
	if i.Len() == 0 {
		return 0, fmt.Errorf("empty ints stack. ")
	}
	
	return i.elements[0], nil
}

func (i *Ints) Push(v interface{}) (err error) {
	i.elements = append(i.elements, v.(int))
	return nil
}

func (i *Ints) Pop() (err error) {
	if i.Len() == 0 {
		return fmt.Errorf("empty ints stack. ")
	}
	
	i.elements = i.elements[1:]
	return nil
}


func main() {
	nums2 := []int{1,2,3}
	nums := NewIntStack(nums2)
	var s Stack = nums
	fmt.Println(s.Top())
	s.Push(4)
	fmt.Println(nums)
	s.Pop()
	fmt.Println(nums)
}
