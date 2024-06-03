/*
Go Integer Data Types
Integer data types are used to store a whole number without decimals, like 35, -50, or 1345000.

The integer data type has two categories:

Signed integers - can store both positive and negative values
Unsigned integers - can only store non-negative values

*/


package main
import ("fmt")

func main() {
  var x int = 500
  var y int = -4500
  fmt.Printf("Type: %T, value: %v", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
}