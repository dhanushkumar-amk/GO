// Go Constants

// If a variable should have a fixed value that cannot be changed, you can use the const keyword.

// The const keyword declares the variable as "constant", which means that it is unchangeable and read-only

// syntax :
// const CONSTNAME type = value

// he value of a constant must be assigned when you declare it.

package main
import ("fmt")

const PI = 3.14

func main() {
  fmt.Println(PI)
}