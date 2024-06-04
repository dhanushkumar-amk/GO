// A function is a block of statements that can be used repeatedly in a program.

// A function will not execute automatically when a page loads.

// A function will be executed by a call to the function.



// Create a Function
// To create (often referred to as declare) a function, do the following:



// Use the func keyword.
// Specify a name for the function, followed by parentheses ().
// Finally, add code that defines what the function should do, inside curly braces {}.

// Syntax
// func FunctionName() {
//   code to be executed
// }

// function example

package main
import ("fmt")

// always declare the function outside the main function
func familyName(fname string) { // single parameter
	fmt.Println( fname)
  }
  
  
func myMessage() {
  fmt.Println("I just got executed!")
}

func main() {
  myMessage() // call the function
  
  
  
  
  
  
//    Parameters and Arguments


  familyName("Dhanush")
}

// Naming Rules for Go Functions

// A function name must start with a letter
// A function name can only contain alpha-numeric characters and underscores (A-z, 0-9, and _ )
// Function names are case-sensitive
// A function name cannot contain spaces
// If the function name consists of multiple words, techniques introduced for multi-word variable naming can be used