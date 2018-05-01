package main

import (
	"fmt"
)

type Stack interface {
	Len() int
	Top() (interface{}, error)
	Push(interface{}) error
	Pop() (interface{}, error)
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
	count := i.Len()
	if count == 0 {
		return 0, fmt.Errorf("empty ints stack. ")
	}
	
	return i.elements[count-1], nil
}

func (i *Ints) Push(v interface{}) (err error) {
	i.elements = append(i.elements, v.(int))
	return nil
}

func (i *Ints) Pop() (result interface{}, err error) {
	count := i.Len()
	if count == 0 {
		return nil, fmt.Errorf("empty ints stack. ")
	}
	
	top := i.elements[count-1]
	i.elements = i.elements[:count-1]
	return top, nil
}


func main() {
	nums2 := []int{1,2,3}
	nums := NewIntStack(nums2)
	var s Stack = nums
	fmt.Println(s.Top())
	s.Push(4)
	s.Push(5)
	s.Push(6)
	fmt.Println(nums)
	s.Pop()
	fmt.Println(nums)
}
