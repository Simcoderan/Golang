// for looping only "FOR" 
package main
import "fmt"

func main(){
	//WHILE LOOP FASHION - FOR
	i := 1
	for i<=3 {
		fmt.Println(i)
		i=i+1

	}

	//INFINITE LOOP
	/*
	for {
		 println ("1")
	}
		  */

	//FOR LOOP
	
	for i:=0;i<3;i++{
		fmt.Println(i)
	}

	//CONTINUE  - AND - BREAK is also used
	for i:=0;i<3;i++{
		//break - stops the loop
	}


	for i:=0;i<3;i++{
		//continue- current iterations stops
		if i=2{
			continue //--- 2 will never come
		}
	
	}

	//RANGE-want to do some thing for some times
	for i:= range 3 {
	fmt.Println(i)
	}       //OUTPUT - 1,2
     

}