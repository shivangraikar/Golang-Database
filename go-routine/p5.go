
package main
import "fmt"
import "time"
import "sync"

var wg sync.WaitGroup
var available int = 10

func reserve(name string, wanted int) {
if available >= wanted {
	fmt.Println(wanted, "tickets blocked for ", name)
	time.Sleep(2 * time.Second)
	available = available - wanted
	fmt.Println(name, "have a safe joureny ")
} else {
	fmt.Println("tickets not available ")
}
wg.Done()
}


func main() {
fmt.Println("work started")
fmt.Println("initial available= ", available)
wg.Add(1)
go reserve("ravi",1)
wg.Add(1)
go reserve("kavi",1)
wg.Wait()
fmt.Println("final available = ", available)
}