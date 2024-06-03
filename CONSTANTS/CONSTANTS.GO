// Go Constants

// If a variable should have a fixed value that cannot be changed, you can use the const keyword.

// The const keyword declares the variable as "constant", which means that it is unchangeable and read-only

// syntax :
// const CONSTNAME type = value

// he value of a constant must be assigned when you declare it.

package main
import ("fmt")

const PI = 3.14

func main() {
  fmt.Println(PI)
  

  // constant types
  
  // two types of constants
//   Typed constants
// Untyped constants

// typed constants
// typed constants are declared with a defined type

const A int = 1 // here we declare int type
fmt.Println(A)

// Untyped constants
// Untyped constants are declared without a type

const B = 2  // here we can't declare the variable type
fmt.Println(B)

// constant variable can't be changed
const C = 1
C = 2
fmt.Println(C)
// error
// cannot assign to C



// we can also declare 
}



