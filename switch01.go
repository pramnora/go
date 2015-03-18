package main

import "fmt"

func main(){
   n:=1
   switch n{
    case 1: fmt.Println("n=1")
    case 2: fmt.Println("n=2")
    default: fmt.Println("n<>1/n<>2")
   }
}
