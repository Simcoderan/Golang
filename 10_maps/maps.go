package main 
import ( 
	"fmt"
	"maps"
)
  





//maps->hash,pbjects,dictionary
func main() {
	//creating map
	m:=make(map[string]string)  //[key]value

	//add elments
	m["name"]="golang"
	m["area"]="backend"

	//get an element
	fmt.Println(m["name"],m["area"])

	//if accessing the key which doesnt exists
	//fmt.Println(m["phone"])     --------output-----empty value(zero value)

	n:=make(map[string]int)
	n["age"]=30
	fmt.Println(n["phone"])       // output-0 (zero value according to int type in value)


	//can get length of map
	fmt.Println(len(m))  //---2


	//delete from map
	delete(m,"area")
	fmt.Println(m)



	//to empty the whole map
	//clear(m)
	//fmt.Println(m)


	//without using   make fn we can make map
	//m:=map[string]int{"phone":40, "phones":3} - |
//                                                |
	//check element in map or not or take action  |
v,ok:=m["price"]                             //   |
	fmt.Println(v)      //we get values 3        <-

	if ok{
		fmt.Println("all ok")
	}else{
		fmt.Println("not ok")
	}      ///all ok



	//to check both maps are equal or not


	m1:=map[string]int{"phone":40, "phones":3}
	
	m2:=map[string]int{"phone":40, "phones":3}

	fmt.Println(maps.Equal(m1,m2))





	


}