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
import ("fmt")

func main() {
	
	// declare array with 
  var arr1 = [3]int{1,2,3}
  arr2 := [5]int{4,5,6,7,8}

  fmt.Println(arr1)
  fmt.Println(arr2)
}
 
 