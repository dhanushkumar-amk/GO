//  the if statement to specify a block of Go code to be executed if a condition is true.


// if condition {
  // code to be executed if condition is true
// }

package main
import ("fmt")

func main() {
  x:= 20
  y:= 18
  if x > y {
    fmt.Println("x is greater than y")
  }
  if x < y {
    fmt.Println("y is greater than x")
	
  }
  
  // if else statement
  
  // syntax:
  
  // if condition {
    // code to be executed if condition is true
  // } else {
    // code to be executed if condition is false
  // }
  
  time := 20
  if (time < 18) {
    fmt.Println("Good day.")
  } else {
    fmt.Println("Good evening.")
  }
}