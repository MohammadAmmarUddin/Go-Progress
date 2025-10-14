package main

import "fmt"

// func add(a,b int){
// 	fmt.Println(a+b)

// }

func main(){
//anonymous func
//IIFE= IMMEDIATELY INVOKED FUNCTION EXPRESSION

func(a int,b int){
	c:=a+b
	fmt.Println(c)
}(5,7)
 
}


//func init()

func init(){
	fmt.Println("I am first called func")
}
