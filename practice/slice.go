package main
import "fmt"

func main(){
	var a [5]int = [5]int{0,1,2,3,4}
	fmt.Println(a)
	var s []int = a[1:2]
	var t []int = a[1:]
	fmt.Println(s)
	fmt.Println(t)

}
