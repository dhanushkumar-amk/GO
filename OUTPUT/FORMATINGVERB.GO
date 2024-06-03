// Formatting Verbs for Printf()
// Go offers several formatting verbs that can be used with the Printf() function.



// General Formatting Verbs
// The following verbs can be used with all data types:

// Verb	Description
// %v	Prints the value in the default format
// %#v	Prints the value in Go-syntax format
// %T	Prints the type of the value
// %%	Prints the % sign


// example for general forming verb 

package main
import ("fmt")

func main() {
  var i = 15.5
  var txt = "Hello World!"

  fmt.Printf("%v\n", i)  //  \n is used to get to new line
  fmt.Printf("%#v\n", i)
  fmt.Printf("%v%%\n", i)
  fmt.Printf("%T\n", i)
  fmt.Printf("")

  fmt.Printf("%v\n", txt)
  fmt.Printf("%#v\n", txt)
  fmt.Printf("%T\n", txt)
}

// output :
// 15.5
// 15.5
// 15.5%
// float64

// Hello World!
// "Hello World!"
// string