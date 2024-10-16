package main
import "fmt"

func main(){
	age:=18
	//if age was 16 then it wont give any output

    //----single loop---
	/*if age >=18 {
		fmt.Println("person is adult")
    } */

	


	//---if-else

	if age >=18 {
		fmt.Println("person is adult")
    } 
	else{
		fmt.Println("person is not an adult")

	}



//----else-if-------


	int number := 16

	if number >= 18 {
		fmt.Println("person is an adult")
	} else if  number >= 12 {
		fmt.Println("person is teenager")
	}else{
		fmt.Println("person is a kid")

	}



	//-logical operators in IF-ELSE---
	var role ="admin"
	var hasPermissions=true


	//OR-operator

	if role == "admin" ||  hasPermissions {        //any one is true then inside code will run
		fmt.Println("yes")

	}

	//AND-operator

	if role == "admin" &&  hasPermissions{  // both conditions should be true
		fmt.Println("yes")
      }






// FEATURE - variable can be declared inside IF itself
  if age :=15;age>=18
  {
	fmt.Println("person is an adult",age)
  }else if age >=12{
	fmt.Println("person is an teenager",age)

  }

  //GOLANG does not have ternary operator,you have to use normal if else



	
}