package main
import "fmt"



//iterating over data structures
func main(){

	nums:=[]int {6,7,8}

	//normal way to iterate

	//for i:=0;i<len(nums);i++{
		//fmt.Println(nums[i])
	//}
     sum:=0
	for_,num:=range nums{
		sum=sum+num
		//fmt.Println(num)
	}
	fmt.Println(sum)

}