package main

import "fmt"

func main() {
  var n, smallest, biggest int
  x := []int{
  49,90,80,68,
  57,85,68,70,
  37,34,83,27,
  19,97, 9,17,
  }

  for _,v:=range x {
    if v>n {
      fmt.Println(v,">",n)
      n = v
      biggest = n
    } else {
      fmt.Println(v,"<",n)
    }
  }

  fmt.Println("The biggest number is ", biggest)
  for _,v:=range x {
    if v>n {
      fmt.Println(v,">",n)
    } else {
      fmt.Println(v,"<",n)
      n = v
      smallest = n
    }
  }
  fmt.Println("The smallest number is ", smallest)
}