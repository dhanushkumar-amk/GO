// Go has three functions to output text:

// Print()
// Println()
// Printf()






// The Printf() Function
// The Printf() function first formats its argument based on the given formatting verb and then prints them.

// Here we will use two formatting verbs:

// %v is used to print the value of the arguments
// %T is used to print the type of the arguments



package main
import ("fmt")

func main() {
  var i string = "Hello"
  var j int = 15

  // this can be done only on printf()
  fmt.Printf("i has value: %v and type: %T\n", i, i)
  fmt.Printf("j has value: %v and type: %T", j, j)
}

// i has value: Hello and type: string
// j has value: 15 and type: int