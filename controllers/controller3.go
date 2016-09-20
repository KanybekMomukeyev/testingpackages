
package controllers

import (
	"fmt"
	"time"
	"strings"
)

func SomeFunction123() int {

	fmt.Println("----------")
	fmt.Println("----------")

	func(){
		phrase := "These are the times that try men's souls\n"

		words := strings.Split(phrase, " ")

		ch := make(chan string, len(words))

		//phrase2 := "These are the times that try men's souls sdfdsf sdfdsfdsfsd ds fds f"
		//words2 := strings.Split(phrase2, " ")

		for _, word := range words {
			ch <- word
		}

		for i:=0; i < len(words); i++ {
			fmt.Print(<-ch + " ")
		}
	}()


	channel := make(chan string,1)

	channel <- "Hello"

	fmt.Println(<-channel)

	channel <- "Go"

	fmt.Println(<-channel)
	time.Sleep(time.Millisecond * 50)


	po := new(PurchaseOrder)
	po.Value = 42.27

	ch := make(chan *PurchaseOrder)

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
