package main
import ("fmt" 
		"strconv" )
// func main(){
// 	fmt.Println("hello ")
// 	i := 23
// 	s := strconv.Itoa(i)
// 	fmt.Println(s)



// }

//string to int
func main()  {
	s:= "123"
	i,err := strconv.Atoi(s)
	if err != nil{
		fmt.Println("error" , err)

	}else{
		fmt.Println(i)
	}
}