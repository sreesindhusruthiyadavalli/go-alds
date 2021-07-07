package main

import "fmt"

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
		fmt.Println("")
		return
	}
	temp := ll.head
	for temp != nil{
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func main(){
	llist := new(LinkedList)
	node1 := new(Node)
	node1.NewNode(3, nil)
	llist.head = node1
	//fmt.Println(llist.Len())
	node2 := new(Node)
	node2.NewNode(7, nil)
	llist.head.next = node2
	//fmt.Println(llist.Len())
	node3 := new(Node)
	node3.NewNode(10, nil)
	llist.head.next.next = node3
	llist.Print()
}





/*
output:
&{1 <nil>}
<nil>
1
&{1 0xc000010260}
&{2 <nil>}
<nil>
2
 */
