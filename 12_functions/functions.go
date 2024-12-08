package main

import "fmt"


func add(a int ,b int ) int {   //fn can return multiple values
	return a+b

}

func getLanguages() (string,string,bool){
	return "golang","js",true
}

func processIt(fn func(a int)int ){   //function passed
	fn(1)                           //function returned
}


func processUp()func(b int)int {
	return func(a int) int {
		return 4        //return from function
	}
}



func main(){

	lang1,lang2,_:=getLanguages()    // "_" used if not using any variable to suppress it
	fmt.Println(lang1,lang2)
 

	 result :=add(3,5)
	 fmt.Println(result)


     //fmt.Println(getLanguages())


    fn:=func(a int) int {   //assigned to an variable
		return 2
	}
	 processIt(fn)


     fnn:=processUp()
     fnn(6)




}