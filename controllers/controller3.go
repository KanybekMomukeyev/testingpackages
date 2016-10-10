
package controllers

import (
	"fmt"
	"time"
)

func SomeFunction123() int {

	mainMethod()
	return  2

	fmt.Println("----------")
	fmt.Println("----------")


	channel := make(chan string,2)
	channel <- "Hello"
	channel <- "Go"
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	time.Sleep(time.Millisecond * 50)


	chh := make(chan int, 2)
	chh <- 1
	chh <- 2
	fmt.Println(<-chh)
	fmt.Println(<-chh)


	po := new(PurchaseOrder)
	po.Value = 42.27

	ch := make(chan *PurchaseOrder)

	fmt.Printf("/n PO Number: %d\n", po)
	go SavePO(po, ch)

	newPo := <- ch
	fmt.Printf("/n PO Number: %d\n", newPo.Number)

	return  0
}

type PurchaseOrder struct {
	Number int
	Value float64
}

func SavePO(po *PurchaseOrder, callbackChannel chan *PurchaseOrder) {
	po.Number = 1234

	callbackChannel <- po
}

func mainMethod() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 1)
			c1 <- "one"
		}
		//close(c1)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 2)
			c2 <- "two"
		}
		//close(c2)
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	//for i := 0; i < 2; i++ {
	//	select {
	//	case msg1 := <-c1:
	//		fmt.Println("received", msg1)
	//	case msg2 := <-c2:
	//		fmt.Println("received", msg2)
	//	}
	//}
}



//package main
//
//import (
//	"fmt"
//	"time"
//	"math/rand"
//	"bufio"
//	"os"
//	"flag"
//)
//
////func boring(msg string) {
////	for i := 0; ; i++ {
////		fmt.Println(msg, i)
////		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
////	}
////}
////
////func main() {
////	go boring("boring!")
////	fmt.Println("I'm listening.")
////	time.Sleep(2 * time.Second)
////	fmt.Println("You're boring; I'm leaving.")
////}
//
//func boring(msg string, c chan string) {
//	for i := 0; ; i++ {
//		c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value.
//		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
//	}
//}
//
//func boring2(msg string) <-chan string { // Returns receive-only channel of strings.
//	c := make(chan string)
//	go func() { // We launch the goroutine from inside the function.
//		for i := 0; ; i++ {
//			c <- fmt.Sprintf("%s %d", msg, i)
//			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
//		}
//	}()
//	return c // Return the channel to the caller.
//}
//
//func fanIn(input1, input2 <-chan string) <-chan string {
//	c := make(chan string)
//	go func() { for { c <- <-input1 } }()
//	go func() { for { c <- <-input2 } }()
//	return c
//}
//
//func fake_fibonacci(n int, c chan int) {
//	for i := 0; i < n; i++ {
//		c <- i
//	}
//	//defer close(c)
//}
//
//func fibonacci(c, quit chan int) {
//	x, y := 0, 1
//	for {
//		select {
//		case c <- x:
//			x, y = y, x+y
//		case <-quit:
//			fmt.Println("quit")
//			return
//		}
//	}
//}
//
//func call_fibonacci()  {
//		c := make(chan int)
//		quit := make(chan int)
//		go func() {
//			for i := 0; i < 10; i++ {
//				fmt.Println(<-c)
//			}
//			quit <- 0
//		}()
//		fibonacci(c, quit)
//}
//
////func main() {
////
////	call_fibonacci()
////
////	//c := make(chan int, 10)
////	//go fake_fibonacci(cap(c), c)
////	//
////	//for i := 0; i < cap(c); i++ {
////	//	fmt.Println(<-c)
////	//}
////
////	//for i := range c {
////	//	fmt.Println(i)
////	//}
////}
//var (
//	msgc = make(chan string) // the message channel
//)
//
//func main() {
//	flag.Parse()
//
//	fmt.Println("start the program")
//	for {
//		// start the app
//		waitc := make(chan struct{}) // a wait lock
//
//		// start the client thread
//		go func() {
//			for {
//				msg := <-msgc // a message to send
//				print(msg)
//
//				if msg == "break" {
//					break
//				}
//
//			}
//		}()
//
//		// start the input thread
//		go func() {
//			for {
//				reader := bufio.NewReader(os.Stdin)
//				text, _ := reader.ReadString('\n')
//				msgc <- text
//			}
//		}()
//
//		<-waitc
//
//		// finished in this round restart the app
//		fmt.Println("restart the app")
//	}
//}
//
//
//func someMethod()  {
//
//
//	cc := fanIn(boring2("Joe"), boring2("Ann"))
//	for i := 0; i < 10; i++ {
//		fmt.Println(<-cc)
//	}
//	fmt.Println("You're both boring; I'm leaving.")
//
//
//	joe := boring2("Joe")
//	ann := boring2("Ann")
//
//	for i := 0; i < 5; i++ {
//		fmt.Println(<-joe)
//		fmt.Println(<-ann)
//	}
//
//	fmt.Println("You're both boring; I'm leaving.")
//
//	// create channel
//	c := make(chan string)
//
//	// call goroutine
//	go boring("boring!", c)
//
//	fmt.Printf("1_You say: %q\n", <-c) // Receive expression is just a value.
//	fmt.Printf("2_You say: %q\n", <-c) // Receive expression is just a value.
//	fmt.Printf("3_You say: %q\n", <-c) // Receive expression is just a value.
//	fmt.Printf("4_You say: %q\n", <-c) // Receive expression is just a value.
//	fmt.Printf("5_You say: %q\n", <-c) // Receive expression is just a value.
//	fmt.Printf("6_You say: %q\n", <-c) // Receive expression is just a value.
//
//	//// receive from goroutine
//	//for i := 0; i < 15; i++ {
//	//	fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value.
//	//}
//
//	time.Sleep(2 * time.Second)
//
//	fmt.Println("You're boring; I'm leaving.")
//
//}