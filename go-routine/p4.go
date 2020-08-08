/* wagp to craete two function:	 numbers() -->1...5
				alphabets() -->a..e  */
package main
import "fmt"
import "time"
import "sync"

var wg sync.WaitGroup

func numbers() {
fmt.Println("writing work started")
for i :=1; i <=5; i++ {
fmt.Printf("%d\n", i)
time.Sleep(250 * time.Millisecond)
}
fmt.Println("writing work ended")
wg.Done()
}

func alphabets() {
fmt.Println("music started")
for i :='a'; i <='e'; i++ {
fmt.Printf("%c\n", i)
time.Sleep(250 * time.Millisecond)
}
fmt.Println("music ended")
wg.Done()
}

func main() {
fmt.Println("work started")
wg.Add(1)
go numbers()
wg.Add(1)
go alphabets()
wg.Wait()
fmt.Println("work done")

}