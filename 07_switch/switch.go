package main 
import "fmt"
import "time"
func main(){
	//simple switch
	i:=5

	switch i {       //break not needed
	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("two")
  
	default:       //optional
		fmt.Println("other")

	}

//types of switch in GOLANG
   //--Multiple Condition Switch--
   switch time.Now().Weekday(){
   case time.Saturday , time.Sunday:
	fmt.Println("its weekend")
   default:
	fmt.Println("its work day")
   }


   //----TYPE SWITCH-----
   whoAmI:=func(i interface{}){
	switch t:= i.(type){    //switch  i.(type) if dont wanna use t
	case int:
		fmt.Println("its an integer")
	case string:
		fmt.Println("its an string")
	case bool:
		fmt.Println("its boolean")
	default:
		fmt.Println("other",t)
	}
   }

   whoAmI("golang")   //output -its a string


}