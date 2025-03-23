package main
import "fmt"

func pangkat(base, exp int)int{
	if exp == 0{
		return 1
	}else{
		return base * pangkat(base, exp-1)
	}
}

func main(){
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Print(pangkat(a, b))
}
