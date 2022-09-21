package queue

import(
	"fmt"
)

const (
	MaxUint = ^uint(0)
	MAXSIZE = int(MaxUint >> 1)
)

type Queue struct{
	slice []int
}

func (q *Queue) NewQueue(){
	//s.top = -1
	q.slice = make([]int,0)
}

func (q *Queue) Enque(x int){
	if len(q.slice) > MAXSIZE - 1 {
		fmt.Println("Queue overflow")
	} else{
		q.slice = append(q.slice, x)
		//s.top++
	}
}

func (q *Queue) Deque() int {
	var x int
	length := len(q.slice)
	if len(q.slice) < 0{
		fmt.Println("Queue underflow")
		return -1
	} else{
		x = q.slice[0]
		q.slice = q.slice[1:length]
		fmt.Printf("slice %+v\n", q.slice)
	}
	return x
}

func (q *Queue) isEmpty() bool{
	return len(q.slice) <= 0
}

func CheckQueue(){
	queue := new(Queue)
	queue.NewQueue()
	fmt.Println(queue)
	queue.Enque(10)
	queue.Enque(12)
	queue.Enque(13)
	fmt.Println(queue)
	fmt.Println(queue.Deque())
	fmt.Println(queue.Deque())
	fmt.Println(queue.isEmpty())
	fmt.Println(queue.Deque())
	fmt.Println(queue.isEmpty())
}
