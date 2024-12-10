package main

import "fmt"

func sum(nums...int)int{      //unc sum(nums...interface)int{   //any type
	total:=0

	for _, num:=range nums {
		total=total+num

	}
	return total

}



func main() {  //this is called variadic number of functions //n no. of parameters can be passed

	fmt.Println(1,2,3,5,"hello")  //no limit of parameters





//for specific type pf input value

       nums:=[]int{3,4,5,6}  //slice
	   result:=sum(nums...)
       // result:=sum(3,5,6,7)
		fmt.Println(result)

}