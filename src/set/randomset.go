/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/


package main
import "fmt"

type Set struct{
    dst map[int]bool
    indexDst map[int]int
    lst []int
}

func (s *Set) NewSet() {
        s.lst = make([]int,0)
        s.dst = make(map[int]bool)
        s.indexDst = make(map[int]int)
}

func (s *Set) Insert(val int) bool {
        //If key exists --> value already in the set
        if _, ok := s.dst[val]; ok {
            return false
        }
        s.dst[val] = true
        s.lst = append(s.lst, val)
        s.indexDst[val] = len(s.lst) - 1
        return true
}

func (s *Set) Remove(val int) bool {
        if _, ok := s.dst[val]; ok {
            //Delete the value from map
            delete (s.dst, val)
            length := len(s.lst)
            //Update the slice
            temp := s.lst[s.indexDst[val]] 
            s.lst[s.indexDst[val]] = s.lst[length -1]
            s.lst[length-1] = temp
            //Update the index map
            s.indexDst[s.lst[s.indexDst[val]]] = s.indexDst[val]
            delete(s.indexDst, val) 
            return true
            
        }
        return false
}

func (s *Set) Random() int {
    length := len(s.dst)
    return s.lst[length -1]
}

func (s *Set) PrintSet() {
    fmt.Println(s)
}

/*func main() {
    s := new(Set)
    s.NewSet()
    s.Insert(6)
    s.Insert(4)
    s.Insert(3)
    s.Insert(5)
    s.Insert(1)
    s.PrintSet()
    s.Remove(6)
    s.PrintSet()
}
*/
