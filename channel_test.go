package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Aris Mahmudi"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(3 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Aris Mahmudi"
}

func GiveYouResponse(channel chan string) {
	time.Sleep(1 * time.Second)
	channel <- "Biarkan saja kucoba agar kau selalu di sisiku"
}

func TestChannelAsParam(t *testing.T) {
	channel := make(chan string)
	channel1 := make(chan string)
	defer close(channel)
	defer close(channel1)

	go GiveMeResponse(channel)
	go GiveYouResponse(channel1)

	data := <-channel
	data1 := <-channel1
	fmt.Println(data)
	fmt.Println(data1)

	time.Sleep(3 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Come on come on turn the radio on"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)
	go func() {

		channel <- "Aris"
		channel <- "Mahmudi"
		channel <- "Mavis"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
}
