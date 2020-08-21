package main
import "fmt"
import "time"
var data = []int{}
var dpo = make (chan bool)

func producer() {
	for i:=1 ; i <= 5; i++ {
		data = append(data, i)
		time.Sleep(100 * time.Millisecond)
		fmt.Println("appending ", i)
	}
	dpo <- true
}

func consumer() {
r := <- dpo
if r {
	fmt.Println(data)
}
}

func main() {
go consumer()
go producer()
str := ""
fmt.Scanln(&str)
}