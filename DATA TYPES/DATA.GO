// Go Data Types
// Data type is an important concept in programming. Data type specifies the size and type of variable values.

// Go is statically typed, meaning that once a variable type is defined, it can only store data of that type.

// Go has three basic data types:

// bool: represents a boolean value and is either true or false
// Numeric: represents integer types, floating point values, and complex types
// string: represents a string value






// example

package main
import ("fmt")

func main() {
  var a bool = true     // Boolean
  var b int = 5         // Integer
  var c float32 = 3.14  // Floating point number
  var d string = "Hi!"  // String

  fmt.Println("Boolean: ", a)
  fmt.Println("Integer: ", b)
  fmt.Println("Float:   ", c)
  fmt.Println("String:  ", d)
}