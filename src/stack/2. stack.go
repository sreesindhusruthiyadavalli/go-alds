package stack

import (
	"fmt"
)

const MaxUint = ^uint(0)
const MAXSIZE = int(MaxUint >> 1)

type Stack struct{
	//top int
	slice []int
}

func (s *Stack) NewStack(){
	//s.top = -1
	s.slice = make([]int,0)
}

func (s *Stack) Push(x int){
	if len(s.slice) > MAXSIZE - 1 {
		fmt.Println("Stack overflow")
	} else{
		s.slice = append(s.slice, x)
		//s.top++
	}
}

func (s *Stack) Pop() int{
	var x int
	length := len(s.slice)
	if len(s.slice) < 0{
		fmt.Println("Stack underflow")
		return -1
	} else{
		x = s.slice[length-1]
		s.slice = s.slice[:length-1]
		//s.top--
		//fmt.Printf("slice %+v, top %d", s.slice,s.top)
	}
	return x
}

func (s *Stack) Peek() int{
	var x int
	if len(s.slice) < 0 {
		fmt.Println("Stack underflow")
		return 0
	} else{
		x = s.slice[len(s.slice)-1]
	}
	return x
}

func (s *Stack) isEmpty() bool{
	return len(s.slice) < 0
}

func CheckStack(){
	stack := new(Stack)
	stack.NewStack()
	fmt.Println(stack)
	stack.Push(10)
	stack.Push(12)
	stack.Push(13)
	fmt.Println(stack)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Peek())
	fmt.Println(stack.isEmpty())
	fmt.Println(stack.Pop())
	fmt.Println(stack.isEmpty())
}


