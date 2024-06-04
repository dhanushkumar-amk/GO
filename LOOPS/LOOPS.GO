// loops in go

// 1 => for

// The for loop is the only loop available in Go.

// The for loop loops through a block of code a specified number of times.


// Syntax :
 
// for statement1; statement2; statement3 {
	// code to be executed for each iteration
//  }'



package main
import ("fmt")

func main() {
  for  i:='a'; i < 5; i++ {
    fmt.Println(i)
  }
}