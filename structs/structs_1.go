package main

import "fmt"

func main(){

	var v1,v2,v3 = 123, "Kunal", true
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)

	//Inline or Anonymous struct
	//type: struct{}
	Kunal := struct{
		firstName,lastName string
		salary  int
		fullTime  bool
	}{
		firstName : "Kunal",
		lastName : "Kathe",
		salary : 29800,
		fullTime : true,
	}

	fmt.Printf("%T",Kunal)
	fmt.Println(Kunal)
}