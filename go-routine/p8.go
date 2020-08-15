package main
import "fmt"
import "time"

var done1 = make(chan bool)
var done2 = make(chan bool)

func write() {
fmt.Println("writing work started")
for i :=1; i <=5; i++ {
fmt.Println("writing", i)
time.Sleep(1 * time.Second)
}
fmt.Println("writing work ended")
done1 <- true
}

func music() {
fmt.Println("music started")
for i :=1; i <=10; i++ {
fmt.Println("song", i)
time.Sleep(250 * time.Millisecond)
}
fmt.Println("music ended")
done2 <- true
}


func main() {
go music()
go write()
r1 := <- done1
r2 := <- done2
if r1  && r2 {
fmt.Println("work over")
}
}
