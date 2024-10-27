package main
import "fmt"

//ARRAYS -numbered sequence of specific length
func main(){

	var nums [5]int   //declare

	fmt.Println(len(nums)) //to get length

	nums[0] = 1  //insert
	fmt.Println(nums[0]) //get number at any index

	fmt.Println(nums)  //print whole array  //OUTPUT-[1 0 0 0]


	var vals[4]bool
	fmt.Println(vals)   //output- [false  false false  false]


	nums:=[3]int{4,5,6}  //to declare in single  line 
	fmt.Println(nums) 


	//2D Arrays
	nums:=[2][2]int{{3,4},{5,6}}
	fmt.Println(nums)


}


//-fixed size,tht is predictable
//-memory optimization
//-constant time  access