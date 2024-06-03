package main
import ("fmt")

func main() {
  var (
	// Although the + operator is often used to add together two values, it can also be used to add together a variable and a value, or a variable and another variable:
	
    sum1 = 100 + 50 // 150 (100 + 50)
    sum2 = sum1 + 250 // 400 (150 + 250)
    sum3 = sum2 + sum2 // 800 (400 + 400)
  )
  fmt.Println(sum3)
  
  
 // ! Athematic operator 
  
//   +	Addition	Adds together two values	x + y	
// -	Subtraction	Subtracts one value from another	x - y	
// *	Multiplication	Multiplies two values	x * y	
// /	Division	Divides one value by another	x / y	
// %	Modulus	Returns the division remainder	x % y	
// ++	Increment	Increases the value of a variable by 1	x++	
// --	Decrement	Decreases the value of a variable by 1	x--	


var a int =10;
var b int = 15; 
fmt.Println( a + b) // addition
fmt.Println( a - b) // Subtraction
fmt.Println( a * b)// multiplication
fmt.Println( a / b) // division
fmt.Println( a % b) // modulus
  
  // assignment operator
  // Assignment operators are used to assign values to variables.
  var x = 10
  fmt.Println(x)
  
  
  // assignment operator and shorthand assignment operator
//   =	x = 5	x = 5	
// +=	x += 3	x = x + 3	
// -=	x -= 3	x = x - 3	
// *=	x *= 3	x = x * 3	
// /=	x /= 3	x = x / 3	
// %=	x %= 3	x = x % 3	
// &=	x &= 3	x = x & 3	
// |=	x |= 3	x = x | 3	
// ^=	x ^= 3	x = x ^ 3	
// >>=	x >>= 3	x = x >> 3	
// <<=	x <<= 3	x = x << 3 


// comparison operator
// ==	Equal to	x == y	
// !=	Not equal	x != y	
// >	Greater than	x > y	
// <	Less than	x < y	
// >=	Greater than or equal to	x >= y	
// <=	Less than or equal to	x <= y



// logical operator
// && 	Logical and	Returns true if both statements are true	x < 5 &&  x < 10	
// || 	Logical or	Returns true if one of the statements is true	x < 5 || x < 4	
// !	Logical not	Reverse the result, returns false if the result is true	!(x < 5 && x < 10)
}

// bitwise operator
// & 	AND	Sets each bit to 1 if both bits are 1	x & y	
// |	OR	Sets each bit to 1 if one of two bits is 1	x | y	
//  ^	XOR	Sets each bit to 1 if only one of two bits is 1	x ^ b	
// <<	Zero fill left shift	Shift left by pushing zeros in from the right	x << 2	
// >>	Signed right shift	Shift right by pushing copies of the leftmost bit in from the left, and let the rightmost bits fall off	x >> 2

