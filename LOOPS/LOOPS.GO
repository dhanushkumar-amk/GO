// loops in go

// 1 => for

// The for loop is the only loop available in Go.

// The for loop loops through a block of code a specified number of times.

// Syntax :

// for statement1; statement2; statement3 {
// code to be executed for each iteration
//  }'

package main

import (
	"fmt"
)

func main() {
  for  i:=0; i < 5; i++ {
    fmt.Println(i)
  }
  
  
  
//   The continue Statement

// The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.

// ex : 
for i:=0; i < 5; i++ {
    if i == 3 {  //if i meets condition the ignore that i variable and loops runs continues  
      continue
    }
   fmt.Println(i)
}





// The break Statement

// The break statement is used to break/terminate the loop execution.

// ex : 
for i:=0; i < 5; i++ {
    if i == 3 {
      break   // here the i comes to meet condition the loop will be terminated
    }
   fmt.Println(i)
  }
  
  
  
  
  
  // The Range Keyword
  
// The range keyword is used to more easily iterate over an array, slice or map. It returns both the index and the value.
  
  // Syntax :
  
// for index, value := array|slice|map {
	// code to be executed for each iteration
//  }
  

fruits := [3]string{"apple", "orange", "banana"}

for idx, val := range fruits {
   fmt.Println( idx, val)   // (idx stores the index, val stores the value
}
  
// 0 apple
// 1 orange
// 2 banana 


// if u want to omit the index value the we use  
  
}