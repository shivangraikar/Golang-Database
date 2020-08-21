package main
import "fmt"




func add(a int, b int, ca chan int) {
ans := a + b
ca <- ans
}

func sub(a int, b int, cs chan int) {
ans := a - b
cs <- ans
}

func main() {
var c1 = make(chan int)
var c2 = make(chan int)
go add(10, 2, c1)
go sub(7, 3, c2)
r1 := <- c1
r2 := <- c2
r := r1 * r2
fmt.Println("result = ", r)
}
