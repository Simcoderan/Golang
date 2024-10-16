package main 

import "fmt"
//before function it can be declared 
//but := short hand cant be used outside function
const age=30
const word="go"  //-whenever needed can be used in the program

func main(){
	const name  string = "golang" //type can skippked it infers itself

	//try to reassign the value its not allowed
	//-name="js"  ---not allowed

	/*int*/
	//const age=30
	fmt.Println(age)


	//--CONSTANT GROUPING--// (Whenever u need many constants)
	const (
		port=5000
		host="localhost"
	)
	fmt.Println(port,host)

}