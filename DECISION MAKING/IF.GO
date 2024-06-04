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
  
  
  
  // else if
  // if condition1 {
    // code to be executed if condition1 is true
//  } else if condition2 {
    // code to be executed if condition1 is false and condition2 is true
//  } else {
    // code to be executed if condition1 and condition2 are both false
//  }

time := 22
if time < 10 {
  fmt.Println("Good morning.")
  
  if time < 5{
    fmt.Println("welcome")
  }
} else if time < 20 {
  fmt.Println("Good day.")
} else {
  fmt.Println("Good evening.")
}
}