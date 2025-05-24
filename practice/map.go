package main
import "fmt"

func main(){
	var zoo map[string]int = map[string]int{
		"d":3,
		"sd":3,
		"ddd":3,
	}
	fmt.Println(zoo)
	val, ok := zoo["d"]
	fmt.Println(val, ok)


}
