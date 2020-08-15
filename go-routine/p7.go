
package main
import "fmt"
import "time"
import "sync"

var wg sync.WaitGroup
var balance float64 = 1000
var mu sync.Mutex

func deposit(amount float64) {

fmt.Println("deposite started")
mu.Lock()
amt1 := balance
time.Sleep(1 * time.Second)
amt1 = amt1 + amount
balance = amt1
fmt.Println("deposite ended")

mu.Unlock()
wg.Done()
}

func withdraw(amount float64) {

fmt.Println("withdraw started")
mu.Lock()
amt1 := balance
time.Sleep(1 * time.Second)
amt1 = amt1 - amount
balance = amt1
fmt.Println("withdraw ended")


mu.Unlock()
wg.Done()
}


func main() {
fmt.Println("initial balance= ", balance)
wg.Add(1)
go deposit(100)
wg.Add(1)
go withdraw(100)
wg.Wait()
fmt.Println("final balance= ", balance)
}