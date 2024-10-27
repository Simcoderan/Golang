package main

import "fmt"

//slices-dynamic arrays ,most used,usefull methods

func main() {

	var nums[]int //declaration-uninitialized is nil
	fmt.Println(nums==nil)
	
	

	var num =make([]int,0,5)
	var num =make([]int,2,5)  //[0 0]    //([]-array  ,  2-capacity  , 5- initial capacity)
	//capacity-> max no. of elements can fit 
	fmt.Println(cap(nums))   //print capacity
	fmt.Println(nums == nil )  //false 

	//elemnt add
	nums=append(nums,1)
	fmt.Println(nums)  // [0 0 1]

	nums=append(nums,2) //[0 0 1 2]
	nums=append(nums,3)  capacity double dynamically

	fmt.Println(cap(nums)) //5


	//way2 to print slice
	nums:=[]int{}
	nums=append(nums,1)


 
	var num =make([]int,0,5)
	nums=append(nums,2)
	var nums2 =make([]int,len(nums))
	

		//copy function
		copy(nums2,nums)


		fmt.Println(nums,nums2)


		//slice operator

		var nums=[]int{1,2,3}

		fmt.Println(nums[0:2])  //[1,2]

		//slice package
		var numss1=[]int{1,2}
		var numss2=[]int{1,2}
		//is it equal or not can be checked
		fmt.Println (slices.Equal(numss1,numss2))  ///return type boolean


		//2D Slice
		var nos=[][]int{{1,2,3},{4,5,6}}
		fmt.Println(nums)











}