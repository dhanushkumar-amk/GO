package main
import ("fmt")

func main() {
  
	// ! WE CAN ALSO DECLARE THE MULTIPLE  VARIABLE IN THE SAME LINE 
	var a, b, c, d int = 1, 3, 5, 7
	// f you use the type keyword, it is only possible to declare one type of variable per line.

  fmt.Println(a)  
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
  
 
// output:
// 1
// 3
// 5
// 7

// If the type keyword is not specified, you can declare different types of variables in the same line:

var num, name,bol = 1, "hello", true;
fmt.Println(num);
fmt.Println(name);
fmt.Println(bol);
// output :
// 1
// hello
// true

// we can also declare the variable inside the group  of block element
var (
	e int  
	f int = 1
	g string = "hello"
  )

 fmt.Println(e) // no value assign so default value of int is 0
 fmt.Println(f)
 fmt.Println(g)
 
 // output 
//  0
//  1
//  hello

}