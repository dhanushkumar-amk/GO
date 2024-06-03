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

// ! With the := sign:
//  ! Use the := sign, followed by the variable value:
// ! Can only be used inside functions

// Syntax
// variablename := value

// ! In this case, the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value)

// ? Variable declaration and value assignment cannot be done separately (must be done in the same line)

package main

import (
	"fmt"
)

func main() {
  var student1 string = "John" //type is string
  
  var student2 = "Jane" //type is inferred datatype can't be given because initial value given to the variable 
  
  
  x := 2 //type is inferred it mean the compiler decides the variable for the value declare within same line and also declare inside the function

  fmt.Println(student1) 
  fmt.Println(student2)
  fmt.Println(x)
  
  
//   Variable declaration with out initial value

	var a string
	var b int
	var c bool
  
	fmt.Println(a) // string => initial value empty 
	fmt.Println(b)  // int => initial value  = 0
	fmt.Println(c)   // boolean initial value = false
	
// 	In this example there are 3 variables:

// a
// b
// c
// These variables are declared but they have not been assigned initial values.

// By running the code, we can see that they already have the default values of their respective types:

// a is ""
// b is 0
// c is false




// we can also assign value after declaration

  var student string  // first we declare
  student = "John"    //the we assign the value
  fmt.Println(student)

	
  
  
  
  // This example shows declaring variables outside of a function, with the var keyword:


}





package main
import ("fmt")

var a int
var b int = 2
var c = 3


// using var we can declare the variable outside function 

func main() {
  a = 1
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}

// but declare the variable using := 

package main
import ("fmt")

a := 1

func main() {
  fmt.Println(a)
}