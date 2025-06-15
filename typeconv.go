//not for impicit type converion

//explicit type 
package main
import "fmt"

// func printFloat(f float64){
// 	fmt.Println(f)
// }
// func main()  {
// 	var num1 int = 12
// 	var num2 float32 = 3.14
// 	result := float32(num1) + num2
// 	fmt.Println(result)
// 	var num int = 42
// 	printFloat(float64(num))
// }

// type conversion pitfalls
func main()  {
	var f float64 = 3.23
	var i int =  int(f)
	fmt.Println(i)

	
}