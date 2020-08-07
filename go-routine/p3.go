package main
import "fmt"
import "time"
import "sync"

var wg sync.WaitGroup

func write() {
fmt.Println("writing work started")
for i :=1; i <=5; i++ {
fmt.Println("writing", i)
time.Sleep(1 * time.Second)
}
fmt.Println("writing work ended")
wg.Done()
}

func music() {
fmt.Println("music started")
for i :=1; i <=10; i++ {
fmt.Println("song", i)
time.Sleep(250 * time.Millisecond)
}
fmt.Println("music ended")
wg.Done()
}

func main() {
fmt.Println("work started")
wg.Add(1)
go music()
wg.Add(1)
go write()
wg.Wait()
fmt.Println("work done")

}