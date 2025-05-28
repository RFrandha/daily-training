package other

import "fmt"

func TryChannel() {
	ch := make(chan int)
	//go func() {
	//	time.Sleep(5 * time.Second)
	//}()
	ch <- 1
	go func() {
		fmt.Println(<-ch)
	}()
	//ch <- 2
	//close(ch)
	//fmt.Println(len(ch), cap(ch))
	//for v := range ch {
	//	fmt.Println(v)
	//} // prints 1
	//fmt.Println(<-ch)

	//ch <- 3 // would block if buffer is full
}
