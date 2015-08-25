package main
import (
		"fmt"
		"time"
		"strconv"
)

func main() {

	
	messages := make(chan string)
	channel_1 := make(chan int)
	var i int

	go func() { 
		
		t0 := time.Now()
		for i=0; i<=5000; i++{
			time.Sleep(50000)	
		}
		t1 := time.Now()

		var str string 
		str = "The call took " + (t1.Sub(t0)) + " to run"
		messages <- str
	}()



	go func() { 
		var i int
		i = 5+9
		channel_1 <- i 
	}()

	msg := <-messages

	msg1 := <- channel_1

	fmt.Println(msg)

	fmt.Println(msg1)
}