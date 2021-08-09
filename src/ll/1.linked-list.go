package ll

import (
	"fmt"
)

type Node struct{
	data int
	next *Node
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
	// using a variable here to traverse the data
	// so that current linked list object remains same.
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
	//No Intermediate variable is used so we are traversing on
	//the actual object created and it's references gets updated on updating ll.head.
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
}

func CheckLl(){
	llist := new(LinkedList)
	//node1 := new(Node)
	//node1.NewNode(3, nil)
	////llist.head = node1
	////fmt.Println(llist.Len())
	//node2 := new(Node)
	//node2.NewNode(7, nil)
	////llist.head.next = node2
	////fmt.Println(llist.Len())
	//node3 := new(Node)
	//node3.NewNode(10, nil)
	////llist.head.next.next = node3
	////llist.Print()
	llist.Insert(3, 0)
	llist.Insert(7, 1)
	llist.Insert(10, 2)
	//llist.Print()
	llist.Len()
	llist.Insert(4, 2)
	llist.Print()
	llist.Len()
}





/*
3 -->7 -->10 -->NULL

 */
