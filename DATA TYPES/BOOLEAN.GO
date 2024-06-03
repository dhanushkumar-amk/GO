// Boolean Data Type
// A boolean data type is declared with the bool keyword and can only take the values true or false.


//  bool keyword


// The default value of a boolean data type is false.

// Example
// This example shows some different ways to declare Boolean variables:

package main
import ("fmt")

func main() {
  var b1 bool = true // typed declaration with initial value
  var b2 = true // untyped declaration with initial value
  var b3 bool // typed declaration without initial value then initial 
  b4 := true // untyped declaration with initial value

  fmt.Println(b1) // Returns true
  fmt.Println(b2) // Returns true
  fmt.Println(b3) // Returns false
  fmt.Println(b4) // Returns true
}