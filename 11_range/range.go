package main
import "fmt"



//iterating over data structures
func main(){

	nums:=[]int {6,7,8}

	//normal way to iterate

	//for i:=0;i<len(nums);i++{
		//fmt.Println(nums[i])
	//}
     //sum:=0
	for i ,num:=range nums{    //i=index
		//sum=sum+num
		//fmt.Println(num)
		fmt.Println(num,i)
	}
	//fmt.Println(sum)

	m:=map[string]string{"fname":"john","lname":"doe"}

	for k,v:=range m{
		fmt.Println(k,v)
	}

	//range for string
	for i,c:=range "golang"{
		fmt.Println(i,c)   //unicode code point "rune"(data structure)  for alpjhabets -output
	}                       //starting byte of rune
	                        //255-.1 byte ,2 bytes if more than 255(unicode has more than 255 as of assci value)

							
	//fmt.Println(i,string(c))  //for character printing						
 
}