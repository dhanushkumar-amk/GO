// Slices are similar to arrays, but are more powerful and flexible.

// Like arrays, slices are also used to store multiple values of the same type in a single variable.

// However, unlike arrays, the length of a slice can grow and shrink as you see fit


// declare slice in three different ways 

// Using the []datatype{values} format
// Create a slice from an array
// Using the make() function


//  1 => Create a Slice With []datatype{values}

// slice_name := []datatype{values}
//myslice := []int{}
// myslice := []int{1,2,3}


// there are two function that can  return  length and capacity

// 1 => len()  to display the length;
// 2 => cap()  to display th capacity of the particular slice 

package main
import ("fmt")

func main() {
  mySlice1 := []int{}  // no value so default integer value is 0
  fmt.Println(len(mySlice1))  // 0
  fmt.Println(cap(mySlice1))  // 0
  fmt.Println(mySlice1)  //  return empty array

  mySlice2 := []string{"Go", "Slices", "Are", "Powerful"}
  fmt.Println(len(mySlice2)) // 4
  fmt.Println(cap(mySlice2))  // 4
  fmt.Println(mySlice2)// [Go Slices Are Powerful]
  
  
  // 2 => Create a Slice From an Array
  
  
//   var myarray = [length]datatype{values} // An array
// myslice := myarray[start:end] // A slice made from the array

arr1 := [6]int{10, 11, 12, 13, 14,15}
myslice3 := arr1[2:4]

fmt.Printf("myslice = %v\n", myslice3)
fmt.Println("length = ", len(myslice3))
fmt.Println("capacity = ", cap(myslice3))

// myslice = [12 13]
// length =  2  here myslice3 length is 2 because it hold only twi values
// capacity =  4  here he slice can grow to the end of the array. This means that the capacity of the slice is 4.

// If myslice started from element 0, the slice capacity would be 6


// Access Elements of a Slice
// You can access a specific slice element by referring to the index number.

var array = []int{1,2,3,4,5,6}

 fmt.Println(len(array))
 fmt.Println(cap(array))
 fmt.Println(array[0])
 fmt.Println(array[1])
 fmt.Println(array[2])
 fmt.Println(array[3])
 fmt.Println(array[4])
 fmt.Println(array[5])

// 1
// 2
// 3
// 4
// 5
// 6

//  Change Elements of a Slice like an array
var array2 = []int{1,2,3,4,5,6}
array2[0] = 0;
array2[1] = 0;
array2[2] = 0;
array2[3] = 0;
array2[4] = 0;

fmt.Println(array2)
// [0 0 0 0 0 6]


// Append(ADD) Elements To a Slice
// You can append elements to the end of a slice using the append()function:

var array3 = []int{1,2,3,4,5,6}
// creating new element and append
var arr = append(array3, 7,8,9,10)
fmt.Println(arr)
// [1 2 3 4 5 6 7 8 9 10]

// here append the element in an same
array3 = append(a)
}