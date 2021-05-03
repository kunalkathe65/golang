package main

import "fmt"

func main(){

	// var colors map[string]string //creates empty map []
	// colors := make(map[string]string) //Alternative way

	colors := map[string]string{
		"red" : "#fff",
		"blue" : "#333",
	}

	colors["green"] = "#777"
	colors["red"] = "#888"
	delete(colors,"red")

	printMap(colors)
}

func printMap(m map[string]string){
	for color,hex := range m{
		fmt.Println("The hex of color",color ,"is",hex)
	}
}