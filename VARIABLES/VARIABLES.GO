//  four types of declaration

// ? 1 => var
// ? 2 => const
// ? 3 => type
// ? 4 => func

// ! variable deceleration

// package main

// import "fmt"
 

// func main()  {
	// var variablename type = value
// 	var variable_name int = 2;
// 	fmt.Println("int: ", variable_name)
	
// }


// Declare Variable

// With the := sign:
// Use the := sign, followed by the variable value:

// Syntax
// variablename := value

// In this case, the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value)


package main
import ("fmt")

func main() {
  var student1 string = "John" //type is string
  var student2 = "Jane" //type is inferred
  x := 2 //type is inferred it mean the compiler decides the variable for the value

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)
  
  
//   Variable declaration with out initial value

	var a string
	var b int
	var c bool
  
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}


}