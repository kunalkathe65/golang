package main

import "fmt"

type Employee struct{
	firstName string
	lastName string
	salary int
	fullTime bool
}

func main(){

	// var1 := 12
	// var2 := 5.5
	var3 := 12 + 5.5
	fmt.Println(var3)

	//type : main.Employee
	Kunal := Employee{
		firstName :"Kunal",
		lastName :"Kathe",
		salary : 29800,
		fullTime : true,
	}

	//same as above
	//type : *main.Employee
	Karan := &Employee{
		firstName :"Kunal",
		lastName :"Kathe",
		salary : 29800,
		fullTime : true,
	}

	fmt.Printf("%T",Kunal)
	fmt.Printf("%T",Karan)

	fmt.Println((*Karan).firstName)
	fmt.Println(Karan.lastName) //same as above, Go will take care of dereferencing
}