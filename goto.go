package main
import "fmt"

func main(){
	fmt.Println("start")
	goto skip
	fmt.Println("this line will skip")
skip:

	fmt.Println("eend")

}