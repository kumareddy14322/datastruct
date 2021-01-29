package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)
func main(){


s := stack.New()
s.Push(2)
s.Push(3)
s.Push(4)
for s.Len() != 0 {

val := s.Pop()

fmt.Println(val, " ")

}

s.Push(10)
s.Push(20)
s.Push(30)

fmt.Println(s, "")


}