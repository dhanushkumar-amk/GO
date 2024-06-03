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

var array = [3]int{1,2,3,4,5,5}
 fmt.Println(arr[1])
 fmt.Println(arr[2])
 fmt.Println(arr[3])
 fmt.Println(arr[4])
 fmt.Println(arr[5])

}