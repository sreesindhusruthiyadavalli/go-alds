package ll

import (
	"fmt"
)

const(
	MaxUint = ^uint(0)
	MAXSIZE = int(MaxUint >> 1)
)

type Stackll struct{
	slice *LinkedList
}

func (s *Stackll) NewStackll(){
	s.slice = new(LinkedList)
}

func (s *Stackll) Push(x int){
	length := s.slice.length
	if length > MAXSIZE - 1 {
		fmt.Println("Stack overflow")
	} else{
		s.slice.Insert(x, length)
	}
}

func (s *Stackll) Pop() int {
	var x int
	length := s.slice.length
	if length < 0{
		fmt.Println("Stack underflow")
		return -1
	} else{
		current := s.slice.GetHead()
		var prev *Node
		for current.next != nil{
			prev = current
			current = current.next
		}
		fmt.Printf("previous %+v, next %+v\n", prev.data, current.data)
		x = current.data
		prev.next = nil
		s.slice.length--
	}
	return x
}

func CheckStackWithLl(){
	stackll := new(Stackll)
	stackll.NewStackll()
	fmt.Println(stackll.slice.length)
	stackll.Push(10)
	stackll.Push(12)
	stackll.Push(13)
	fmt.Println(stackll.slice.length)
	fmt.Println(stackll.Pop())
	fmt.Println(stackll.slice.length)
	//stackll.ll.Print()
	fmt.Println(stackll.Pop())
	stackll.slice.Print()
	//fmt.Println(stack.Peek())
	//fmt.Println(stack.isEmpty())
	//fmt.Println(stack.Pop())
	//fmt.Println(stack.isEmpty())
}


