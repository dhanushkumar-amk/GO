//  four types of declaration

// ? 1 => var
// ? 2 => const
// ? 3 => type
// ? 4 => func

// ! variable deceleration

// package main

// import "fmt"
 

// func main()  {
// 	// var variablename type = value
// 	var variable_name int = 2;
// 	fmt.Println("int: ", variable_name)
	
}


// Declare Variable
package main
import ("fmt")

func main() {
  var student1 string = "John" //type is string
  var student2 = "Jane" //type is inferred
  x := 2 //type is inferred

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)
}