package channelgoroutine_test

import (
	"fmt"
	"testing"
	"time"
)

type aStruct struct {
	value chan string
}

func TestConsumeChannel(t *testing.T) {
	aStr := aStruct{
		value: make(chan string),
	}

	go func() {
		populateData(aStr.value)
	}()

	go func() {
		for v := range aStr.value {
			fmt.Println(v)
		}
	}()

	<-time.After(5 * time.Second)
	fmt.Println("Program exited")
}

func populateData(c chan <- string){
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("counting %d", i)
	}
}
