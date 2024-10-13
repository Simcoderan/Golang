package main

import "fmt"

func main(){
	//string
	var name string ="golang"
	//--can be written as
	var word = "simran" //----> type automatically infer
	fmt.Println(name)
	fmt.Println(word)

	//boolean variable
	 var isAdult = true
	 fmt.Println(isAdult)

	 //number
	 var age int = 30
	 fmt.Println(age)

	 //short-hand--syntax
	 names:="sim"  //--used when we need to declared and assign at the same time
	 fmt.Println(names)


	 //situation where we need to mention the type --NO SHORTCUT--
    var words string //--not assign now

    //--afterwards assign
	words="go"
	fmt.Println(words)


	//float
	var price float32= 50.5
	fmt.Println(price)

	//----or----

	var prices =50.55
	fmt.Println(prices) //---it works
}