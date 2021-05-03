package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct{
	firstName string
	lastName string
	// contact contactInfo (ALTERNATIVE WAY)
	contactInfo //embedded struct
}

func main(){
	// alex := person{"Alex","Anderson"} (NOT RECOMMENDED)
	// var alex person (ASSIGN ZERO VALUES TO STRUCT PROPS)
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	alex := person{
		firstName : "Alex",
		lastName : "Anderson",
		contactInfo : contactInfo{
			email:"alex@xxx.com",
			zipCode :99900,
		},
	}

	// ptr := &alex
	// ptr.updateName("Kunal","Kathe") //ALTERNATIVE WAY)
	alex.updateName("Kunal","Kathe")
	alex.print()

	mySlice := []string{"Hi,","there!"}
	updateSlice(mySlice)

	num := 12
	test(&num)
	fmt.Println(num)
	
}	

func ( p person) print(){
	fmt.Println(p)
	fmt.Printf("%+v",p)
}

func ( ptr *person) updateName(newFirstName string,newLastName string){
	(*ptr).firstName = newFirstName
	(*ptr).lastName = newLastName
}

//because slice is reference type
func updateSlice(s []string){
	s[0] = "Bye!"
	fmt.Println(s)
}

func test(num *int){
	*num = 13
}