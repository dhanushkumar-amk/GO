// Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.

// Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.

// 1 => With the var keyword
// 2 => With the := sign

// 1=> With the var keyword
// var array_name = [length]datatype{values}
// here length is defined

// or

// var array_name = [...]datatype{values}
// here length is inferred
//?  here the length defined by [...] we can add an number of values  the array

// 2 =>

// array_name := [length]datatype{values}
// here length is defined

// or

// array_name := [...]datatype{values}

// ?  here the length defined by [...] we can add an number of values  the array
// here length is inferred

package main

import (
	"fmt"
)

func main() {
	
	// declare array with length fixed
  var arr1 = [3]int{1,2,3}
  arr2 := [5]int{4,5,6,7,8}

  
  
  fmt.Println(arr1)
   fmt.Println(arr2)
   
   
   // output :
//    [1 2 3]
// [4 5 6 7 8]



// declares two arrays (arr1 and arr2) with inferred lengths:
// here we can add any number of  values in an array

var a1 = [...]int{1,2,3}
a2 := [...]int{4,5,6,7,8,6,5,4,3,5,6,7,5,4,3} // there is no limit to add the values 

fmt.Println(a1)
fmt.Println(a2)



// Array with string 

// with user defined length

var person = [3]string{"Dhanush", "sanjayKumar", "ashwin"};
fmt.Println("person : "   , person)
// person :  [Dhanush sanjayKumar ashwin]


// without defined length

var person1 = [...]string{"Dhanush", "sanjayKumar", "ashwin"};
fmt.Println("person : "   , person1)

// person :  [Dhanush sanjayKumar ashwin]



// Access Elements of an Array
// ou can access a specific array element by referring to the index number.
var person2 = [...]string{"Dhanush", "sanjayKumar", "ashwin"};
fmt.Println("person : "   , person2[0])
fmt.Println("person : "   , person2[1])
fmt.Println("person : "   , person2[2])

// person :  Dhanush
// person :  sanjayKumar
// person :  ashwin


var person2[2] = ""

}
 
 