package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestGoRoutines(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Hehehehe")
	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("This is", number)
}

func TestManyGoRoutines(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(1 * time.Second)
}
