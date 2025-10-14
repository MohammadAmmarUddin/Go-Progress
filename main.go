package main

import "fmt"


func add(a,b int){
	fmt.Println(a+b)

}

func main(){
add(3,3)
 
}

func init(){
	fmt.Println("I'll be called first")
}
