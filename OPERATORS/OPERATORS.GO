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
}