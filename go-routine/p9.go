package main
import "fmt"


var done1 = make(chan int)
var done2 = make(chan int)

func add(a int, b int) {
ans := a + b
done1 <- ans
}

func sub(a int, b int) {
ans := a - b
done2 <- ans
}

func main() {
go add(10, 2)
r1 := <- done1
go sub(7, 3)
r2 := <- done2
r := r1 * r2
fmt.Println("result = ", r)
}
