// package main
// import "fmt"

// func greedy(){
// 	fmt.Println("heeloo function")
// }
// func main(){
// 	greedy()

// }

// package main
// import "fmt"
// func add(a  , b int ) int {
// 	return  a+b

// }
// func main(){
// 	sum := add(3, 4)
// 	fmt.Println("sum:", sum)
// }

package main
import "fmt"

func factorial(n int) int {
	if n==0 {
		return 1

	}
	return n*factorial(n-1)


}
func main(){
	fmt.Println("factorial of 5:" , factorial(5))
}