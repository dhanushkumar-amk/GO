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
  
  
  
  
  
//   Go has five keywords/types of signed integers:

// Type	Size	Range
// int	Depends on platform:
// 32 bits in 32 bit systems and
// 64 bit in 64 bit systems	-2147483648 to 2147483647 in 32 bit systems and
// -9223372036854775808 to 9223372036854775807 in 64 bit systems
// int8	8 bits/1 byte	-128 to 127
// int16	16 bits/2 byte	-32768 to 32767
// int32	32 bits/4 byte	-2147483648 to 2147483647
// int64	64 bits/8 byte	-9223372036854775808 to 9223372036854775807



// 1byte = 8 bits 


// Unsigned Integers
// Unsigned integers, declared with one of the uint keywords, can only store non-negative values:

var x uint = 500
var y uint = 4500
fmt.Printf("Type: %T, value: %v", x, x)
fmt.Printf("Type: %T, value: %v", y, y)

}