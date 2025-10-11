package main

import "fmt"

/*
tuto-1
*/
// func main(){

// 	fmt.Println("hellog")

// }

/* tuto-2  */
func add(num1, num2 int){

	sum := num1+num2

	fmt.Println(sum)
}

/*
 func return value
*/


func getNumber(num1 int, num2 int) (int,int) {

 sum := num1 + num2 
 mul:= num1 * num2

 return sum,mul
}

func main(){

    a:=1
	b:=3
	  add(a,b)
	  fmt.Println(getNumber(2,3))
}

