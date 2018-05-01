package main

import (
	"fmt"
)

type Queue interface {
	Len() int
	Head() (interface{}, error)
	Enqueue(interface{}) error
	Dequeue() (interface{}, error)
}

type Ints struct {
	elements []int
}

func NewIntQueue(elements []int) *Ints {
	return &Ints{elements}
}

func (i *Ints) Len() int {
	return len(i.elements)
}

func (i *Ints) Head() (result interface{}, err error) {
	if i.Len() == 0 {
		return 0, fmt.Errorf("empty ints queue. ")
	}
	
	return i.elements[0], nil
}

func (i *Ints) Enqueue(v interface{}) (err error) {
	i.elements = append(i.elements, v.(int))
	return nil
}

func (i *Ints) Dequeue() (result interface{}, err error) {
	if i.Len() == 0 {
		return nil,fmt.Errorf("empty ints queue. ")
	}
	head := i.elements[0]
	i.elements = i.elements[1:]
	return head, nil
}


func main() {
	nums2 := []int{1,2,3}
	nums := NewIntQueue(nums2)
	var s Queue = nums
	fmt.Println(s.Head())
	s.Enqueue(4)
	fmt.Println(nums)
	s.Dequeue()
	fmt.Println(nums)
}
