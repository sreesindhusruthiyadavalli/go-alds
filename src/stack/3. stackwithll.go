package stack

/*import (
	"fmt"
	"github.com/sreesindhusruthiyadavalli/go-alds/src/ll"
)

//const MaxUint = ^uint(0)
//const MAX_SIZE = int(MaxUint >> 1)

/*type Node struct{
	data int
	next *Nodego mo
}

func (n *Node) NewNode(data int, next *Node){
	n.data = data
	n.next = next
}

type LinkedList struct{
	head *Node
}

func (ll *LinkedList) Len() int{
	count := 0
	var temp *Node
	if ll.head == nil{
		return 0
	}
	temp = ll.head
	//fmt.Println(temp)
	for temp != nil {
		count ++
		temp = temp.next
		//fmt.Println(temp)
	}
	return count
}

func (ll *LinkedList) Print() {
	if ll.head == nil{
		fmt.Print("NULL")
		return
	}
	for ll.head != nil {
		fmt.Printf("%+v -->",ll.head.data)
		ll.head = ll.head.next
		ll.Print()
	}
}

func (ll *LinkedList) Insert(data int, position int){
	var currPosition int
	currentNode := ll.head
	temp := new(Node)
	temp.data = data
	//This inserts new element as linked list is empty
	if ll.head == nil{
		ll.head = temp
		return
	}
	//This inserts at beginning of the linked list.
	if position == 0{
		temp.next = ll.head
		ll.head = temp
	}
	for currPosition < position-1 && currentNode != nil {
		currentNode = currentNode.next
		currPosition ++
		//fmt.Println(currentNode.data, currPosition)
	}
	//Update the links
	existingNextNode := currentNode.next
	currentNode.next = temp
	currentNode = currentNode.next
	currentNode.next = existingNextNode
}*/

/*type Stackll struct{
	top int
	ll *ll.LinkedList
}

func (s *Stackll) NewStackll(){
	s.top = -1
	s.ll = new(LinkedList)
}

func (s *Stackll) Push(x int){
	if s.top > MAX_SIZE - 1 {
		fmt.Println("Stack overflow")
	} else{
		s.top++
		s.ll.Insert(x, s.top)
	}
}

func (s *Stackll) Pop() int {
	var x int
	if s.top < 0{
		fmt.Println("Stack underflow")
		return -1
	} else{
		current := s.ll.head
		var prev *Node
		for current.next != nil{
			prev = current
			current = current.next
		}
		fmt.Printf("previous %+v, next %+v\n", prev.data, current.data)
		x = current.data
		prev.next = nil
		s.top--
	}
	return x
}

func main(){
	stackll := new(Stackll)
	stackll.NewStackll()
	fmt.Println(stackll.ll.Len())
	stackll.Push(10)
	stackll.Push(12)
	stackll.Push(13)
	fmt.Println(stackll.ll.Len())
	fmt.Println(stackll.Pop())
	fmt.Println(stackll.ll.Len())
	//stackll.ll.Print()
	fmt.Println(stackll.Pop())
	stackll.ll.Print()
	//fmt.Println(stack.Peek())
	//fmt.Println(stack.isEmpty())
	//fmt.Println(stack.Pop())
	//fmt.Println(stack.isEmpty())
}*/


