// The switch Statement
// Use the switch statement to select one of many code blocks to be executed.

// The switch statement in Go is similar to the ones in C, C++, Java, JavaScript, and PHP. The difference is that it only runs the matched case so it does not need a break statement.


// ? SYNTAX 


// switch expression {
// case x:
   // code block
// case y:
   // code block
// case z:
// ...
// default:
   // code block
// }





package main 

import fmt;


func main(){
    a := 2;
	
	switch a {
	case 1: {
		fmt.Println("Monday")
	}
	case 2: {
		fmt.Println("Tuesday")
	}
	case 3: {
		fmt.Println("WednesDay")
	}
	case 4: {
		fmt.Println("ThursDay")
	}
}
}