package main
import "fmt"
import "time"

func write() {
fmt.Println("writing work started")
for i :=1; i <=5; i++ {
fmt.Println("writing", i)
time.Sleep(1 * time.Second)
}
fmt.Println("writing work ended")
}

func music() {
fmt.Println("music started")
for i :=1; i <=10; i++ {
fmt.Println("song", i)
time.Sleep(250 * time.Millisecond)
}
fmt.Println("music ended")
}

func main() {
go music()
go write()
time.Sleep(7 * time.Second)
}